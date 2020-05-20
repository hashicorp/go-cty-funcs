package cidr

import (
	"fmt"
	"testing"

	"github.com/zclconf/go-cty/cty"
)

func TestSubnet(t *testing.T) {
	tests := []struct {
		Prefix  cty.Value
		Newbits cty.Value
		Netnum  cty.Value
		Want    cty.Value
		Err     bool
	}{
		{
			cty.StringVal("192.168.2.0/20"),
			cty.NumberIntVal(4),
			cty.NumberIntVal(6),
			cty.StringVal("192.168.6.0/24"),
			false,
		},
		{
			cty.StringVal("fe80::/48"),
			cty.NumberIntVal(16),
			cty.NumberIntVal(6),
			cty.StringVal("fe80:0:0:6::/64"),
			false,
		},
		{ // IPv4 address encoded in IPv6 syntax gets normalized
			cty.StringVal("::ffff:192.168.0.0/112"),
			cty.NumberIntVal(8),
			cty.NumberIntVal(6),
			cty.StringVal("192.168.6.0/24"),
			false,
		},
		{ // not enough bits left
			cty.StringVal("192.168.0.0/30"),
			cty.NumberIntVal(4),
			cty.NumberIntVal(6),
			cty.UnknownVal(cty.String),
			true,
		},
		{ // can't encode 16 in 2 bits
			cty.StringVal("192.168.0.0/168"),
			cty.NumberIntVal(2),
			cty.NumberIntVal(16),
			cty.StringVal("fe80:0:0:6::/64"),
			true,
		},
		{ // not a valid CIDR mask
			cty.StringVal("not-a-cidr"),
			cty.NumberIntVal(4),
			cty.NumberIntVal(6),
			cty.StringVal("fe80:0:0:6::/64"),
			true,
		},
		{ // can't have an octet >255
			cty.StringVal("10.256.0.0/8"),
			cty.NumberIntVal(4),
			cty.NumberIntVal(6),
			cty.StringVal("fe80:0:0:6::/64"),
			true,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("cidrsubnet(%#v, %#v, %#v)", test.Prefix, test.Newbits, test.Netnum), func(t *testing.T) {
			got, err := Subnet(test.Prefix, test.Newbits, test.Netnum)

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
