package steggo

import (
	"bytes"
	"fmt"
	"testing"
)

func TestHuffman(t *testing.T) {
	var buff bytes.Buffer
	buff.Write([]byte("hello world"))
	b := buff.Bytes()

	huffman := Huffman(b)
	unHuffman, err := UnHuffman(huffman)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(unHuffman); i++ {
		fmt.Printf("%08b %08b\n", b[i], unHuffman[i])
	}

	if !bytes.Equal(b, unHuffman) {
		t.Error("The data before Huffman and after UnHuffman are not equal")
	}
}
