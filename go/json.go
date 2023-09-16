package iso

import "bytes"

var jsonNull = []byte("null")

func isJsonNull(bs []byte) bool {
	return bytes.Equal(bs, jsonNull)
}

func unquote(bs []byte) []byte {
	if len(bs) < 2 {
		return bs
	}
	if bs[0] == '"' && bs[len(bs)-1] == '"' {
		return bs[1 : len(bs)-1]
	}
	return bs
}
