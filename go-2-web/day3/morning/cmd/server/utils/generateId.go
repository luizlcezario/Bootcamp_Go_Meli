package utils

import (
	"crypto/sha256"
	"fmt"
)

func GeneratedId(s string) string {
	h := sha256.New()
	if _, err := h.Write([]byte(s)); err != nil {
		panic(err.Error())
	} else {
		bt := h.Sum(nil)
		return fmt.Sprintf("%x", bt)
	}
}
