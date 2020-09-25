package datastructures

import (
	"fmt"
)

// Dequeue is a double ended queue
type Dequeue struct {
	Queue
}

// EnqueueFront add a item to the front of the DeQueue
func (d *Dequeue) EnqueueFront(item string) {
	var name Queue

	name.Enqueue(item)

	newQueue := append(name, d.Queue...)

	d.Queue = newQueue
}

// DequeueEnd add a item to the end of the DeQueue
func (d *Dequeue) DequeueEnd() (string, error) {

	if d.Queue.IsEmpty() {
		err := fmt.Errorf("de-queue is empty")
		return "", err
	}

	item := d.Queue[len(d.Queue)-1]
	d.Queue = d.Queue[:len(d.Queue)-1]

	return item, nil
}
