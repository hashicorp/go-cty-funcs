package crypto

import (
	"fmt"
	"testing"

	"github.com/zclconf/go-cty/cty"
)

func TestMd5(t *testing.T) {
	tests := []struct {
		String cty.Value
		Want   cty.Value
		Err    bool
	}{
		{
			cty.StringVal("tada"),
			cty.StringVal("ce47d07243bb6eaf5e1322c81baf9bbf"),
			false,
		},
		{ // Confirm that we're not trimming any whitespaces
			cty.StringVal(" tada "),
			cty.StringVal("aadf191a583e53062de2d02c008141c4"),
			false,
		},
		{ // We accept empty string too
			cty.StringVal(""),
			cty.StringVal("d41d8cd98f00b204e9800998ecf8427e"),
			false,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("md5(%#v)", test.String), func(t *testing.T) {
			got, err := Md5(test.String)

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
