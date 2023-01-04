package steggo

import (
	"github.com/bkaradzic/go-lz4"
)

func Encode(s []byte, key ...[]byte) (string, error) {
	dst, _ := lz4.Encode(nil, s)
	if len(key) != 0 {
		dst = encrypt(dst, key[0])
	}
	h := Huffman(dst)
	return h, nil
}

func Decode(s string, key ...[]byte) ([]byte, error) {
	b, err := UnHuffman(s)
	if err != nil {
		return []byte{}, err
	}

	if len(key) != 0 {
		b = decrypt(b, key[0])
	}

	decode, err := lz4.Decode(nil, b)
	if err != nil {
		return []byte{}, err
	}

	return decode, nil
}
