package node

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
	if _, ok := g.edges[id]; ok {
		return
	}
	g.edges[id] = make(map[NodeID]struct{})
}

func (g *Graph) AddEdge(a, b NodeID) {
	if a == b {
		return
	}

	// 确保节点存在
	g.AddNode(a)
	g.AddNode(b)

	g.edges[a][b] = struct{}{}
	g.edges[b][a] = struct{}{}
}

func (g *Graph) Neighbors(id NodeID) []NodeID {
	neighbors, ok := g.edges[id]
	if !ok {
		return nil
	}

	result := make([]NodeID, 0, len(neighbors))
	for n := range neighbors {
		result = append(result, n)
	}
	return result
}

func (g *Graph) RemoveNode(id NodeID) {
	neighbors, ok := g.edges[id]
	if !ok {
		return
	}

	// 删除其他节点指向它的边
	for n := range neighbors {
		delete(g.edges[n], id)
	}

	delete(g.edges, id)
}
func (g *Graph) HasEdge(a, b NodeID) bool {
	if _, ok := g.edges[a]; !ok {
		return false
	}
	_, ok := g.edges[a][b]
	return ok
}
