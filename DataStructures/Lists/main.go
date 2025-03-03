package main

import (
	linkedlist "algorithms-and-data-structures/DataStructures/Lists/LinkedList"
	"fmt"
)

func main() {
	testForLinkedList()
	testPrependForLinkedList()
	testRemoveForLinkedList()
}

func testForLinkedList() {
	var newList *linkedlist.List = linkedlist.CreateLinkedList([]interface{}{})
	newList.Append("Hello")
	newList.Append("World")
	newList.Append("!")
	newList.Append(677)
	newList.Draw()
}

func testPrependForLinkedList() {
	var newList *linkedlist.List = linkedlist.CreateLinkedList([]interface{}{})
	newList.Append("Hello")
	newList.Prepend("World")
	newList.Append("!")
	newList.Prepend(677)
	newList.Draw()
}

func testRemoveForLinkedList() {
	var newList *linkedlist.List = linkedlist.List1
	newList.Draw()
	fmt.Println("Delete 1")
	newList.RemoveFirst()
	newList.Draw()
	fmt.Println("Delete 4")
	newList.RemoveFirst()
	newList.RemoveFirst()
	newList.RemoveFirst()
	newList.RemoveFirst()
	newList.Draw()
}
