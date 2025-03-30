# Go Generic Containers

This package provides generic container data structures for Go, including Set and ComparablePair implementations.

## Set

A generic Set implementation that can store any comparable type. The Set provides standard set operations and relationships.

### Features

- **Basic Operations**
  - `NewSet[T](...items)`: Create a new set with optional initial items
  - `Add(item)`: Add an item to the set
  - `Remove(item)`: Remove an item from the set
  - `Contains(item)`: Check if an item exists in the set
  - `Size()`: Get the number of items in the set

- **Set Operations**
  - `Union(other)`: Create a new set containing all items from both sets
  - `Intersection(other)`: Create a new set containing only items common to both sets
  - `Difference(other)`: Create a new set containing items in the current set but not in the other set

- **Set Relationships**
  - `IsSubset(other)`: Check if the current set is a subset of another set
  - `IsSuperset(other)`: Check if the current set is a superset of another set

### Example

```go
set1 := NewSet[int](1, 2, 3)
set2 := NewSet[int](3, 4, 5)

union := set1.Union(set2)        // Contains: 1, 2, 3, 4, 5
intersection := set1.Intersection(set2)  // Contains: 3
difference := set1.Difference(set2)      // Contains: 1, 2
```

## ComparablePair

A generic pair structure that can hold two ordered comparable values. Useful for storing and comparing pairs of values.

### Features

- Supports any two types that satisfy the `constraints.Ordered` interface
- Provides comparison operations between pairs
- Type-safe implementation using Go generics

### Methods

- `Equal(other)`: Check if two pairs are equal
- `Less(other)`: Compare pairs lexicographically (First values first, then Second values)

### Example

```go
// Create pairs with different types
intPair := ComparablePair[int, int]{1, 2}
stringPair := ComparablePair[string, int]{"test", 1}

// Compare pairs
result := intPair.Less(ComparablePair[int, int]{1, 3})  // true
equal := stringPair.Equal(ComparablePair[string, int]{"test", 1})  // true
```

## Installation

```bash
go get github.com/yourusername/containers
```

## Usage

```go
import "github.com/yourusername/containers"
```

## Testing

The package includes comprehensive test coverage. Run the tests using:

```bash
go test ./containers
```
```

This README provides:
1. An overview of both data structures
2. Detailed feature lists
3. Usage examples
4. Installation and testing instructions
5. Clear documentation of the API

Let me know if you'd like me to add or modify any sections!