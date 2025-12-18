package timingwheel

import (
	"container/list"
)

type TaskID int64 // 任务唯一标识

// 任务结构体
type task struct {
	id         TaskID
	expireTick int64
}

type wheelLevel struct {
	slots     []list.List
	slotCount int
	interval  int64 // 本层每个 slot 覆盖 tick 数
	cursor    int   // 当前槽索引，用于下沉
}

type TimingWheel struct {
	levels []*wheelLevel
}

func NewTimingWheel() *TimingWheel {
	tw := &TimingWheel{}

	tw.levels = []*wheelLevel{
		newWheelLevel(60, 1),    // 秒轮
		newWheelLevel(60, 60),   // 分轮
		newWheelLevel(24, 3600), // 小时轮
	}

	return tw
}

func newWheelLevel(slotCount int, interval int64) *wheelLevel {
	slots := make([]list.List, slotCount)
	return &wheelLevel{
		slots:     slots,
		slotCount: slotCount,
		interval:  interval,
		cursor:    0,
	}
}

// AddTask 添加任务到时间轮
func (tw *TimingWheel) AddTask(taskID TaskID, expireTick int64, currentTick int64) {
	diff := expireTick - currentTick
	tw.addTask(taskID, expireTick, diff)
}

// addTask 内部方法：将任务添加到合适的层级和槽位
func (tw *TimingWheel) addTask(taskID TaskID, expireTick int64, diff int64) {
	for _, level := range tw.levels {
		span := int64(level.slotCount) * level.interval
		if diff < span {
			slot := (level.cursor + int(diff/level.interval)) % level.slotCount
			level.slots[slot].PushBack(&task{
				id:         taskID,
				expireTick: expireTick,
			})
			return
		}
	}

	// 超过最大范围，放在最高层最后一个槽
	last := tw.levels[len(tw.levels)-1]
	last.slots[last.slotCount-1].PushBack(&task{
		id:         taskID,
		expireTick: expireTick,
	})
}

// RemoveTask 移除指定ID的任务
func (tw *TimingWheel) RemoveTask(taskID TaskID) {
	// 遍历所有层级
	for _, level := range tw.levels {
		// 遍历层级中的所有槽位
		for i := 0; i < level.slotCount; i++ {
			slot := &level.slots[i]
			// 遍历槽位中的所有任务
			for e := slot.Front(); e != nil; {
				next := e.Next()
				t := e.Value.(*task)
				// 如果找到匹配的任务ID，移除该任务
				if t.id == taskID {
					slot.Remove(e)
					return // 任务ID应该是唯一的，找到后直接返回
				}
				e = next
			}
		}
	}
}

// Tick 处理时间轮Tick，处理所有到期的任务
func (tw *TimingWheel) Tick(currentTick int64, handle func(TaskID)) {
	tw.advance(0, currentTick, handle)
}

// advance 内部方法：处理指定层级的tick
func (tw *TimingWheel) advance(levelIdx int, currentTick int64, handle func(TaskID)) {
	level := tw.levels[levelIdx]
	slot := &level.slots[level.cursor]

	// 收集所有需要处理的任务
	var tasksToProcess []*task
	for e := slot.Front(); e != nil; {
		next := e.Next()
		t := e.Value.(*task)

		if t.expireTick <= currentTick {
			// 任务已到期，需要处理
			tasksToProcess = append(tasksToProcess, t)
		} else {
			// 任务未到期，需要下沉到下一层
			diff := t.expireTick - currentTick
			tw.addTask(t.id, t.expireTick, diff)
		}

		// 从当前槽位移除任务
		slot.Remove(e)
		e = next
	}

	// 处理所有到期的任务
	for _, t := range tasksToProcess {
		handle(t.id)
	}

	// 更新当前层级的游标
	level.cursor = (level.cursor + 1) % level.slotCount

	// 如果游标回到0，说明当前层级已经走完一圈，需要进位到上一层
	if level.cursor == 0 && levelIdx+1 < len(tw.levels) {
		tw.advance(levelIdx+1, currentTick, handle)
	}
}
