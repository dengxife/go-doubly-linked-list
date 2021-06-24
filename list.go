package main

import "reflect"

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

// Head 返回链表的头节点元素
func (l *List) Head() (head *Node) {
	head = l.head
	return
}

// Tail 返回链表的尾节点元素
func (l *List) Tail() (tail *Node) {
	tail = l.tail
	return
}

// Size 返回链表的数量
func (l *List) Size() (size int64) {
	size = l.size
	return
}

// isEmpty 判断链表是否为空
func (l List) isEmpty() bool {
	if l.Size() == 0 {
		return true
	}
	return false
}

// Append 在链表的末尾添加一个节点元素
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
}

// InsertPrev 在链表的头部添加一个节点元素
func (l *List) InsertPrev(data interface{}) {
	node := NewNode(data)

	if l.isEmpty() {
		l.head = node
		l.tail = node
	} else {
		head := l.Head()
		head.prev = node
		node.next = head

		l.head = node
	}

	l.size++
}

// Remove 删除链表中的任意节点
func (l *List) Remove(node *Node) bool {

	// 如果要删除的节点为nil 或者 链表为空 直接返回 false
	if node == nil || l.isEmpty() {
		return false
	}

	currentNode := l.Head()

	for {
		if reflect.DeepEqual(node.data, currentNode.data) {

			// 如果当前是头部节点
			if currentNode.prev == nil {
				currentNode.Next().prev = nil
				l.head = currentNode.Next()
				currentNode.next = nil

				l.size--
				return true
			}

			// 如果当前你是是尾部节点
			if currentNode.next == nil {
				currentNode.Prev().next = nil
				l.tail = currentNode.Prev()
				currentNode.prev = nil

				l.size--
				return true
			}

			// 如果是中间节点
			currentNode.Prev().next = currentNode.next
			currentNode.Next().prev = currentNode.prev
			currentNode.prev = nil
			currentNode.next = nil

			l.size--

			return true
		}

		currentNode = currentNode.Next()
	}

	return false
}
