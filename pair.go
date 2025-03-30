package containers

import (
	"golang.org/x/exp/constraints"
)

// ComparablePair represents a pair of comparable values
type ComparablePair[T constraints.Ordered, U constraints.Ordered] struct {
	First  T
	Second U
}

// Equal checks if two pairs are equal
func (p ComparablePair[T, U]) Equal(other ComparablePair[T, U]) bool {
	return p.First == other.First && p.Second == other.Second
}

// Less checks if the current pair is less than another pair
// It compares First values first, then Second values if First values are equal
func (p ComparablePair[T, U]) Less(other ComparablePair[T, U]) bool {
	if p.First != other.First {
		return p.First < other.First
	}
	return p.Second < other.Second
}
