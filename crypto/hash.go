package crypto

import (
	"fmt"
	"hash"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

func makeStringHashFunction(hf func() hash.Hash, enc func([]byte) string) function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "str",
				Type: cty.String,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (ret cty.Value, err error) {
			s := args[0].AsString()
			h := hf()
			h.Write([]byte(s))
			rv := enc(h.Sum(nil))
			return cty.StringVal(rv), nil
		},
	})
}

func makeFileHashFunction(baseDir string, hf func() hash.Hash, enc func([]byte) string) function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "path",
				Type: cty.String,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (ret cty.Value, err error) {
			path := args[0].AsString()
			src, err := readFileBytes(baseDir, path)
			if err != nil {
				return cty.UnknownVal(cty.String), err
			}

			h := hf()
			h.Write(src)
			rv := enc(h.Sum(nil))
			return cty.StringVal(rv), nil
		},
	})
}

func readFileBytes(baseDir, path string) ([]byte, error) {
	path, err := homedir.Expand(path)
	if err != nil {
		return nil, fmt.Errorf("failed to expand ~: %s", err)
	}

	if !filepath.IsAbs(path) {
		path = filepath.Join(baseDir, path)
	}

	// Ensure that the path is canonical for the host OS
	path = filepath.Clean(path)

	src, err := ioutil.ReadFile(path)
	if err != nil {
		// ReadFile does not return Terraform-user-friendly error messages, so
		// we'll provide our own.
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("no file exists at %s", path)
		}
		return nil, fmt.Errorf("failed to read %s", path)
	}

	return src, nil
}
