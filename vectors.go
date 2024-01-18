package mathlib

import (
	"fmt"
	"math"
	"strings"
)

// Vector define a vector
type Vector struct {
	X float64
	Y float64
	Z float64
}

// Module return a module from vector
func (v Vector) Module() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}

// Add perform the addition between a and b
func (v Vector) Add(u Vector) (r Vector) {
	r.X = v.X + u.X
	r.Y = v.Y + u.Y
	r.Z = v.Z + u.Z
	return
}

// Sub perform the subtraction between a and b
func (v Vector) Sub(u Vector) (r Vector) {
	r.X = v.X - u.X
	r.Y = v.Y - u.Y
	r.Z = v.Z - u.Z
	return
}

// Scalar perform the scalar product between a and b
func (v Vector) Scalar(u Vector) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

// Cross perform the cross product between a and b
func (v Vector) Cross(u Vector) (r Vector) {
	r.X = v.Y*u.Z - v.Z*u.Y
	r.Y = v.Z*u.X - v.X*u.Z
	r.Z = v.X*u.Y - v.Y*u.X
	return
}

// String redefine string() to Vector
func (v Vector) String() string {
	strOut := ""
	if v.X != 0 {
		strOut += fmt.Sprintf("%si", cleanZeros(v.X))
	}
	if v.Y != 0 {
		if len(strOut) > 0 {
			strOut += fmt.Sprintf(" + %sj", cleanZeros(v.Y))
		} else {
			strOut += fmt.Sprintf("%sj", cleanZeros(v.Y))
		}
	}
	if v.Z != 0 {
		if len(strOut) > 0 {
			strOut += fmt.Sprintf(" + %sk", cleanZeros(v.Z))
		} else {
			strOut += fmt.Sprintf("%sk", cleanZeros(v.Z))
		}
	}
	if len(strOut) == 0 {
		return "0"
	}
	return strOut
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
