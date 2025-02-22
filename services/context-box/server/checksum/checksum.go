package checksum

import (
	"bytes"

	"golang.org/x/crypto/blake2b"
)

// CalculateChecksum calculates a Blake2b 256 bit checksum of the passed data
// and returns it as a []byte.
func CalculateChecksum(data string) []byte {
	res := blake2b.Sum256([]byte(data))
	// Creating a slice using an array you can just make a simple slice expression
	return res[:]
}

// CompareChecksums compares checksums passed in as []byte.
func CompareChecksums(ch1, ch2 []byte) bool {
	return bytes.Equal(ch1, ch2)
}
