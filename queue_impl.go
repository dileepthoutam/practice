package main

import (
	"fmt"
)

type Queue struct {
	queue []string
}

func (q *Queue) Enqueue(name string) {
	q.queue = append(q.queue, name)
}

func (q *Queue) Dequeue() error {
	if len(q.queue) > 0 {
		q.queue = q.queue[1:]
		return nil
	}
	return fmt.Errorf("Queue is empty.")
}

func (q *Queue) Front() (string, error) {
	if len(q.queue) > 0 {
		return q.queue[0], nil
	}

	return "", fmt.Errorf("Queue is empty.")
}

func (q *Queue) Size() int {
	return len(q.queue)
}

func (q *Queue) Empty() bool {
	return q.Size() == 0
}
