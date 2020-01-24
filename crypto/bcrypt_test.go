package crypto

import (
	"testing"

	"github.com/zclconf/go-cty/cty"
	"golang.org/x/crypto/bcrypt"
)

func TestBcrypt(t *testing.T) {
	// single variable test
	p, err := Bcrypt(cty.StringVal("test"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(p.AsString()), []byte("test"))
	if err != nil {
		t.Fatalf("Error comparing hash and password: %s", err)
	}

	// testing with two parameters
	p, err = Bcrypt(cty.StringVal("test"), cty.NumberIntVal(5))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(p.AsString()), []byte("test"))
	if err != nil {
		t.Fatalf("Error comparing hash and password: %s", err)
	}

	// Negative test for more than two parameters
	_, err = Bcrypt(cty.StringVal("test"), cty.NumberIntVal(10), cty.NumberIntVal(11))
	if err == nil {
		t.Fatal("succeeded; want error")
	}
}
