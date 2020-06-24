package aaaaaaab

func RemoveLastByte (i []byte) ([]byte) {
	newNoOfBytes := len (i) - 1
	if newNoOfBytes < 0 {
		newNoOfBytes = 0
	}
	
	newSlice := make ([]byte, newNoOfBytes)
	copy (newSlice, i)
	return newSlice
}
