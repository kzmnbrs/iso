//go:build go1.20 || go1.21
// +build go1.20 go1.21

package iso

import "unsafe"

func b2s(bb []byte) string {
	return unsafe.String(unsafe.SliceData(bb), len(bb))
}

func s2b(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
