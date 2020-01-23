package cidr

import (
	"fmt"
	"testing"

	"github.com/zclconf/go-cty/cty"
)

func TestSubnets(t *testing.T) {
	tests := []struct {
		Prefix  cty.Value
		Newbits []cty.Value
		Want    cty.Value
		Err     string
	}{
		{
			cty.StringVal("10.0.0.0/21"),
			[]cty.Value{
				cty.NumberIntVal(3),
				cty.NumberIntVal(3),
				cty.NumberIntVal(3),
				cty.NumberIntVal(4),
				cty.NumberIntVal(4),
				cty.NumberIntVal(4),
				cty.NumberIntVal(7),
				cty.NumberIntVal(7),
				cty.NumberIntVal(7),
			},
			cty.ListVal([]cty.Value{
				cty.StringVal("10.0.0.0/24"),
				cty.StringVal("10.0.1.0/24"),
				cty.StringVal("10.0.2.0/24"),
				cty.StringVal("10.0.3.0/25"),
				cty.StringVal("10.0.3.128/25"),
				cty.StringVal("10.0.4.0/25"),
				cty.StringVal("10.0.4.128/28"),
				cty.StringVal("10.0.4.144/28"),
				cty.StringVal("10.0.4.160/28"),
			}),
			``,
		},
		{
			cty.StringVal("10.0.0.0/30"),
			[]cty.Value{
				cty.NumberIntVal(1),
				cty.NumberIntVal(3),
			},
			cty.UnknownVal(cty.List(cty.String)),
			`would extend prefix to 33 bits, which is too long for an IPv4 address`,
		},
		{
			cty.StringVal("10.0.0.0/8"),
			[]cty.Value{
				cty.NumberIntVal(1),
				cty.NumberIntVal(1),
				cty.NumberIntVal(1),
			},
			cty.UnknownVal(cty.List(cty.String)),
			`not enough remaining address space for a subnet with a prefix of 9 bits after 10.128.0.0/9`,
		},
		{
			cty.StringVal("10.0.0.0/8"),
			[]cty.Value{
				cty.NumberIntVal(1),
				cty.NumberIntVal(0),
			},
			cty.UnknownVal(cty.List(cty.String)),
			`must extend prefix by at least one bit`,
		},
		{
			cty.StringVal("10.0.0.0/8"),
			[]cty.Value{
				cty.NumberIntVal(1),
				cty.NumberIntVal(-1),
			},
			cty.UnknownVal(cty.List(cty.String)),
			`must extend prefix by at least one bit`,
		},
		{
			cty.StringVal("fe80::/48"),
			[]cty.Value{
				cty.NumberIntVal(1),
				cty.NumberIntVal(33),
			},
			cty.UnknownVal(cty.List(cty.String)),
			`may not extend prefix by more than 32 bits`,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("cidrsubnets(%#v, %#v)", test.Prefix, test.Newbits), func(t *testing.T) {
			got, err := Subnets(test.Prefix, test.Newbits...)
			wantErr := test.Err != ""

			if wantErr {
				if err == nil {
					t.Fatal("succeeded; want error")
				}
				if err.Error() != test.Err {
					t.Fatalf("wrong error\ngot:  %s\nwant: %s", err.Error(), test.Err)
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
