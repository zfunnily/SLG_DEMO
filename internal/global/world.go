package global

import (
	"encoding/json"
	"os"
	"slg_sever/internal/config"
	"slg_sever/internal/node"
	"slg_sever/internal/world"
)

var slgWorld *world.World

func GetWorld() *world.World {
	return slgWorld
}

func loadNode(m *world.World, n config.MapNode, nodeType node.NodeType) {
	m.Nodes[node.NodeID(n.ID)] = &node.Node{
		ID:   node.NodeID(n.ID),
		Type: nodeType,
	}
	for _, to := range n.Links {
		m.Graph.AddEdge(node.NodeID(n.ID), node.NodeID(to))
	}
}

func loadMap() {
	data, err := os.ReadFile("data/map.json")
	if err != nil {
		panic(err)
	}

	var cfg config.MapConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	loadNode(slgWorld, cfg.Capital, node.NodeCapital)

	for _, n := range cfg.Large {
		loadNode(slgWorld, n, node.NodeLarge)
	}
	for _, n := range cfg.Medium {
		loadNode(slgWorld, n, node.NodeMiddle)
	}
	for _, n := range cfg.HQ {
		loadNode(slgWorld, n, node.NodeBase)
	}
	for _, n := range cfg.Small {
		loadNode(slgWorld, n, node.NodeSmall)
	}
}

func loadPlayer() {

}

func InitWorld() {
	slgWorld = world.NewWorld()
	loadMap()
}
