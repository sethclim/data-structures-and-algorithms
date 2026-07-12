package singlyLinkedList

import (
	"testing"
)

func TestInsertFront(t *testing.T) {
	list := LinkedList{}

	list.InsertFront(1)
	list.InsertFront(2)

	if list.Head.Val != 2 {
		t.Fatal("expected head value 2")
	}

	if list.Head.Next.Val != 1 {
		t.Fatal("expected second node value 1")
	}
}

func TestInsertEnd(t *testing.T) {
	list := LinkedList{}

	list.InsertEnd(1)
	list.InsertEnd(2)

	if list.Head.Val != 1 {
		t.Fatal("expected head value 1")
	}

	if list.Head.Next.Val != 2 {
		t.Fatal("expected second node value 2")
	}
}

func TestSearch(t *testing.T) {
	list := LinkedList{}

	list.InsertFront(1)
	list.InsertFront(2)

	expect := list.Head

	res := list.Search(2)

	if expect != res {
		t.Fatal("expected to find 2nd node in list")
	}

	res2 := list.Search(5)

	if res2 != nil {
		t.Fatal("expected nil")
	}
}

func TestDelete(t *testing.T) {
	list := LinkedList{}

	list.InsertEnd(1)
	list.InsertEnd(2)
	list.InsertEnd(3)

	list.Delete(2)

	if list.Head.Next.Val != 3 {
		t.Fatal("expected first node to point to last")
	}
}
