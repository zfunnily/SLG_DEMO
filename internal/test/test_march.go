package test

import (
	"slg_sever/internal/battle"
	"slg_sever/internal/config"
	"slg_sever/internal/global"
	"slg_sever/internal/node"
)

func StartMarch(data *config.MarchData) {
	println("test StartMarch...")

	currentTick := global.GetWorld().GetNowTick()
	for _, unit := range data.MarchUnits {
		p := make([]node.NodeID, 0)
		for _, a := range unit.Path {
			p = append(p, node.NodeID(a))
		}

		global.GetWorld().MarchMgr.CreateMarch(battle.BattleUnitID(unit.PlayerID), p, currentTick)
	}
}
