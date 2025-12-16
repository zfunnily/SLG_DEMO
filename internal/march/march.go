package march

import (
	"slg_sever/internal/battle"
	"slg_sever/internal/node"
)

type MarchID int64

type March struct {
	ID       MarchID
	UnitID   battle.BattleUnitID
	From     node.NodeID
	To       node.NodeID
	ArriveAt int64 // tick
}
