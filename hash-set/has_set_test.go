package hashSet

import "testing"

func TestMyHashSet_AddAndContains(t *testing.T) {
	set := NewHashSet()

	keys := []int{10, 20, 30}

	for _, key := range keys {
		set.Add(key)
	}

	for _, key := range keys {
		if !set.Contains(key) {
			t.Errorf("expected set to contain %d", key)
		}
	}

	if set.Contains(100) {
		t.Errorf("expected set to not contain 100")
	}
}

func TestMyHashSet_DuplicateAdd(t *testing.T) {
	set := NewHashSet()

	set.Add(5)
	set.Add(5)

	// Since this is a set, duplicate adds should do nothing
	if !set.Contains(5) {
		t.Errorf("expected set to contain 5")
	}

	// Optional: verify only one node exists in the bucket
	idx := set.hash(5)
	list := set.Arr[idx]

	count := 0
	for node := list.Head; node != nil; node = node.Next {
		count++
	}

	if count != 1 {
		t.Errorf("expected 1 element in bucket, got %d", count)
	}
}

func TestMyHashSet_Remove(t *testing.T) {
	set := NewHashSet()

	set.Add(42)

	if !set.Contains(42) {
		t.Fatalf("expected set to contain 42 before removal")
	}

	set.Remove(42)

	if set.Contains(42) {
		t.Errorf("expected set to not contain 42 after removal")
	}
}

func TestMyHashSet_Collisions(t *testing.T) {
	// Force collisions by using small capacity
	set := NewHashSet(2)

	// Both hash to index 0
	set.Add(2)
	set.Add(4)
	set.Add(6)

	for _, key := range []int{2, 4, 6} {
		if !set.Contains(key) {
			t.Errorf("expected set to contain %d", key)
		}
	}

	set.Remove(4)

	if set.Contains(4) {
		t.Errorf("expected 4 to be removed")
	}

	if !set.Contains(2) || !set.Contains(6) {
		t.Errorf("removing 4 should not remove other colliding keys")
	}
}

func TestNewHashSet_DefaultCapacity(t *testing.T) {
	set := NewHashSet()

	if set.Capacity != DefaultCapacity {
		t.Errorf("expected capacity %d, got %d", DefaultCapacity, set.Capacity)
	}

	if len(set.Arr) != DefaultCapacity {
		t.Errorf("expected arr length %d, got %d", DefaultCapacity, len(set.Arr))
	}
}

func TestNewHashSet_CustomCapacity(t *testing.T) {
	set := NewHashSet(10)

	if set.Capacity != 10 {
		t.Errorf("expected capacity 10, got %d", set.Capacity)
	}

	if len(set.Arr) != 10 {
		t.Errorf("expected arr length 10, got %d", len(set.Arr))
	}
}
