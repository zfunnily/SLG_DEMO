package slg

type NodeType int

const (
	NodeBase    NodeType = 1 // 大本营
	NodeSmall   NodeType = 2 // 小城市
	NodeMiddle  NodeType = 3 // 中城市
	NodeLarge   NodeType = 4 // 大城市
	NodeCapital NodeType = 5 // 都城
)

type NodeInfo struct {
	ID       NodeID
	Type     NodeType
	OwnerUID int64
}
