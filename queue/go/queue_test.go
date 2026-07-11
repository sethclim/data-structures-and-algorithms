package queue

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)

	if len(q.items) != 2 {
		t.Fatal("expected 2 items")
	}

	if q.items[0] != 1 {
		t.Fatal("expected 1 first")
	}

	if q.items[1] != 2 {
		t.Fatal("expected 2 second")
	}

}

func TestDequeue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)

	out1, err := q.Dequeue()

	if err != nil {
		t.Fatal("Unexpected Error %w", err)
	}

	if out1 != 1 {
		t.Fatal("expected 1")
	}

	out2, err := q.Dequeue()

	if err != nil {
		t.Fatal("Unexpected Error %w", err)
	}
	if out2 != 2 {
		t.Fatal("expected 2")
	}

}

func TestPeek(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)

	res, err := q.Peek()

	if err != nil {
		t.Fatal("Unexpected Error %w", err)
	}

	if res != 1 {
		t.Fatal("expected 1")
	}
}
