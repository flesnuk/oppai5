package vector

import "math"

// Vector2 2D vector with double values
type Vector2 struct {
	X float64
	Y float64
}

// New makes a new Vector2
func New(x, y float64) Vector2 {
	return Vector2{x, y}
}

// Sub subtracts the vector by another vector
func (v Vector2) Sub(other Vector2) Vector2 {
	v.X -= other.X
	v.Y -= other.Y
	return v
}

// Mul multiplies the vector by a value
func (v Vector2) Mul(value float64) Vector2 {
	v.X *= value
	v.Y *= value
	return v
}

// Len returns the length of the vector
func (v Vector2) Len() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Dot is the product of two vectors
func (v Vector2) Dot(other Vector2) float64 {
	return v.X*other.X + v.Y*other.Y
}
