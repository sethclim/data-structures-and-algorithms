package hashSet

import linkedlist "github.com/sethclim/data-structures-and-algorithms/linked-list"

type MyHashSet struct {
	Arr      []*linkedlist.LinkedList
	Capacity int
}

const DefaultCapacity = 800

func NewHashSet(capacity ...int) *MyHashSet {
	cap := DefaultCapacity
	if len(capacity) > 0 {
		cap = capacity[0]
	}

	arr := make([]*linkedlist.LinkedList, cap)
	for i := range arr {
		arr[i] = linkedlist.NewLinkedList()
	}
	return &MyHashSet{
		Arr:      arr,
		Capacity: cap,
	}
}

func (this *MyHashSet) Add(key int) {
	idx := this.hash(key)
	list := this.Arr[idx]

	if list.Head == nil {
		this.Arr[idx].InsertFront(key)
		return
	}

	if list.Search(key) != nil {
		return
	}

	this.Arr[idx].InsertFront(key)
}

func (this *MyHashSet) Remove(key int) {
	idx := this.hash(key)
	this.Arr[idx].Delete(key)
}

func (this *MyHashSet) Contains(key int) bool {
	idx := this.hash(key)

	return this.Arr[idx].Search(key) != nil
}

func (this *MyHashSet) hash(key int) int {
	return key % this.Capacity
}
