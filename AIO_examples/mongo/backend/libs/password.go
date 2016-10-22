package libs

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	Pass string
}

var SecretKey string = "ZA1XSWSekret128cdevfraASDFlkjhHg"

func (w Password) Gen(p string) string {

	password := []byte(p + SecretKey)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	w.Pass = string(hashedPassword)

	return string(hashedPassword)
}

func (w Password) Compare(hs string, p string) bool {
	hash := []byte(hs)
	pass := []byte(p + SecretKey)
	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword(hash, pass)
	// fmt.Println(err) // nil means it is a match

	r := false
	if err == nil {
		r = true
	}

	return r
}

func (w Password) Token() string {
	t, err := w.Random(32)
	if err != nil {
		panic(err)
	}

	return t
}

func (w Password) RandomByte(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (w Password) Random(s int) (string, error) {
	b, err := w.RandomByte(s)
	return base64.URLEncoding.EncodeToString(b), err
}
