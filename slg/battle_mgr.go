package slg

type BattleID int64

type Battle struct {
	ID     BattleID
	NodeID NodeID
	Units  []BattleUnitID
}

type BattleMgr struct {
	battles map[BattleID]*Battle
}

func NewBattleMgr() *BattleMgr {
	return &BattleMgr{
		battles: make(map[BattleID]*Battle),
	}
}

func (b *BattleMgr) AddBattle(bt *Battle) {
	b.battles[bt.ID] = bt
}
