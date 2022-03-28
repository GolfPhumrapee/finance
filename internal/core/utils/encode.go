package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"net/url"
)

// EncodeStringBase64 encode string base64
func EncodeStringBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// EncodeParam endcode param
func EncodeParam(s string) string {
	return url.QueryEscape(s)
}

func md5HashByte(text string) []byte {
	md5 := md5.New()
	_, _ = md5.Write([]byte(text))
	hashByte := md5.Sum(nil)
	return hashByte
}

// MD5HashHex md5 hash byte to string hex
func MD5HashHex(text string) string {
	hashByte := md5HashByte(text)
	return hex.EncodeToString(hashByte)
}

// MD5Encode md5 hash byte to string base64 with prefix {MD5}
func MD5HashEncode(text string) string {
	hash := md5HashByte(text)
	return base64.StdEncoding.EncodeToString(hash)
}
