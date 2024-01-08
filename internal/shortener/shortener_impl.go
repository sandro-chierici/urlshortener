package shortener

import (
	"crypto/sha256"
	"fmt"
	"hash"
)

type SimpleShortener struct {
	h hash.Hash
}

func New() *SimpleShortener {
	var sh = &SimpleShortener{}
	sh.h = sha256.New()
	return sh
}

// shortener algo
func (s *SimpleShortener) Evaluate(url string) string {

	s.h.Reset()
	s.h.Write([]byte(url))

	return fmt.Sprintf("%x", s.h.Sum(nil))
}
