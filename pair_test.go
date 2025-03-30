package containers

import (
	"sort"
	"testing"
)

func TestComparablePairEqual(t *testing.T) {
	// Test with integers
	pair1 := ComparablePair[int, int]{1, 2}
	pair2 := ComparablePair[int, int]{1, 2}
	pair3 := ComparablePair[int, int]{1, 3}
	pair4 := ComparablePair[int, int]{2, 2}

	if !pair1.Equal(pair2) {
		t.Error("Equal() should return true for identical pairs")
	}
	if pair1.Equal(pair3) {
		t.Error("Equal() should return false when Second values differ")
	}
	if pair1.Equal(pair4) {
		t.Error("Equal() should return false when First values differ")
	}

	// Test with different types
	strPair1 := ComparablePair[string, int]{"test", 1}
	strPair2 := ComparablePair[string, int]{"test", 1}
	strPair3 := ComparablePair[string, int]{"test", 2}

	if !strPair1.Equal(strPair2) {
		t.Error("Equal() should work with different types")
	}
	if strPair1.Equal(strPair3) {
		t.Error("Equal() should return false for different Second values with different types")
	}
}

func TestComparablePairLess(t *testing.T) {
	// Test with integers
	pair1 := ComparablePair[int, int]{1, 2}
	pair2 := ComparablePair[int, int]{1, 3}
	pair3 := ComparablePair[int, int]{2, 1}
	pair4 := ComparablePair[int, int]{1, 1}

	if !pair1.Less(pair2) {
		t.Error("Less() should return true when First values are equal and Second value is less")
	}
	if !pair1.Less(pair3) {
		t.Error("Less() should return true when First value is less")
	}
	if pair1.Less(pair4) {
		t.Error("Less() should return false when First values are equal and Second value is greater")
	}
	if pair3.Less(pair1) {
		t.Error("Less() should return false when First value is greater")
	}

	// Test with different types
	strPair1 := ComparablePair[string, int]{"a", 1}
	strPair2 := ComparablePair[string, int]{"a", 2}
	strPair3 := ComparablePair[string, int]{"b", 1}

	if !strPair1.Less(strPair2) {
		t.Error("Less() should work with different types when First values are equal")
	}
	if !strPair1.Less(strPair3) {
		t.Error("Less() should work with different types when First values differ")
	}
	if strPair3.Less(strPair1) {
		t.Error("Less() should return false for greater First values with different types")
	}

	// Test with floating point numbers
	floatPair1 := ComparablePair[float64, float64]{1.0, 2.0}
	floatPair2 := ComparablePair[float64, float64]{1.0, 2.1}
	floatPair3 := ComparablePair[float64, float64]{1.1, 1.0}

	if !floatPair1.Less(floatPair2) {
		t.Error("Less() should work with floating point numbers when First values are equal")
	}
	if !floatPair1.Less(floatPair3) {
		t.Error("Less() should work with floating point numbers when First values differ")
	}
}

func TestComparablePairSort(t *testing.T) {
	// Test with integers
	pair1 := ComparablePair[int, int]{1, 2}
	pair2 := ComparablePair[int, int]{1, 3}
	pair3 := ComparablePair[int, int]{2, 1}

	unsorted := []ComparablePair[int, int]{pair3, pair1, pair2}
	sort.Slice(unsorted, func(i, j int) bool {
		return unsorted[i].Less(unsorted[j])
	})

	if !unsorted[0].Equal(pair1) || !unsorted[1].Equal(pair2) || !unsorted[2].Equal(pair3) {
		t.Error("Sort() should sort pairs correctly")
	}
}
