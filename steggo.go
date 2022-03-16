package steggo

import (
	"github.com/bkaradzic/go-lz4"
)

func Encode(s string, key string) (string, error) {
	dst, _ := lz4.Encode(nil, []byte(s))
	h := Huffman(dst)
	return h, nil
}

func Decode(s string, key string) (string, error) {
	b, err := UnHuffman(s)
	if err != nil {
		return "", err
	}
	decode, err := lz4.Decode(nil, b)
	if err != nil {
		return "", err
	}

	return string(decode), nil
}
