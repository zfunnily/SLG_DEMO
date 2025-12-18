package world

import (
	"slg_sever/internal/battle"
	"slg_sever/internal/event"
	"slg_sever/internal/march"
	"slg_sever/internal/node"
	"slg_sever/internal/player"
)

type World struct {
	Graph *node.Graph
	Nodes map[node.NodeID]*node.Node

	MarchMgr  *march.MarchMgr
	PlayerMgr *player.PlayerMgr
	BattleMgr *battle.BattleMgr

	events []event.Event

	PrintTick int64
	NowTick   int64
}

func NewWorld() *World {
	w := &World{
		Graph:     node.NewGraph(),
		Nodes:     make(map[node.NodeID]*node.Node),
		MarchMgr:  march.NewMarchMgr(),
		PlayerMgr: player.NewPlayerMgr(),
		BattleMgr: battle.NewBattleMgr(),
		PrintTick: 0,
	}

	w.MarchMgr.RegisterEventSink(w.onEvent)

	return w
}

func (w *World) onEvent(ev event.Event) {
	w.events = append(w.events, ev)
}

func (w *World) GetNowTick() int64 {
	return w.NowTick
}

func (w *World) Tick(tick int64) {
	w.NowTick = tick

	if w.PrintTick == 0 {
		w.PrintTick = tick
	}

	if tick-w.PrintTick >= 10 {
		println("now_tick.", tick)
		w.PrintTick = tick
	}

	w.MarchMgr.Tick(tick)

	for _, ev := range w.events {
		w.dispatch(ev)
	}

	w.events = w.events[:0]
}

func (w *World) dispatch(ev event.Event) {
	switch e := ev.(type) {
	case *march.ArriveEvent:
		w.OnMarchArrive(e)

	default:
		// unknown event
	}
}

func (w *World) OnMarchArrive(e *march.ArriveEvent) {
	println("march.arrive MarchID%d, unitID:%d, fromNode:%d, toNode:%d, arriveAt:%d, isFinal:%v", e.MarchID, e.UnitID, e.FromNode, e.ToNode, e.ArriveAt, e.IsFinal)
}
