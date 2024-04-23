package utils

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func GeneratedId(s string) string {
	h := sha256.New()
	env := os.Getenv("IDKEYGEN")
	if _, err := h.Write([]byte(s + env)); err != nil {
		panic(err.Error())
	} else {
		bt := h.Sum(nil)
		return fmt.Sprintf("%x", bt)
	}
}
