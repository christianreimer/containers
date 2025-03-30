package containers

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	// Test empty set
	emptySet := NewSet[int]()
	if emptySet.Size() != 0 {
		t.Error("NewSet() should create an empty set")
	}

	// Test set with items
	set := NewSet(1, 2, 3)
	if set.Size() != 3 {
		t.Error("NewSet(items...) should create a set with the given items")
	}
	if !set.Contains(1) || !set.Contains(2) || !set.Contains(3) {
		t.Error("NewSet(items...) should contain all provided items")
	}
}

func TestAdd(t *testing.T) {
	set := NewSet[string]()
	set.Add("test")
	if !set.Contains("test") {
		t.Error("Add() should add the item to the set")
	}
	if set.Size() != 1 {
		t.Error("Add() should increase the set size")
	}
}

func TestRemove(t *testing.T) {
	set := NewSet(1, 2, 3)
	set.Remove(2)
	if set.Contains(2) {
		t.Error("Remove() should remove the item from the set")
	}
	if set.Size() != 2 {
		t.Error("Remove() should decrease the set size")
	}
	// Test removing non-existent item
	set.Remove(4)
	if set.Size() != 2 {
		t.Error("Remove() of non-existent item should not change set size")
	}
}

func TestContains(t *testing.T) {
	set := NewSet(1, 2, 3)
	if !set.Contains(1) {
		t.Error("Contains() should return true for existing items")
	}
	if set.Contains(4) {
		t.Error("Contains() should return false for non-existing items")
	}
}

func TestSize(t *testing.T) {
	set := NewSet(1, 2, 3)
	if set.Size() != 3 {
		t.Error("Size() should return the correct number of items")
	}
	set.Remove(1)
	if set.Size() != 2 {
		t.Error("Size() should update after removing items")
	}
}

func TestUnion(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(3, 4, 5)
	union := set1.Union(set2)

	expected := NewSet[int](1, 2, 3, 4, 5)
	if union.Size() != expected.Size() {
		t.Error("Union() should have the correct size")
	}
	for item := range expected {
		if !union.Contains(item) {
			t.Errorf("Union() should contain item %v", item)
		}
	}
}

func TestIntersection(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(3, 4, 5)
	intersection := set1.Intersection(set2)

	if intersection.Size() != 1 {
		t.Error("Intersection() should have the correct size")
	}
	if !intersection.Contains(3) {
		t.Error("Intersection() should contain common items")
	}
}

func TestDifference(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(3, 4, 5)
	difference := set1.Difference(set2)

	if difference.Size() != 2 {
		t.Error("Difference() should have the correct size")
	}
	if !difference.Contains(1) || !difference.Contains(2) {
		t.Error("Difference() should contain items unique to the first set")
	}
	if difference.Contains(3) {
		t.Error("Difference() should not contain items from the second set")
	}
}

func TestIsSubset(t *testing.T) {
	set1 := NewSet(1, 2)
	set2 := NewSet(1, 2, 3)
	set3 := NewSet(1, 3, 4)

	if !set1.IsSubset(set2) {
		t.Error("IsSubset() should return true when all items are contained")
	}
	if set1.IsSubset(set3) {
		t.Error("IsSubset() should return false when items are not contained")
	}
}

func TestIsSuperset(t *testing.T) {
	set1 := NewSet[int](1, 2, 3)
	set2 := NewSet[int](1, 2)
	set3 := NewSet[int](1, 2, 4)

	if !set1.IsSuperset(set2) {
		t.Error("IsSuperset() should return true when all items are contained")
	}
	if set1.IsSuperset(set3) {
		t.Error("IsSuperset() should return false when items are not contained")
	}
}

func TestPrint(t *testing.T) {
	// Test empty set
	emptySet := NewSet[int]()
	if emptySet.Print() != "Set{}" {
		t.Errorf("Print() for empty set should be 'Set{}', got '%s'", emptySet.Print())
	}

	// Test single item
	singleSet := NewSet(1)
	if singleSet.Print() != "Set{1}" {
		t.Errorf("Print() for single item should be 'Set{1}', got '%s'", singleSet.Print())
	}

	// Test multiple items
	// Note: Since maps don't guarantee order, we need to check for all possible valid outputs
	multiSet := NewSet(1, 2)
	result := multiSet.Print()
	if result != "Set{1, 2}" && result != "Set{2, 1}" {
		t.Errorf("Print() for multiple items should be either 'Set{1, 2}' or 'Set{2, 1}', got '%s'", result)
	}

	// Test with strings
	strSet := NewSet("a", "b")
	result = strSet.Print()
	if result != "Set{a, b}" && result != "Set{b, a}" {
		t.Errorf("Print() for strings should be either 'Set{a, b}' or 'Set{b, a}', got '%s'", result)
	}
}
