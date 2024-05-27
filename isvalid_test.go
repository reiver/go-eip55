package eip55_test

import (
	"testing"

	"github.com/reiver/go-eip55"
)

func TestIsValid(t *testing.T) {

	tests := []struct{
		Encoded string
		Expected bool
	}{
		{
			Encoded: "0x52908400098527886E0F7030069857D2E4169EE7",
			Expected: true,
		},
		{
			Encoded: "0x8617E340B3D01FA5F11F306F4090FD50E238070D",
			Expected: true,
		},



		{
			Encoded: "0xde709f2102306220921060314715629080e2fb77",
			Expected: true,
		},
		{
			Encoded: "0x27b1fdb04752bbc536007a920d24acb045561c26",
			Expected: true,
		},



		{
			Encoded: "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed",
			Expected: true,
		},
		{
			Encoded: "0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359",
			Expected: true,
		},
		{
			Encoded: "0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB",
			Expected: true,
		},
		{
			Encoded: "0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb",
			Expected: true,
		},









		{
			Encoded: "0xde709f2102306220921060314715629080e2fB77",
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := eip55.IsValid(test.Encoded)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual EIP-55 encoding is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}
