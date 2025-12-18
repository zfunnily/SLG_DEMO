package march

import (
	"fmt"
	"slg_sever/internal/battle"
	"slg_sever/internal/event"
	"slg_sever/internal/node"
	timingwheel "slg_sever/pkg/timewheel"
	"slg_sever/pkg/uuid"
)

const (
	UNKNOWN int = iota
	MOVE
	WAIT
)

// MarchSpeed 移动速度（单位：tick/节点）
const MarchSpeed = 5

type EventSink func(event.Event)

type MarchMgr struct {
	marches   map[MarchID]*March
	moveWheel *timingwheel.TimingWheel
	sink      EventSink
}

func NewMarchMgr() *MarchMgr {
	return &MarchMgr{
		marches:   make(map[MarchID]*March),
		moveWheel: timingwheel.NewTimingWheel(),
	}
}

func (m *MarchMgr) RegisterEventSink(sink EventSink) {
	m.sink = sink
}

func (m *MarchMgr) emit(event event.Event) {
	if m.sink != nil {
		m.sink(event)
	}
}

// CreateMarch 创建新的行军
func (m *MarchMgr) CreateMarch(unitID battle.BattleUnitID, path []node.NodeID, currentTick int64) MarchID {
	if len(path) < 2 {
		return 0 // 路径至少需要两个节点
	}

	id := MarchID(uuid.GenerateID())

	// 简单计算行军时间：固定速度 * 距离（这里假设距离为1）
	arriveAt := currentTick + MarchSpeed

	// 创建行军对象
	march := &March{
		ID:       id,
		UnitID:   unitID,
		Path:     path,
		Index:    0,
		ArriveAt: arriveAt,
	}

	// 保存行军
	m.marches[id] = march

	// 在时间轮中添加到达事件
	m.moveWheel.AddTask(timingwheel.TaskID(id), arriveAt, currentTick)

	fmt.Printf("Created march %d from %d to %d, arrive at %d\n", id, path[0], path[1], arriveAt)

	return id
}

// ProcessMarchArrive 处理行军到达事件
func (m *MarchMgr) ProcessMarchArrive(marchID MarchID, currentTick int64) *ArriveEvent {
	march, exists := m.marches[marchID]
	if !exists {
		return nil
	}

	from := march.Path[march.Index]

	march.Index++

	isFinal := march.Index >= len(march.Path)-1
	to := march.Path[march.Index]
	e := &ArriveEvent{
		MarchID:  march.ID,
		UnitID:   march.UnitID,
		FromNode: from,
		ToNode:   to,
		ArriveAt: march.ArriveAt,
		IsFinal:  isFinal,
	}

	if isFinal {
		delete(m.marches, marchID)
		return e
	}

	// 继续下一段行军
	march.ArriveAt += MarchSpeed
	m.moveWheel.AddTask(
		timingwheel.TaskID(marchID),
		march.ArriveAt,
		currentTick,
	)

	return e
}

// SpeedUpMarch 加速行军
func (m *MarchMgr) SpeedUpMarch(marchID MarchID, speedUpTicks int64, currentTick int64) {
	march, exists := m.marches[marchID]
	if !exists {
		return
	}

	// 计算新的到达时间
	newArriveAt := march.ArriveAt - speedUpTicks
	if newArriveAt < currentTick {
		newArriveAt = currentTick
	}

	// 更新行军的到达时间
	march.ArriveAt = newArriveAt

	// 重新在时间轮中添加任务
	m.moveWheel.RemoveTask(timingwheel.TaskID(marchID))
	m.moveWheel.AddTask(timingwheel.TaskID(marchID), newArriveAt, currentTick)

	fmt.Printf("Speed up march %d, new arrive at %d\n", marchID, newArriveAt)
}

// Tick 处理时间轮Tick
func (m *MarchMgr) Tick(currentTick int64) {
	m.moveWheel.Tick(currentTick, func(taskID timingwheel.TaskID) {
		// 解析任务ID，获取行军ID
		marchID := MarchID(taskID)
		// 处理行军到达
		e := m.ProcessMarchArrive(marchID, currentTick)
		if e != nil {
			m.emit(e)
		}
	})
}

// GetMarch 获取行军信息
func (m *MarchMgr) GetMarch(marchID MarchID) (*March, bool) {
	march, exists := m.marches[marchID]
	return march, exists
}

// AddMarch 添加行军
func (m *MarchMgr) AddMarch(mc *March) {
	m.marches[mc.ID] = mc
}

// RemoveMarch 移除行军
func (m *MarchMgr) RemoveMarch(id MarchID) {
	delete(m.marches, id)
}
