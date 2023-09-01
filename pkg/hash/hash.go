package hash

import "crypto/sha256"

func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}
