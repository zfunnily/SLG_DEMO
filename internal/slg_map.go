package internal

import (
	"encoding/json"
	"os"
	"slg_sever/config"
	"slg_sever/slg"
)

var slgMap *slg.SLGMap

func GetSLGMap() *slg.SLGMap {
	return slgMap
}

func loadNode(m *slg.SLGMap, n config.MapNode, nodeType slg.NodeType) {
	m.Nodes[slg.NodeID(n.ID)] = &slg.NodeInfo{
		ID:   slg.NodeID(n.ID),
		Type: nodeType,
	}
	for _, to := range n.Links {
		m.Graph.AddEdge(slg.NodeID(n.ID), slg.NodeID(to))
	}
}

func NewSLGMap() *slg.SLGMap {
	return &slg.SLGMap{
		Graph:     slg.NewGraph(),
		Nodes:     make(map[slg.NodeID]*slg.NodeInfo),
		MarchMgr:  slg.NewMarchMgr(),
		PlayerMgr: slg.NewPlayerMgr(),
		BattleMgr: slg.NewBattleMgr(),
	}
}

func init() {
	data, err := os.ReadFile("data/map.json")
	if err != nil {
		panic(err)
	}

	var cfg config.MapConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	slgMap := NewSLGMap()

	loadNode(slgMap, cfg.Capital, slg.NodeCapital)

	for _, n := range cfg.Large {
		loadNode(slgMap, n, slg.NodeLarge)
	}
	for _, n := range cfg.Medium {
		loadNode(slgMap, n, slg.NodeMiddle)
	}
	for _, n := range cfg.HQ {
		loadNode(slgMap, n, slg.NodeBase)
	}
	for _, n := range cfg.Small {
		loadNode(slgMap, n, slg.NodeSmall)
	}
}
