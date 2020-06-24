package aaaaaaab

// Function RemoveLastByte () removes the last byte in a slice passed to it. If an empty slice
// is passed, an empty slice also gets returned.
func RemoveLastByte (i []byte) ([]byte) {
	newNoOfBytes := len (i) - 1
	if newNoOfBytes < 0 {
		newNoOfBytes = 0
	}
	
	newSlice := make ([]byte, newNoOfBytes)
	copy (newSlice, i)
	return newSlice
}
