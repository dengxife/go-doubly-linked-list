package main

// Node 节点的结构体
type Node struct {
	prev *Node
	next *Node
	data interface{}
}

// NewNode 创建一个节点
func NewNode(data interface{}) (node *Node) {
	node = &Node{
		data: data,
	}
	return
}

// Prev 返回节点的前个节点的引用
func (n *Node) Prev() (prev *Node) {
	prev = n.prev
	return
}

// Next 返回节点的后个节点的引用
func (n *Node) Next() (next *Node) {
	next = n.next
	return
}

// GetNode 返回节点的值
func (n *Node) GetNode() (data interface{}) {
	if n == nil {
		return
	}
	data = n.data
	return
}

// isHead 判断当前节点是不是最后的节点
func (n *Node) isHead() bool {
	if n.prev != nil {
		return false
	}
	return true
}

// isTail 判断当前节点是不是最后的节点
func (n Node) isTail() bool {
	if n.next != nil {
		return false
	}
	return true
}
