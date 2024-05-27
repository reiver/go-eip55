package eip55

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

func Encode(src [ArrayLength]byte) string {

	var encoded [ArrayLength*2 + lenprefix]byte
	{
		copy(encoded[:lenprefix], prefix)
		hex.Encode(encoded[lenprefix:], src[:])
	}

	var digest []byte
	{
		hashfunc := sha3.NewLegacyKeccak256()
		hashfunc.Write(encoded[lenprefix:])
		digest = hashfunc.Sum(nil)
	}

	{
		var limit int = len(encoded)

		for i:=lenprefix; i<limit; i++ {
			var b byte = encoded[i]

			if b < 'a' || 'z' < b {
				continue
			}

			var bb byte = digest[(i-lenprefix)/2]
			if 0 == (i%2) {
				bb = bb >> 4
			}

			const mask byte = 0b0000_1000
			bb &= mask

			if mask == bb && 'a' <= b && b <= 'z' {
				encoded[i] -= 'a'
				encoded[i] += 'A'
			}
		}
	}

	return string(encoded[:])
}
