package mathlib

import (
	"fmt"
	"math"
	"strings"
)

// Vector: define a vector
type Vector struct {
	X float64
	Y float64
	Z float64
}

// Module: return a module from vector
func (v Vector) Module() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}

// Add: perform the addition between a and b
func (a Vector) Add(b Vector) (r Vector) {
	r.X = a.X + b.X
	r.Y = a.Y + b.Y
	r.Z = a.Z + b.Z
	return
}

// Sub: perform the subtraction between a and b
func (a Vector) Sub(b Vector) (r Vector) {
	r.X = a.X - b.X
	r.Y = a.Y - b.Y
	r.Z = a.Z - b.Z
	return
}

// Scalar: perform the scalar product between a and b
func (a Vector) Scalar(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Cross: perform the cross product between a and b
func (a Vector) Cross(b Vector) (r Vector) {
	r.X = a.Y*b.Z - a.Z*b.Y
	r.Y = a.Z*b.X - a.X*b.Z
	r.Z = a.X*b.Y - a.Y*b.X
	return
}

// String - redefine string() to Vector
func (r Vector) String() string {
	str_out := ""
	if r.X != 0 {
		str_out += fmt.Sprintf("%si", cleanZeros(r.X))
	}
	if r.Y != 0 {
		if len(str_out) > 0 {
			str_out += fmt.Sprintf(" + %sj", cleanZeros(r.Y))
		} else {
			str_out += fmt.Sprintf("%sj", cleanZeros(r.Y))
		}
	}
	if r.Z != 0 {
		if len(str_out) > 0 {
			str_out += fmt.Sprintf(" + %sk", cleanZeros(r.Z))
		} else {
			str_out += fmt.Sprintf("%sk", cleanZeros(r.Z))
		}
	}
	if len(str_out) == 0 {
		return "0"
	}
	return str_out
}

func cleanZeros(fstr float64) string {
	str := fmt.Sprintf("%f", fstr)
	index := strings.Index(str, ".")
	if index < 0 {
		return str
	}
	for i := len(str) - 1; i > index; i-- {
		if str[i] != 48 {
			index = i + 1
			break
		}
	}
	return str[:index]
}
