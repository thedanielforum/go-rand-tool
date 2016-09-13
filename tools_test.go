package randtool

import (
	"testing"
)

func TestGenerateSeed(t *testing.T) {
	// Execute to check for panic
	GenerateSeed()
}

func TestGenerateString(t *testing.T) {
	// Check for correct generation
	var err error
	r := make(map[string]bool)
	a, err := GenerateString(8)
	b, err := GenerateString(8)
	c, err := GenerateString(8)
	d, err := GenerateString(8)
	if err != nil {
		t.Fatal(err.Error())
	}
	r[a] = true
	r[b] = true
	r[c] = true
	r[d] = true
	// If the output is unique we expect 4 slices
	if len(r) != 4 {
		t.Fatalf("Key coalition expected 4 unique keys got %d", len(r))
	}
}
