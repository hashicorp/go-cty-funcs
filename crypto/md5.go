package crypto

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

// MakeFileMd5Func is a function that is like Md5Func but reads the contents of
// a file rather than hashing a given literal string.
func MakeFileMd5Func(baseDir string) function.Function {
	return makeFileHashFunction(baseDir, md5.New, hex.EncodeToString)
}

// Md5Func is a function that computes the MD5 hash of a given string and
// encodes it with hexadecimal digits.
var Md5Func = makeStringHashFunction(md5.New, hex.EncodeToString)

// Md5 computes the MD5 hash of a given string and encodes it with hexadecimal
// digits.
func Md5(str cty.Value) (cty.Value, error) {
	return Md5Func.Call([]cty.Value{str})
}
