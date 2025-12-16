package bootstrap

import (
	"encoding/json"
	"os"
	"slg_sever/config"
	"slg_sever/internal/battle"
	"slg_sever/internal/march"
	"slg_sever/internal/node"
	"slg_sever/internal/player"
	"slg_sever/internal/world"
)

var slgMap *world.World

func GetSLGMap() *world.World {
	return slgMap
}

func loadNode(m *world.World, n config.MapNode, nodeType node.NodeType) {
	m.Nodes[node.NodeID(n.ID)] = &node.NodeInfo{
		ID:   node.NodeID(n.ID),
		Type: nodeType,
	}
	for _, to := range n.Links {
		m.Graph.AddEdge(node.NodeID(n.ID), node.NodeID(to))
	}
}

func NewSLGMap() *world.World {
	return &world.World{
		Graph:     node.NewGraph(),
		Nodes:     make(map[node.NodeID]*node.NodeInfo),
		MarchMgr:  march.NewMarchMgr(),
		PlayerMgr: player.NewPlayerMgr(),
		BattleMgr: battle.NewBattleMgr(),
		PrintTick: 0,
	}
}

func InitWorld() {
	data, err := os.ReadFile("data/map.json")
	if err != nil {
		panic(err)
	}

	var cfg config.MapConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	slgMap = NewSLGMap()

	loadNode(slgMap, cfg.Capital, node.NodeCapital)

	for _, n := range cfg.Large {
		loadNode(slgMap, n, node.NodeLarge)
	}
	for _, n := range cfg.Medium {
		loadNode(slgMap, n, node.NodeMiddle)
	}
	for _, n := range cfg.HQ {
		loadNode(slgMap, n, node.NodeBase)
	}
	for _, n := range cfg.Small {
		loadNode(slgMap, n, node.NodeSmall)
	}
}
