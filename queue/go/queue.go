package queue

import "fmt"

type Queue struct {
	items []int
}

func NewQueue() *Queue {
	return &Queue{
		items: []int{},
	}
}

func (q *Queue) Enqueue(val int) {
	q.items = append(q.items, val)
}

func (q *Queue) Dequeue() (int, error) {

	if len(q.items) <= 0 {
		return -1, fmt.Errorf("Can't Dequeue, queue is empty!")
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item, nil
}

func (q *Queue) Peek() (int, error) {
	if len(q.items) <= 0 {
		return -1, fmt.Errorf("Can't peek, queue is empty!")
	}

	return q.items[0], nil
}
