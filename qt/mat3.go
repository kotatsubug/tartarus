package qt

type Mat3 [9]float32

func Identity3() Mat3 {
	return Mat3{1,0,0,0,1,0,0,0,1}
}

func (m *Mat3) At(col, row int) *float32 {
	return &m[3 * col + row]
}

func (m Mat3) Trace() float32 {
	return (m[0] + m[4] + m[8])
}

func (m Mat3) Determinant() float32 {
	A, B, C := Mat2{m[4],m[5],m[7],m[8]}, Mat2{m[1],m[2],m[7],m[8]}, Mat2{m[1],m[2],m[4],m[5]}
	return (m[0] * A.Determinant() - m[3] * B.Determinant() + m[6] * C.Determinant())
}