package steggo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAES(t *testing.T) {
	var tests = []struct {
		plainText, decryptedText, key []byte
		equal                         bool
	}{
		{[]byte("steggo"), []byte("steggo"), []byte("steggo"), true},
		{[]byte("steggosteggo"), []byte("steggosteggo"), []byte("steggo"), true},
		{[]byte("steggo"), []byte("steg"), []byte("steggo"), false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s %s %s", test.plainText, test.decryptedText, test.key), func(t *testing.T) {
			cipherText := encrypt(test.plainText, test.key)
			decryptedText := decrypt(cipherText, test.key)
			if test.equal {
				assert.Equal(t, test.decryptedText, decryptedText)
			} else {
				assert.NotEqual(t, test.decryptedText, decryptedText)
			}
		})
	}
}
