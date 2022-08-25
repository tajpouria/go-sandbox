package utils

import (
	"math/rand"
	"testing"
)

func TestRandom(t *testing.T) {
	a := RandomURL(uint32(rand.Intn(100)))
	b := RandomURL(uint32(rand.Intn(100)))
	if a == b {
		t.Errorf("Generated URLs are identical")
	} else {
		t.Log("Generated URLs are random")
	}
}
