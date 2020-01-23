package cidr

import (
	"fmt"
	"testing"

	"github.com/zclconf/go-cty/cty"
)

func TestHost(t *testing.T) {
	tests := []struct {
		Prefix  cty.Value
		Hostnum cty.Value
		Want    cty.Value
		Err     bool
	}{
		{
			cty.StringVal("192.168.1.0/24"),
			cty.NumberIntVal(5),
			cty.StringVal("192.168.1.5"),
			false,
		},
		{
			cty.StringVal("192.168.1.0/24"),
			cty.NumberIntVal(-5),
			cty.StringVal("192.168.1.251"),
			false,
		},
		{
			cty.StringVal("192.168.1.0/24"),
			cty.NumberIntVal(-256),
			cty.StringVal("192.168.1.0"),
			false,
		},
		{
			cty.StringVal("192.168.1.0/30"),
			cty.NumberIntVal(255),
			cty.UnknownVal(cty.String),
			true, // 255 doesn't fit in two bits
		},
		{
			cty.StringVal("192.168.1.0/30"),
			cty.NumberIntVal(-255),
			cty.UnknownVal(cty.String),
			true, // 255 doesn't fit in two bits
		},
		{
			cty.StringVal("not-a-cidr"),
			cty.NumberIntVal(6),
			cty.UnknownVal(cty.String),
			true, // not a valid CIDR mask
		},
		{
			cty.StringVal("10.256.0.0/8"),
			cty.NumberIntVal(6),
			cty.UnknownVal(cty.String),
			true, // can't have an octet >255
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("cidrhost(%#v, %#v)", test.Prefix, test.Hostnum), func(t *testing.T) {
			got, err := Host(test.Prefix, test.Hostnum)

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
