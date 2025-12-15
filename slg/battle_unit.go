package slg

type BattleUnitID int64

type BattleUnit struct {
	ID       BattleUnitID
	PlayerID int64
	NodeID   NodeID

	Hp  int64
	Atk int64
	Def int64
}
