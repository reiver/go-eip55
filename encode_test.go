package eip55_test

import (
	"testing"

	"encoding/hex"
	"strings"

	"github.com/reiver/go-eip55"
)

func TestEncode(t *testing.T) {

	tests := []struct{
		Address  string
		Expected string
	}{
		{
			Address: strings.ToLower("52908400098527886E0F7030069857D2E4169EE7"),
			Expected:              "0x52908400098527886E0F7030069857D2E4169EE7",
		},
		{
			Address: strings.ToLower("8617E340B3D01FA5F11F306F4090FD50E238070D"),
			Expected:              "0x8617E340B3D01FA5F11F306F4090FD50E238070D",
		},



		{
			Address: strings.ToLower("de709f2102306220921060314715629080e2fb77"),
			Expected:              "0xde709f2102306220921060314715629080e2fb77",
		},
		{
			Address: strings.ToLower("27b1fdb04752bbc536007a920d24acb045561c26"),
			Expected:              "0x27b1fdb04752bbc536007a920d24acb045561c26",
		},



		{
			Address: strings.ToLower("5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed"),
			Expected:              "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed",
		},
		{
			Address: strings.ToLower("fB6916095ca1df60bB79Ce92cE3Ea74c37c5d359"),
			Expected:              "0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359",
		},
		{
			Address: strings.ToLower("dbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB"),
			Expected:              "0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB",
		},
		{
			Address: strings.ToLower("D1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb"),
			Expected:              "0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb",
		},
	}

	for testNumber, test := range tests {

		var address [eip55.Length]byte
		{
			var p []byte = address[0:0]
			_, err := hex.AppendDecode(p, []byte(test.Address))
			if nil != err {
				panic(err)
			}
		}
		{
			expected := test.Address
			actual := hex.EncodeToString(address[:])

			if expected != actual {
				t.Errorf("For test #%d, the encoding-decoding of the address was wrong — this is a problem with the test not the code.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}

		actual := eip55.Encode(address)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual EIP-55 encoding is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
