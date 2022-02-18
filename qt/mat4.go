package qt

import (
	"math"
)

type Mat4 [16]float32 // column major

func Identity4() Mat4 {
	return Mat4{1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1}
}

func (m *Mat4) At(col, row int) *float32 {
	return &m[4 * col + row]
}

func (m1 Mat4) Add(m2 Mat4) Mat4 {
	return Mat4{
		m1[0] + m2[0],
		m1[1] + m2[1],
		m1[2] + m2[2],
		m1[3] + m2[3],
		m1[4] + m2[4],
		m1[5] + m2[5],
		m1[6] + m2[6],
		m1[7] + m2[7],
		m1[8] + m2[8],
		m1[9] + m2[9],
		m1[10] + m2[10],
		m1[11] + m2[11],
		m1[12] + m2[12],
		m1[13] + m2[13],
		m1[14] + m2[14],
		m1[15] + m2[15]}
}

func (m1 Mat4) Mul4(m2 Mat4) Mat4 {
	return Mat4{
		m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3],
		m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3],
		m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3],
		m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3],
		m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7],
		m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7],
		m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7],
		m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7],
		m1[0]*m2[8] + m1[4]*m2[9] + m1[8]*m2[10] + m1[12]*m2[11],
		m1[1]*m2[8] + m1[5]*m2[9] + m1[9]*m2[10] + m1[13]*m2[11],
		m1[2]*m2[8] + m1[6]*m2[9] + m1[10]*m2[10] + m1[14]*m2[11],
		m1[3]*m2[8] + m1[7]*m2[9] + m1[11]*m2[10] + m1[15]*m2[11],
		m1[0]*m2[12] + m1[4]*m2[13] + m1[8]*m2[14] + m1[12]*m2[15],
		m1[1]*m2[12] + m1[5]*m2[13] + m1[9]*m2[14] + m1[13]*m2[15],
		m1[2]*m2[12] + m1[6]*m2[13] + m1[10]*m2[14] + m1[14]*m2[15],
		m1[3]*m2[12] + m1[7]*m2[13] + m1[11]*m2[14] + m1[15]*m2[15]}
}

func (m Mat4) Mul(c float32) Mat4 {
	return Mat4{
		m[0] * c,
		m[1] * c,
		m[2] * c,
		m[3] * c,
		m[4] * c,
		m[5] * c,
		m[6] * c,
		m[7] * c,
		m[8] * c,
		m[9] * c,
		m[10] * c,
		m[11] * c,
		m[12] * c,
		m[13] * c,
		m[14] * c,
		m[15] * c}
}

func (m Mat4) Trace() float32 {
	return (m[0] + m[5] + m[10] + m[15])
}

func (m Mat4) Determinant() float32 {
	// Laplace (cofactor) expansion
	var det float32 = 0
	var s Mat3 // Sub matrix to expand by

	for col := 0; col < 4; col++ {
		a, b := 0, 0
		for i := 0; i < 4; i++ {
			if i == 0 {
				continue
			}
			b = 0
			for j := 0; j < 4; j++ {
				if col == j {
					continue
				}
				*s.At(b, a) = *m.At(j, i)
				b++
			}
			a++
		}

		if col % 2 == 0 {
			det += *m.At(col, 0) * s.Determinant()
		} else {
			det -= *m.At(col, 0) * s.Determinant()
		}
	}

	return det
}

// Replace position component of matrix.
func (m Mat4) SetPosition(v Vec3) Mat4 {
	return Mat4{m[0], m[1], m[2], m[3], m[4], m[5], m[6], m[7], m[8], m[9], m[10], m[11], v[0], v[1], v[2], m[15]}
}

