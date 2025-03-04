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

func (list *List) Prepend(value interface{}) {
	node := &ListNode{Data: value, NextLink: list.HeadLink}
	if list.HeadLink == nil {
		list.HeadLink = node
		list.TailLink = node
	} else {
		node.NextLink = list.HeadLink
		list.HeadLink = node
	}
	list.length++
}

func (list *List) Find(value interface{}) (*ListNode, bool) {
	if list.TailLink.Data == value {
		return list.TailLink, true
	}
	var node *ListNode = list.HeadLink
	var currentNodeIndex int = 0
	for currentNodeIndex <= list.length-1 {
		if node.Data == value {
			return node, true
		}
		node = node.NextLink
		currentNodeIndex++
	}
	return nil, false
}

func (list *List) InsertAt(position int, value interface{}) error {
	if list.length <= position || 0 > position {
		return fmt.Errorf("position %d is outside of this list", position)
	}
	if position == 0 {
		list.Prepend(value)
		return nil
	}
	if position == list.length-1 {
		list.Append(value)
		return nil
	}
	var newNode *ListNode = &ListNode{Data: value}
	var node *ListNode = list.HeadLink
	var currentNodeIndex int = 0
	for currentNodeIndex != position-1 {
		node = node.NextLink
		currentNodeIndex++
	}
	nextNode := node.NextLink
	node.NextLink = newNode
	newNode.NextLink = nextNode
	list.length++
	return nil
}

func (list *List) GetAt(position int) (*ListNode, error) {
	if list.length <= position || 0 > position {
		return nil, errors.New("no element in this position")
	}
	if position == 0 {
		return list.HeadLink, nil
	}
	if position == list.length-1 {
		return list.TailLink, nil
	}
	var node *ListNode = list.HeadLink
	var currentNodeIndex int = 0
	for currentNodeIndex != position {
		node = node.NextLink
		currentNodeIndex++
	}
	return node, nil
}

func (list *List) Remove(position int) (*ListNode, error) {
	if list.length <= position {
		return nil, errors.New("no element in this position")
	}
	if position == 0 {
		return list.RemoveFirst()
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
	list.length--
	return removedNode, nil
}

func (list *List) RemoveFirst() (*ListNode, error) {
	if list.length == 0 {
		return nil, errors.New("list is empty")
	}
	linkToNode := list.HeadLink
	list.HeadLink = list.HeadLink.NextLink
	list.length--
	return linkToNode, nil
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
