package qt

import (
	"math"
)

type Vec2 [2]float32
type Vec3 [3]float32

func (a Vec3) Add(b Vec3) Vec3 {
	return Vec3{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

func (a Vec3) Sub(b Vec3) Vec3 {
	return Vec3{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func (a Vec3) Mul(c float32) Vec3 {
	return Vec3{a[0] * c, a[1] * c, a[2] * c}
}

func (a Vec3) MulV(b Vec3) Vec3 {
	return Vec3{a[0] * b[0], a[1] * b[1], a[2] * b[2]}
}

func (a Vec3) Cross(b Vec3) Vec3 {
	return Vec3{a[1]*b[2] - a[2]*b[1], a[2]*b[0] - a[0]*b[2], a[0]*b[1] - a[1]*b[0]}
}

func (v Vec3) Normalize() Vec3 {
	f := 1.0 / float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])))
	return v.Mul(f)
}