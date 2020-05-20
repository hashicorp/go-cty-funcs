package uuid

import (
	"fmt"
	"testing"

	"github.com/zclconf/go-cty/cty"
)

func TestV5(t *testing.T) {
	tests := []struct {
		Namespace cty.Value
		Name      cty.Value
		Want      cty.Value
		Err       bool
	}{
		{
			cty.StringVal("dns"),
			cty.StringVal("tada"),
			cty.StringVal("faa898db-9b9d-5b75-86a9-149e7bb8e3b8"),
			false,
		},
		{
			cty.StringVal("url"),
			cty.StringVal("tada"),
			cty.StringVal("2c1ff6b4-211f-577e-94de-d978b0caa16e"),
			false,
		},
		{
			cty.StringVal("oid"),
			cty.StringVal("tada"),
			cty.StringVal("61eeea26-5176-5288-87fc-232d6ed30d2f"),
			false,
		},
		{
			cty.StringVal("x500"),
			cty.StringVal("tada"),
			cty.StringVal("7e12415e-f7c9-57c3-9e43-52dc9950d264"),
			false,
		},
		{
			cty.StringVal("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
			cty.StringVal("tada"),
			cty.StringVal("faa898db-9b9d-5b75-86a9-149e7bb8e3b8"),
			false,
		},
		{
			cty.StringVal("tada"),
			cty.StringVal("tada"),
			cty.UnknownVal(cty.String),
			true,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("uuidv5(%#v, %#v)", test.Namespace, test.Name), func(t *testing.T) {
			got, err := V5(test.Namespace, test.Name)

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
