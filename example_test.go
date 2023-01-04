package steggo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasic(t *testing.T) {
	plain := []byte("steggoo")
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

	encode, err = Encode(plain, plain)
	if err != nil {
		t.Error(err)
		return
	}
	decode, err = Decode(encode, plain)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, plain, decode)
}
