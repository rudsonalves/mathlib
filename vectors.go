package mathlib

import "math"

type Vector struct {
	x float64
	y float64
	z float64
}

func (v Vector) Module() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2) + math.Pow(v.z, 2))
}

func (a Vector) Add(b Vector) (r Vector) {
	r.x = a.x + b.x
	r.y = a.y + b.y
	r.z = a.z + b.z
	return
}

func (a Vector) Sub(b Vector) (r Vector) {
	r.x = a.x - b.x
	r.y = a.y - b.y
	r.z = a.z - b.z
	return
}

func (a Vector) Scalar(b Vector) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func (a Vector) Cross(b Vector) (r Vector) {
	r.x = a.y*b.z - a.z*b.y
	r.y = a.z*b.x - a.x*b.z
	r.z = a.x*b.y - a.y*b.z
	return
}
