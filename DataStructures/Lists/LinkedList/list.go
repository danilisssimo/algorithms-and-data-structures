package linkedlist

import (
	"errors"
	"fmt"
)

type ListNode struct {
	Data     interface{}
	NextLink *ListNode
}

type List struct {
	HeadLink *ListNode
	TailLink *ListNode
	length   int
}

// May be wrong
func (list *List) Append(value interface{}) {
	node := &ListNode{
		Data: value, NextLink: nil,
	}
	if list.HeadLink == nil {
		list.HeadLink = node
		list.TailLink = node
	} else {
		list.TailLink.NextLink = node
		list.TailLink = node
	}
	list.length++
}

// May be wrong
func (list *List) Prepend(value interface{}) {
	node := &ListNode{Data: value, NextLink: list.HeadLink}
	if list.HeadLink == nil {
		list.HeadLink = node
		list.TailLink = node
	} else {
		list.TailLink.NextLink = node
	}
	list.length++
}

func (list *List) RemoveFirst() (*ListNode, error) {
	if list.length == 0 {
		return nil, errors.New("list is empty")
	}
	linkToNode := list.HeadLink
	list.HeadLink = list.HeadLink.NextLink
	return linkToNode, nil
}

func (list *List) Remove(position int) (*ListNode, error) {
	if list.length <= position {
		return nil, errors.New("no element in this position")
	}
	var node *ListNode = list.HeadLink
	var currentNodeIndex int = 0
	for currentNodeIndex != position-1 {
		node = node.NextLink
		currentNodeIndex++
	}
	var nextNodeLink *ListNode = nil

	removedNode := node.NextLink // save removed node

	if removedNode.NextLink != nil { // if removed node not Tail
		nextNodeLink = removedNode.NextLink
		removedNode.NextLink = nil // delete link to next node
	}
	node.NextLink = nextNodeLink // previous node to delete
	return removedNode, nil
}

func (list *List) GetLength() int {
	return list.length
}

func (list *List) Draw() {
	var currentNode *ListNode = list.HeadLink
	fmt.Printf("Length: %d;\nList: ", list.length)
	for i := 0; i < list.length; i++ {
		fmt.Printf("%+v -> ", currentNode.Data)
		currentNode = currentNode.NextLink
	}
	fmt.Printf("nil;\n")
}

func GetNewList() *List {
	return &List{
		HeadLink: nil,
		TailLink: nil,
		length:   0,
	}
}
