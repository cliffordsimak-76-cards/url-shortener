package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

// StringToMD5 returns first 8 bytes MD5 checksum of the data.
func StringToMD5(str string) string {
	h := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", h[:8])
}

// GenerateRandom returns random bytes.
func GenerateRandom(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateRandom returns a new HMAC.
func SignHMAC256(msg []byte, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(msg)
	return h.Sum(nil)
}
