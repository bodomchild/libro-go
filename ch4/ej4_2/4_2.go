package ej4_2

import (
	"crypto/sha512"
	"encoding/base64"
)

func SHA512Enc(input []byte) string {
	hasher := sha512.New()
	hasher.Write(input)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

func SHA384Enc(input []byte) string {
	hasher := sha512.New384()
	hasher.Write(input)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

func SHA256Enc(input []byte) string {
	hasher := sha512.New512_256()
	hasher.Write(input)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
