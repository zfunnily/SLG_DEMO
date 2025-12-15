package slg

type NodeID int64

type Graph struct {
	edges map[NodeID]map[NodeID]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		edges: make(map[NodeID]map[NodeID]struct{}),
	}
}

func (g *Graph) AddNode(id NodeID) {
	if _, ok := g.edges[id]; !ok {
		g.edges[id] = make(map[NodeID]struct{})
	}
}

func (g *Graph) AddEdge(a, b NodeID) {
	g.AddNode(a)
	g.AddNode(b)
	g.edges[a][b] = struct{}{}
	g.edges[b][a] = struct{}{}
}

func (g *Graph) Neighbors(id NodeID) []NodeID {
	ns := make([]NodeID, 0)
	for n := range g.edges[id] {
		ns = append(ns, n)
	}
	return ns
}
