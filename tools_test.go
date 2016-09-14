package randtool

import (
	"testing"
)

func TestGenerateSeed(t *testing.T) {
	// Execute to check for panic
	GenerateSeed()
}

func TestGenerateString(t *testing.T) {
	// Check for correct generation and "basic entropy"
	var err error
	r := make(map[string]bool)
	a, err := GenerateString(6)
	b, err := GenerateString(6)
	c, err := GenerateString(6)
	d, err := GenerateString(6)
	// If we got a error on execution
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

func TestGenerateString2(t *testing.T) {
	// Check expected string length
	rs, _ := GenerateString(32)
	if len(rs) != 32 {
		t.Fatalf("Expected string with length of 32 got %d", len(rs))
	}
}

func TestGenerateString3(t *testing.T) {
	// < 1 should generate a invalid length error
	_, err := GenerateString(0)
	if err == nil {
		t.Fatal("Expected err 'Random string length must be greater than 0' got 'nil'")
	}
}

func TestGenerateStringIgnErr(t *testing.T) {
	// Expecting string of 32 char
	sl := len(GenerateStringIgnErr(32))
	if sl != 32 {
		t.Fatalf("Expected string with length of 32 got %d", sl)
	}
}

func TestGenerateStringIgnErr2(t *testing.T) {
	// Check for empty string on invalid input
	s := GenerateStringIgnErr(0)
	if s != "" {
		t.Fatalf("Expected '' got %s", s)
	}
}
