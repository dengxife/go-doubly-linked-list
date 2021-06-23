package main

// List 链表结构体
type List struct {
	head *Node
	tail *Node
	size int64
}

// NewList 返回链表
func NewList() (list *List) {
	list = &List{}
	return
}

// Head 返回链表的头节点
func (l *List) Head() (head *Node) {
	head = l.head
	return
}

// Tail 返回链表的尾节点
func (l *List) Tail() (tail *Node) {
	tail = l.tail
	return
}

// Size 返回链表的数量
func (l *List) Size() (size int64) {
	size = l.size
	return
}

// Append 在链表的末尾添加一个节点
func (l *List) Append(data interface{}) {
	node := NewNode(data)

	if l.Size() == 0 {
		l.head = node
		l.tail = node
	} else {
		tail := l.tail
		tail.next = node
		node.prev = tail

		l.tail = node
	}

	l.size++
	return
}
