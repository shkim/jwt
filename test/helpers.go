package test

import (
	"crypto"
	"crypto/rsa"
	"io/ioutil"

	"github.com/golang-jwt/jwt/v4"
)

func LoadRSAPrivateKeyFromDisk(location string) *rsa.PrivateKey {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

func LoadRSAPublicKeyFromDisk(location string) *rsa.PublicKey {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

// MakeSampleToken creates and returns a encoded JWT token that has been signed with the specified cryptographic key.
func MakeSampleToken[T jwt.Key](c jwt.Claims, method jwt.SigningMethod[T], key interface{}) string {
	token := jwt.NewWithClaims(method, c)
	s, e := token.SignedString(key)

	if e != nil {
		panic(e.Error())
	}

	return s
}

func LoadECPrivateKeyFromDisk(location string) crypto.PrivateKey {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseECPrivateKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

func LoadECPublicKeyFromDisk(location string) crypto.PublicKey {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseECPublicKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}