// * angle in radians. axis must be normalized.
func AxisAngleToMat3(axis Vec3, angle float32) Mat3 {
	c, s := float32(math.Cos(float64(angle))), float32(math.Sin(float64(angle)))
	t := 1 - c
	x, y, z := axis[0], axis[1], axis[2]

	// [R] = [I] + s*[~axis] + t*[~axis]^2
	//     = c*[I] + s*[~axis] + t*([~axis]^2 + [I])   (~ implies skew symmetry)
	//
	//        || txx+c txy-zs txz+ys ||
	//  [R] = || txy+zs tyy+c tyz-xs ||
	//        || txz-ys tyz+xs tzz+c ||

	return Mat3{
		t*x*x + c,
		t*x*y + z*s,
		t*x*z - y*s,
		t*x*y - z*s,
		t*y*y + c,
		t*y*z + x*s,
		t*x*z + y*s,
		t*y*z - x*s,
		t*z*z + c}
}

// Rotate a 4x4 matrix by a 3x3 rotation matrix about a pivot point.
func (m Mat4) RotateAboutPivot(r Mat3, p Vec3) Mat4 {
	// Matrix representing rotation about a given pivot point P is
	// [R] = [T]^-1 * [R0] * [T]
	// where  [T]^-1 = inverse transform = translation of P to origin (0,0,0) -> [-Px,-Py,-Pz]
	//        [R0]   = rotation about origin -> [some 3x3 rotation matrix]
	//        [T]    = translation of origin to point -> [+Px,+Py,+Pz]
	// 
	//        || r00 r01 r02  Px - Px*r00 - Py*r01 - Pz*r02 ||
	//  [R] = || r10 r11 r12  Py - Px*r10 - Py*r11 - Pz*r12 ||
	//        || r20 r21 r22  Pz - Px*r20 - Py*r21 - Pz*r22 ||
	//        ||  0   0   0                1                ||

	R := Mat4{
		r[0],
		r[1],
		r[2],
		0,
		r[3],
		r[4],
		r[5],
		0,
		r[6],
		r[7],
		r[8],
		0,
		p[0] - p[0]*r[0] - p[1]*r[3] - p[2]*r[6],
		p[1] - p[0]*r[1] - p[1]*r[4] - p[2]*r[7],
		p[2] - p[0]*r[2] - p[1]*r[5] - p[2]*r[8],
		1}
	
	return m.Mul4(R)
}

// Scales matrix M. Preserves position/rotation.
// TODO: SHOULD preserve rotation... Check
func (m Mat4) SetScale(s Vec3) Mat4 {
	i := Mat4{s[0], 0, 0, 0, 0, s[1], 0, 0, 0, 0, s[2], 0, 0, 0, 0, 1}
	return m.Mul4(i)
}

func OrthographicProjectionMat4(xMin, xMax, yMin, yMax, zNear, zFar float32) Mat4 {
	dx, dy, dz := (xMax - xMin), (yMax - yMin), (zFar - zNear)
	return Mat4{float32(2.0 / dx), 0, 0, 0, 0, float32(2.0 / dy), 0, 0, 0, 0, float32(-2.0 / dz), 0, float32(-(xMax + xMin) / dx), float32(-(yMax + yMin) / dy), float32(-(zFar + zNear) / dz), 1}
}

func PerspectiveProjectionMat4(fovy, aspect, zNear, zFar float32) Mat4 {
	dz, f := (zNear - zFar), float32(1.0 / math.Tan(float64(fovy) / 2.0))
	return Mat4{float32(f / aspect), 0, 0, 0, 0, float32(f), 0, 0, 0, 0, float32((zNear + zFar) / dz), -1, 0, 0, float32((2. * zFar * zNear) / dz), 0}
}

// Generates a transformation matrix from world space into the eye space.
// "eye" - location of the camera
// "center" - where it should be looking
// "up" - global Y of camera.
func LookAtV(eye, center, up Vec3) Mat4 {
	f := (center.Sub(eye)).Normalize()
	s := f.Cross(up.Normalize())
	u := (s.Normalize()).Cross(f)
	M := Mat4{s[0], u[0], -f[0], 0, s[1], u[1], -f[1], 0, s[2], u[2], -f[2], 0, 0, 0, 0, 1}
	N := Mat4{1,0,0,0,0,1,0,0,0,0,1,0,float32(-eye[0]),float32(-eye[1]),float32(-eye[2]),1}
	return M.Mul4(N)
}