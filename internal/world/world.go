package world

import (
	"slg_sever/internal/battle"
	"slg_sever/internal/march"
	"slg_sever/internal/node"
	"slg_sever/internal/player"
)

type World struct {
	Graph *node.Graph
	Nodes map[node.NodeID]*node.NodeInfo

	MarchMgr  *march.MarchMgr
	PlayerMgr *player.PlayerMgr
	BattleMgr *battle.BattleMgr

	PrintTick int64
}

func (m *World) Tick(tick int64) {
	if m.PrintTick == 0 {
		m.PrintTick = tick
	}

	if tick-m.PrintTick >= 10 {
		println(tick)
		m.PrintTick = tick
	}
}
