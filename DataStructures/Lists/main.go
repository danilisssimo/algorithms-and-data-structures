package main

import (
	linkedlist "algorithms-and-data-structures/DataStructures/Lists/LinkedList"
	"fmt"
)

func main() {
	testForLinkedList()
	testPrependForLinkedList()
	testRemoveFirstForLinkedList()
	fmt.Println("--TEST FOR REMOVING--")
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

func testRemoveFirstForLinkedList() {
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

func testRemoveForLinkedList() {
	var newList *linkedlist.List = linkedlist.List2
	newList.Draw()
	fmt.Println("Delete last")
	newList.Remove(7)
	newList.Draw()
	fmt.Println("Delete 3th")
	newList.Remove(3)
	newList.Draw()
	fmt.Println("Delete first")
	newList.Remove(0)
	newList.Draw()
	newList = linkedlist.List3
	newList.Draw()
	newList.Remove(0)
	newList.Draw()

	newList = linkedlist.List4
	newList.Draw()
	_, err := newList.Remove(3)
	if err != nil {
		fmt.Println(err)
	}
}
