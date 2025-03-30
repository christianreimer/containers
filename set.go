package containers

// Set represents a generic set data structure that can store comparable values.
type Set[T comparable] map[T]any

// NewSet creates a new Set with the given items.
// It returns an empty Set if no items are provided.
func NewSet[T comparable](items ...T) Set[T] {
	set := make(Set[T])
	for _, item := range items {
		set[item] = struct{}{}
	}
	return set
}

// Add adds an item to the Set.
func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

// Remove removes an item from the Set.
// If the item doesn't exist, the operation is a no-op.
func (s Set[T]) Remove(item T) {
	delete(s, item)
}

// Contains checks if an item exists in the Set.
// Returns true if the item is present, false otherwise.
func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

// Size returns the number of items in the Set.
func (s Set[T]) Size() int {
	return len(s)
}

// Union creates a new Set containing all items from both sets.
// The result includes items that exist in either the current Set or the other Set.
func (s Set[T]) Union(other Set[T]) Set[T] {
	result := NewSet[T]()
	for item := range s {
		result.Add(item)
	}
	for item := range other {
		result.Add(item)
	}
	return result
}

// Intersection creates a new Set containing only items that exist in both sets.
// The result includes items that exist in both the current Set and the other Set.
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := NewSet[T]()
	for item := range s {
		if other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

// Difference creates a new Set containing items that exist in the current Set but not in the other Set.
// The result includes items that are unique to the current Set.
func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := NewSet[T]()
	for item := range s {
		if !other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

// IsSubset checks if all items in the current Set exist in the other Set.
// Returns true if the current Set is a subset of the other Set, false otherwise.
func (s Set[T]) IsSubset(other Set[T]) bool {
	for item := range s {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// IsSuperset checks if all items in the other Set exist in the current Set.
// Returns true if the current Set is a superset of the other Set, false otherwise.
func (s Set[T]) IsSuperset(other Set[T]) bool {
	for item := range other {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}
