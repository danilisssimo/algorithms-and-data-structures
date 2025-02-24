package main

import (
	fifo "algorithms-and-data-structures/DataStructures/Stack/FIFO"
	"fmt"
)

func main() {
	testForFIFO()
}

func testForFIFO() {
	var sliceOfInt []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	var queue *fifo.Queue = fifo.GetNewQueue(10)
	for _, el := range sliceOfInt {
		queue.Enqueue(el)
	}
	queue.Draw()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	queue.Enqueue(646)
	queue.Enqueue(642)
	queue.Enqueue(642)
	queue.Draw()
}
