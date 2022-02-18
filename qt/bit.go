package qt

// Inserts newBit into val at index (LSB = 0)
func InsertBit(val, newBit, index uint64) uint64 {
	x := val
	y := x
	x <<= 1

	if newBit == 1 {
		x |= (1 << index)
	} else {
		x &= ^(1 << index)
	}

	x &= (1 << index)
	y &= (^(1 << index))
	x |= y

	return x
}
