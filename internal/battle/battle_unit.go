package battle

import "slg_sever/internal/node"

type BattleUnitID int64

type BattleUnit struct {
	ID       BattleUnitID
	PlayerID int64
	NodeID   node.NodeID

	Hp  int64
	Atk int64
	Def int64
}
