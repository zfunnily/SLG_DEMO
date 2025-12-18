package uuid

import (
	"fmt"
	"sync"
	"time"
)

// Snowflake 雪花算法结构体
type Snowflake struct {
	mu            sync.Mutex
	lastTimestamp int64 // 上次生成ID的时间戳
	dataCenterID  int64 // 数据中心ID
	machineID     int64 // 机器ID
	sequence      int64 // 序列号
}

// 常量定义
const (
	// 时间戳起始点 (2024-01-01 00:00:00 UTC)
	epoch = int64(1704067200000)

	// 各部分的位数
	dataCenterIDBits = uint(5)
	machineIDBits    = uint(5)
	sequenceBits     = uint(12)

	// 各部分的最大值
	maxDataCenterID = int64(-1) ^ (int64(-1) << dataCenterIDBits)
	maxMachineID    = int64(-1) ^ (int64(-1) << machineIDBits)
	maxSequence     = int64(-1) ^ (int64(-1) << sequenceBits)

	// 各部分的偏移量
	timestampLeftShift = sequenceBits + machineIDBits + dataCenterIDBits
	dataCenterIDShift  = sequenceBits + machineIDBits
	machineIDShift     = sequenceBits
)

var (
	snowflake *Snowflake
	once      sync.Once
)

// Init 初始化雪花算法，如果参数无效则直接panic
func Init(dataCenterID, machineID int64) {
	once.Do(func() {
		// 验证数据中心ID和机器ID的有效性
		if dataCenterID < 0 || dataCenterID > maxDataCenterID {
			panic(fmt.Sprintf("dataCenterID must be between 0 and %d", maxDataCenterID))
		}

		if machineID < 0 || machineID > maxMachineID {
			panic(fmt.Sprintf("machineID must be between 0 and %d", maxMachineID))
		}

		snowflake = &Snowflake{
			lastTimestamp: 0,
			dataCenterID:  dataCenterID,
			machineID:     machineID,
			sequence:      0,
		}
	})
}

// GenerateID 生成唯一ID，如果出错则直接panic
func GenerateID() int64 {
	if snowflake == nil {
		panic("snowflake not initialized")
	}

	snowflake.mu.Lock()
	defer snowflake.mu.Unlock()

	// 获取当前时间戳（毫秒）
	timestamp := time.Now().UnixMilli()

	// 如果当前时间小于上次生成ID的时间戳，说明系统时钟回退了
	if timestamp < snowflake.lastTimestamp {
		panic(fmt.Sprintf("clock moved backwards. Refusing to generate id for %d milliseconds", snowflake.lastTimestamp-timestamp))
	}

	// 如果当前时间等于上次生成ID的时间戳，则递增序列号
	if timestamp == snowflake.lastTimestamp {
		snowflake.sequence = (snowflake.sequence + 1) & maxSequence

		// 如果序列号达到最大值，则等待下一个毫秒
		if snowflake.sequence == 0 {
			for timestamp <= snowflake.lastTimestamp {
				timestamp = time.Now().UnixMilli()
			}
		}
	} else {
		// 如果是新的毫秒，则重置序列号
		snowflake.sequence = 0
	}

	// 更新上次生成ID的时间戳
	snowflake.lastTimestamp = timestamp

	// 组合各个部分生成ID
	id := ((timestamp - epoch) << timestampLeftShift) |
		(snowflake.dataCenterID << dataCenterIDShift) |
		(snowflake.machineID << machineIDShift) |
		snowflake.sequence

	return id
}

// IDInfo 存储ID的各个组成部分
type IDInfo struct {
	Timestamp    int64     // 时间戳（毫秒）
	DataCenterID int64     // 数据中心ID
	MachineID    int64     // 机器ID
	Sequence     int64     // 序列号
	Time         time.Time // 生成时间
}

// ParseID 解析ID为各个组成部分，如果ID无效则直接panic
func ParseID(id int64) *IDInfo {
	if id < 0 {
		panic("invalid id")
	}

	timestamp := (id >> timestampLeftShift) + epoch
	dataCenterID := (id >> dataCenterIDShift) & maxDataCenterID
	machineID := (id >> machineIDShift) & maxMachineID
	sequence := id & maxSequence

	return &IDInfo{
		Timestamp:    timestamp,
		DataCenterID: dataCenterID,
		MachineID:    machineID,
		Sequence:     sequence,
		Time:         time.UnixMilli(timestamp),
	}
}
