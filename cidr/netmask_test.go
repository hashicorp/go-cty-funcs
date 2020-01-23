package cidr

import (
	"fmt"
	"testing"

	"github.com/zclconf/go-cty/cty"
)

func TestNetmask(t *testing.T) {
	tests := []struct {
		Prefix cty.Value
		Want   cty.Value
		Err    bool
	}{
		{
			cty.StringVal("192.168.1.0/24"),
			cty.StringVal("255.255.255.0"),
			false,
		},
		{
			cty.StringVal("192.168.1.0/32"),
			cty.StringVal("255.255.255.255"),
			false,
		},
		{
			cty.StringVal("0.0.0.0/0"),
			cty.StringVal("0.0.0.0"),
			false,
		},
		{
			cty.StringVal("1::/64"),
			cty.StringVal("ffff:ffff:ffff:ffff::"),
			false,
		},
		{
			cty.StringVal("not-a-cidr"),
			cty.UnknownVal(cty.String),
			true, // not a valid CIDR mask
		},
		{
			cty.StringVal("110.256.0.0/8"),
			cty.UnknownVal(cty.String),
			true, // can't have an octet >255
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("cidrnetmask(%#v)", test.Prefix), func(t *testing.T) {
			got, err := Netmask(test.Prefix)

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
