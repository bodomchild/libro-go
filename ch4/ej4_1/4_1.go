package ej4_1

import "crypto/sha256"

func popCount(b byte) int {
	count := 0
	for ; b != 0; count++ {
		b &= b - 1
	}
	return count
}

func bitDiff(firstHash, secondHash []byte) int {
	count := 0

	for i := 0; i < len(firstHash) || i < len(secondHash); i++ {
		switch {
		case i > len(firstHash):
			count += popCount(secondHash[i])
		case i > len(secondHash):
			count += popCount(firstHash[i])
		default:
			count += popCount(firstHash[i] ^ secondHash[i])
		}
	}

	return count
}

func DifferentBitsSHA256(a, b []byte) int {
	shaA := sha256.Sum256(a)
	shaB := sha256.Sum256(b)

	return bitDiff(shaA[:], shaB[:])
}
