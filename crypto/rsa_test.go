package crypto

import (
	"fmt"
	"testing"

	"github.com/zclconf/go-cty/cty"
)

func TestRsaDecrypt(t *testing.T) {
	tests := []struct {
		Ciphertext cty.Value
		Privatekey cty.Value
		Want       cty.Value
		Err        bool
	}{
		// Base-64 encoded cipher decrypts correctly
		{
			cty.StringVal(CipherBase64),
			cty.StringVal(PrivateKey),
			cty.StringVal("message"),
			false,
		},
		// Wrong key
		{
			cty.StringVal(CipherBase64),
			cty.StringVal(WrongPrivateKey),
			cty.UnknownVal(cty.String),
			true,
		},
		// Bad key
		{
			cty.StringVal(CipherBase64),
			cty.StringVal("bad key"),
			cty.UnknownVal(cty.String),
			true,
		},
		// Empty key
		{
			cty.StringVal(CipherBase64),
			cty.StringVal(""),
			cty.UnknownVal(cty.String),
			true,
		},
		// Bad cipher
		{
			cty.StringVal("bad cipher"),
			cty.StringVal(PrivateKey),
			cty.UnknownVal(cty.String),
			true,
		},
		// Empty cipher
		{
			cty.StringVal(""),
			cty.StringVal(PrivateKey),
			cty.UnknownVal(cty.String),
			true,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("RsaDecrypt(%#v, %#v)", test.Ciphertext, test.Privatekey), func(t *testing.T) {
			got, err := RsaDecrypt(test.Ciphertext, test.Privatekey)

			if test.Err {
				if err == nil {
					t.Fatal("succeeded; want error")
				}
				return
			} else if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if !got.RawEquals(test.Want) {
				t.Errorf("wrong result\ngot:  %#v\nwant: %#v", got, test.Want)
			}
		})
	}
}
