package uuid

import (
	"testing"
)

func TestV4(t *testing.T) {
	result, err := V4()
	if err != nil {
		t.Fatal(err)
	}

	resultStr := result.AsString()
	if got, want := len(resultStr), 36; got != want {
		t.Errorf("wrong result length %d; want %d", got, want)
	}
}
