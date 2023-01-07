package steggo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasic(t *testing.T) {
	plain := []byte("steggo")
	encode, err := Encode(plain)
	if err != nil {
		t.Error(err)
		return
	}
	decode, err := Decode(encode)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, plain, decode)

	password := []byte("atom")
	encode, err = Encode(plain, password)
	if err != nil {
		t.Error(err)
		return
	}
	decode, err = Decode(encode, password)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, plain, decode)
}
