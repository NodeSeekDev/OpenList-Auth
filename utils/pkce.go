package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

func GenerateCodeVerifier(n int) (string, error) {
	// generate CodeVerifier for 115Open
	if n < 43 || n > 128 {
		return "", errors.New("code_verifier length must be between 43 and 128")
	}
	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	s := base64.RawURLEncoding.EncodeToString(buf)
	return s[:n], nil
}
