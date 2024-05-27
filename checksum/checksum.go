package checksum

import (
	"crypto/sha256"
)

// CreateChecksum computes the double SHA256 and returns the first 4 bytes as the checksum
func CreateChecksum(payload []byte) [4]byte {
	hash1 := sha256.Sum256(payload)
	hash2 := sha256.Sum256(hash1[:])
	return [4]byte{hash2[0], hash2[1], hash2[2], hash2[3]}
}
