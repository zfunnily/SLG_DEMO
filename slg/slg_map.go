package slg

type SLGMap struct {
	Graph *Graph
	Nodes map[NodeID]*NodeInfo

	MarchMgr  *MarchMgr
	PlayerMgr *PlayerMgr
	BattleMgr *BattleMgr

	printTick int64
}

func (m *SLGMap) Tick(tick int64) {
	if m.printTick == 0 {
		m.printTick = tick
	}

	print(tick)
}
