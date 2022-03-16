package steggo

import (
	"testing"
)

func TestBasic(t *testing.T) {
	plain := "99999999999999"
	encode, err := Encode(plain, "")
	if err != nil {
		return
	}
	decode, err := Decode(encode, "")
	if err != nil {
		return
	}
	if plain != decode {
		t.Error("decode failed.")
	}
}
