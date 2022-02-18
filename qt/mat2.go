package qt

type Mat2 [4]float32

func Identity2() Mat2 {
	return Mat2{1,0,0,1}
}

func (m *Mat2) At(col, row int) *float32 {
	return &m[2 * col + row]
}

func (m Mat2) Trace() float32 {
	return (m[0] + m[3])
}

func (m Mat2) Determinant() float32 {
	return (m[0]*m[3] - m[2]*m[1])
}