package steggo

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHuffman(t *testing.T) {
	var buff bytes.Buffer
	buff.Write([]byte("hello world"))
	b := buff.Bytes()

	huffman := Huffman(b)
	unHuffman, err := UnHuffman(huffman)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(unHuffman); i++ {
		fmt.Printf("%08b %08b\n", b[i], unHuffman[i])
	}

	assert.Equal(t, b, unHuffman)
}
