// Code generated by 'github.com/traefik/yaegi/extract crypto/aes'. DO NOT EDIT.

// +build go1.15,!go1.16

package stdlib

import (
	"crypto/aes"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["crypto/aes"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BlockSize": reflect.ValueOf(constant.MakeFromLiteral("16", token.INT, 0)),
		"NewCipher": reflect.ValueOf(aes.NewCipher),

		// type definitions
		"KeySizeError": reflect.ValueOf((*aes.KeySizeError)(nil)),
	}
}
