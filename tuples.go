package containers

type Pair[T any, U any] struct {
	First  T
	Second U
}

type Triple[T any, U any, V any] struct {
	First  T
	Second U
	Third  V
}

type Quad[T any, U any, V any, W any] struct {
	First  T
	Second U
	Third  V
	Fourth W
}
