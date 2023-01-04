package steggo

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
)

func encrypt(plainText, key []byte) []byte {
	keyMd5 := md5.Sum(key)
	block, _ := aes.NewCipher(keyMd5[:])
	cfb := cipher.NewCFBEncrypter(block, keyMd5[:])
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return cipherText
}

func decrypt(cipherText, key []byte) []byte {
	keyMd5 := md5.Sum(key)
	block, _ := aes.NewCipher(keyMd5[:])
	cfb := cipher.NewCFBDecrypter(block, keyMd5[:])
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return plainText
}
