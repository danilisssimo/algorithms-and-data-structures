package main

import (
	linkedlist "algorithms-and-data-structures/DataStructures/Lists/LinkedList"
)

func main() {
	testForLinkedList()
}

func testForLinkedList() {
	var newList *linkedlist.List = linkedlist.CreateLinkedList([]interface{}{})
	newList.Append("Hello")
	newList.Append("World")
	newList.Append("!")
	newList.Append(677)
	newList.Draw()
}
