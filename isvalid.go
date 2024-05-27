package eip55

import (
	"encoding/hex"
)

func IsValid(value string) bool {
	{
		const expectedlength int = ArrayLength*2 + lenprefix
		if expectedlength != len(value) {
			return false
		}
	}

	{
		const expected string = "0x"
		actual := value[:lenprefix]

		if expected != actual {
			return false
		}
	}

	var decoded [ArrayLength]byte
	{
		temp, err := hex.DecodeString(value[lenprefix:])
		if nil != err {
			return false
		}

		copy(decoded[:], temp)
	}

	var reencoded string = Encode(decoded)

	return reencoded == value
}
