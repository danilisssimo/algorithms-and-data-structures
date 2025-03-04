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
	fmt.Println("-- TEST FOR FINDING --")
	testFindForLinkedList()
	fmt.Println("-- TEST FOR GET AT --")
	testGetAtForLinkedList()
	fmt.Println("-- TEST FOR INSERT AT --")
	testInsertAtForLinkedList()
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

func testFindForLinkedList() {
	var newList *linkedlist.List = linkedlist.CreateLinkedList([]interface{}{1, 4, 3, 6, 7, 8, 9, 12, 43, 5, 2, 6})
	node, res := newList.Find(12)
	fmt.Printf("Node: %+v, Result: %t\n", node, res)
	node, res = newList.Find(0)
	fmt.Printf("Node: %+v, Result: %t\n", node, res)
}

func testGetAtForLinkedList() {
	var newList *linkedlist.List = linkedlist.CreateLinkedList([]interface{}{1, 4, 3, 6, 7, 8, 9, 12, 43, 5, 2, 6})
	node, err := newList.GetAt(3)
	fmt.Printf("Node: %+v, Err: %+v\n", node, err)
	node, err = newList.GetAt(-1)
	fmt.Printf("Node: %+v, Err: %+v\n", node, err)
	node, err = newList.GetAt(40)
	fmt.Printf("Node: %+v, Err: %+v\n", node, err)
	node, err = newList.GetAt(0)
	fmt.Printf("Node: %+v, Err: %+v\n", node, err)
	node, err = newList.GetAt(11)
	fmt.Printf("Node: %+v, Err: %+v\n", node, err)
}

func testInsertAtForLinkedList() {
	var newList *linkedlist.List = linkedlist.CreateLinkedList([]interface{}{1, 4, 3, 6, 7, 8, 9, 12, 43, 5, 2, 6})
	err := newList.InsertAt(2, "New value")
	if err != nil {
		fmt.Println(err)
	} else {
		newList.Draw()
	}
	err = newList.InsertAt(-1, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		newList.Draw()
	}
	err = newList.InsertAt(50, 99)
	if err != nil {
		fmt.Println(err)
	} else {
		newList.Draw()
	}
}
