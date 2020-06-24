package aaaaaaab

func RemoveLastByte (i []byte) ([]byte) {
	newSlice := make ([]byte, len (i) - 1)
	copy (newSlice, i)
	return newSlice
}
