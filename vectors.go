package mathlib

import (
	"math"
)

// Vector define a vector
type Vector struct {
	X float64
	Y float64
	Z float64
}

// Module returna a module from vector
func (v Vector) Module() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}

// Add make a addition from vectores a and b
func (a Vector) Add(b Vector) (r Vector) {
	r.X = a.X + b.X
	r.Y = a.Y + b.Y
	r.Z = a.Z + b.Z
	return
}

func (a Vector) Sub(b Vector) (r Vector) {
	r.X = a.X - b.X
	r.Y = a.Y - b.Y
	r.Z = a.Z - b.Z
	return
}

func (a Vector) Scalar(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vector) Cross(b Vector) (r Vector) {
	r.X = a.Y*b.Z - a.Z*b.Y
	r.Y = a.Z*b.X - a.X*b.Z
	r.Z = a.X*b.Y - a.Y*b.Z
	return
}
