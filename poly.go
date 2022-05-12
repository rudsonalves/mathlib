package mathlib

import (
	"fmt"
	"math"
)

// SolvePoly2: returns the real roots of the second order polynomial
func SolvePoly2(a, b, c float64) (x1, x2 float64, err error) {
	d := delta2(a, b, c)
	if d < 0 {
		return 0, 0, fmt.Errorf("there are no real roots")
	}

	rd := math.Sqrt(d)
	x1 = (-b + rd) / (2 * a)
	x2 = (-b - rd) / (2 * a)
	return
}

// delta2: internal function
func delta2(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}
