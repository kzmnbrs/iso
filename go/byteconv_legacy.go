//go:build !go1.20 && !go1.21
// +build !go1.20,!go1.21

package iso

import (
	"reflect"
	"unsafe"
)

func b2s(bb []byte) string {
	bh := *(*reflect.SliceHeader)(unsafe.Pointer(&bb))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: bh.Data,
		Len:  bh.Len,
	}))
}

func s2b(s string) []byte {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}))
}
