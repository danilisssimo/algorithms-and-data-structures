package fifo

import (
	"errors"
	"fmt"
)

type Queue struct {
	elements           []interface{}
	length             int
	nextFreePlaceIndex int
}

// Dequeue delete and return first element of queue
func (queue *Queue) Dequeue() interface{} {
	if queue.length <= 0 {
		return errors.New("Queue is null")
	}
	result := queue.elements[0]
	queue.elements = queue.elements[1:]
	queue.length--
	queue.nextFreePlaceIndex--
	return result
}

// Enqueue append element to last position of queue
func (queue *Queue) Enqueue(el interface{}) {
	if queue.nextFreePlaceIndex >= queue.length {
		queue.length++
		queue.elements = append(queue.elements, el)
	} else {
		queue.elements[queue.nextFreePlaceIndex] = el
	}
	queue.nextFreePlaceIndex++
}

// GetLength returned length of queue
func (queue *Queue) GetLength() int {
	return queue.length
}

// Draw printing all queue
func (queue *Queue) Draw() {
	for _, el := range queue.elements {
		fmt.Printf("%+v ", el)
	}
	fmt.Printf("\n")
}

// GetNewQueue init new queue and return it
func GetNewQueue(size int) *Queue {
	var queue Queue = Queue{
		elements:           make([]interface{}, size),
		length:             size,
		nextFreePlaceIndex: 0,
	}
	return &queue
}
