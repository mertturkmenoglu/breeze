package random

import "crypto/rand"

// Generate n random bytes
func Bytes(n uint32) ([]byte, error) {
	if n < 1 {
		return nil, ErrInvalidBytesCount
	}

	arr := make([]byte, n)
	_, err := rand.Read(arr)

	if err != nil {
		return nil, err
	}

	return arr, nil
}
