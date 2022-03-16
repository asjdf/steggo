package steggo

import (
	"errors"
	"strings"
)

// HuffmanConvertMap
const (
	BIT00   = "\u200c"
	BIT01   = "\u200d"
	BIT10   = "\u2060"
	BIT11   = "\u2062"
	BIT0101 = "\u2063"
	BIT0000 = "\u2064"
)

var HuffmanMap = map[uint8]string{
	0x00: BIT00,
	0x01: BIT01,
	0x02: BIT10,
	0x03: BIT11,
}

var UnHuffmanMap = map[string]uint8{
	BIT00: 0x00,
	BIT01: 0x01,
	BIT10: 0x02,
	BIT11: 0x03,
}

// Huffman to convert byte to string
func Huffman(b []byte) string {
	out := ""
	for i := 0; i < len(b); i++ {
		for j := uint(0); j < 8; j = j + 2 {
			s := HuffmanMap[b[i]&(0x03<<j)>>j]
			out += s
		}
	}
	out = strings.ReplaceAll(out, BIT01+BIT01, BIT0101)
	out = strings.ReplaceAll(out, BIT00+BIT00, BIT0000)
	return out
}

func UnHuffman(in string) (b []byte, err error) {
	in = strings.ReplaceAll(in, BIT0101, BIT01+BIT01)
	in = strings.ReplaceAll(in, BIT0000, BIT00+BIT00)
	b = make([]byte, 0)
	pos := uint(0)
	temp := byte(0)
	for _, i := range strings.Split(in, "") {
		if v, ok := UnHuffmanMap[i]; ok {
			temp |= v << pos
			pos += 2
			if pos == 8 {
				b = append(b, temp)
				pos = 0
				temp = 0
			}
		} else {
			return nil, errors.New("invalid huffman code")
		}
	}
	return b, nil
}
