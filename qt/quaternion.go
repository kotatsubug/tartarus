package qt

import (
	"math"
)

type Quat struct {
	W, X, Y, Z float32
}

// Get versor of q.
func (q Quat) Normalize() Quat {
	norm := q.Norm()
	return Quat{q.W/norm, q.X/norm, q.Y/norm, q.Z/norm}
}

func (q Quat) Norm() float32 {
	return float32(math.Sqrt(float64(q.W*q.W + q.X*q.X + q.Y*q.Y + q.Z*q.Z)))
}

func (q Quat) Conjugate() Quat {
	return Quat{q.W, -q.X, -q.Y, -q.Z}
}

func NewVersor(x, y, z float32) Quat {
	t := 1.0 - x*x - y*y - z*z
	var w float32
	if t < 0.0 {
		w = 0.0
	} else {
		w = -float32(math.Sqrt(float64(t)))
	}

	return Quat{w, x, y, z}.Normalize()
}





// * q must be normalized in order to provide a functional rotation matrix.
func QuatToMat3(q Quat) Mat3 {
	// Q = w + ix + jy + kz
	//      || 1 - 2yy - 2zz  2xy - 2zw      2xz + 2yw     ||
	//  R = || 2xy + 2zw      1 - 2xx - 2zz  2yz - 2xw     ||
	//      || 2xz - 2yw      2yz + 2xw      1 - 2xx - 2yy ||
	return Mat3{
		1.0 - 2.0*q.Y*q.Y - 2.0*q.Z*q.Z,
		2.0*q.X*q.Y + 2.0*q.Z*q.W,
		2.0*q.X*q.Z - 2.0*q.Y*q.W,
		2.0*q.X*q.Y - 2.0*q.Z*q.W,
		1.0 - 2.0*q.X*q.X - 2.0*q.Z*q.Z,
		2.0*q.Y*q.Z + 2.0*q.X*q.W,
		2.0*q.X*q.Z + 2.0*q.Y*q.W,
		2.0*q.Y*q.Z - 2.0*q.X*q.W,
		1.0 - 2.0*q.X*q.X - 2.0*q.Y*q.Y}
}