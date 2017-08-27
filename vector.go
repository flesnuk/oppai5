package main

import "math"

// Vector2 2D vector with double values
type Vector2 struct {
	X float64
	Y float64
}

func (v Vector2) sub(other Vector2) Vector2 {
	v.X -= other.X
	v.Y -= other.Y
	return v
}

func (v Vector2) mul(value float64) Vector2 {
	v.X *= value
	v.Y *= value
	return v
}

func (v Vector2) len() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
