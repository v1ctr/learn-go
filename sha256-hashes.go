package main

import (
	"crypto/sha256"
	"fmt"
)

// You can compute other hashes using a similar pattern to the one shown above.
// For example, to compute SHA512 hashes import crypto/sha512 and use sha512.New().

func main() {
	s := "sha256 this string"

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
