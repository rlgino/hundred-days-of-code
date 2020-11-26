package interfaces

// Operation is the operation interface
type Operation interface {
	Calculate(n1, n2 int) int
}

// Abstraction it's the abstraction of the implementation
type Abstraction interface {
	Calculate(n1, n2 int) int
}
