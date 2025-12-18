package march

import (
	"slg_sever/internal/battle"
	"slg_sever/internal/node"
)

// ArriveEvent 行军到达事件
type ArriveEvent struct {
	MarchID  MarchID
	UnitID   battle.BattleUnitID
	FromNode node.NodeID
	ToNode   node.NodeID
	ArriveAt int64
	IsFinal  bool // 是否到达终点
}

func (e *ArriveEvent) EventType() string {
	return "march.arrive"
}
