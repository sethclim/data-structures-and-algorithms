package stack

import "testing"

func TestPush(t *testing.T) {
	s := NewStack()

	s.Push(1)
	s.Push(2)

	if s.items[0] != 1 {
		t.Fatal("Expected first item to be 1")
	}

	if s.items[1] != 2 {
		t.Fatal("Expected first item to be 2")
	}
}

func TestPop(t *testing.T) {
	s := NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	res, err := s.Pop()

	if err != nil {
		t.Fatal("unexpected error %w", err)
	}

	if res != 3 {
		t.Fatal("Expected popped item to be 3")
	}

	res2, err := s.Pop()

	if err != nil {
		t.Fatal("unexpected error %w", err)
	}

	if res2 != 2 {
		t.Fatal("Expected popped item to be 2")
	}
}
