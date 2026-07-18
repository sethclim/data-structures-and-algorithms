package singlyLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
	}
}

func (list *LinkedList) InsertFront(val int) {
	newNode := &ListNode{Val: val}

	if list.Head != nil {
		newNode.Next = list.Head
	}

	list.Head = newNode
}

func (list *LinkedList) InsertEnd(val int) {
	newNode := &ListNode{Val: val}

	if list.Head == nil {
		list.Head = newNode
	} else {
		curr := list.Head
		for curr.Next != nil {
			curr = curr.Next
		}

		curr.Next = newNode
	}
}

func (list *LinkedList) Search(target int) *ListNode {
	curr := list.Head
	for curr != nil {

		if curr.Val == target {
			return curr
		}

		curr = curr.Next
	}

	return nil
}

func (list *LinkedList) Delete(target int) bool {
	var prev *ListNode

	for curr := list.Head; curr != nil; curr = curr.Next {
		if curr.Val == target {
			if prev == nil {
				// deleting head
				list.Head = curr.Next
			} else {
				// deleting middle or tail
				prev.Next = curr.Next
			}

			return true
		}

		prev = curr
	}

	return false
}
