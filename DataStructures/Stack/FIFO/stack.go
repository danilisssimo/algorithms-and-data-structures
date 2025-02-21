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

func (queue *Queue) Dequeue() interface{} {
	if queue.length <= 0 {
		return errors.New("Queue is null")
	}
	result := queue.elements[0]
	queue.elements = queue.elements[1:]
	queue.length--
	queue.nextFreePlaceIndex--
	fmt.Printf("Length: %d, NextFreePlaceIndex: %d\n", queue.length, queue.nextFreePlaceIndex)
	return result
}

func (queue *Queue) Enqueue(el interface{}) {
	if queue.nextFreePlaceIndex >= queue.length {
		queue.length++
		queue.elements = append(queue.elements, el)
	} else {
		queue.elements[queue.nextFreePlaceIndex] = el
	}
	queue.nextFreePlaceIndex++
	fmt.Printf("Length: %d, NextFreePlaceIndex: %d\n", queue.length, queue.nextFreePlaceIndex)
}

func (queue *Queue) Draw() {
	for _, el := range queue.elements {
		fmt.Printf("%+v ", el)
	}
	fmt.Printf("\n")
}

func GetNewQueue(size int) *Queue {
	var queue Queue = Queue{
		elements:           make([]interface{}, size),
		length:             size,
		nextFreePlaceIndex: 0,
	}
	return &queue
}
