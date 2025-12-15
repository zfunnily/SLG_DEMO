package slg

type MarchID int64

type March struct {
	ID       MarchID
	UnitID   BattleUnitID
	From     NodeID
	To       NodeID
	ArriveAt int64 // tick
}

type MarchMgr struct {
	marches map[MarchID]*March
}

func NewMarchMgr() *MarchMgr {
	return &MarchMgr{
		marches: make(map[MarchID]*March),
	}
}

func (m *MarchMgr) AddMarch(mc *March) {
	m.marches[mc.ID] = mc
}

func (m *MarchMgr) RemoveMarch(id MarchID) {
	delete(m.marches, id)
}
