package randtool

import (
	"testing"
)

/*
Unit tests
 */

func TestGenInt64(t *testing.T) {
	// Execute to check for panic
	GenInt64()
}

func TestSeedMathRand(t *testing.T) {
	// Execute to check for panic
	SeedMathRand()
}

func TestGenStr(t *testing.T) {
	// Check for correct generation and "basic entropy"
	var err error
	r := make(map[string]bool)
	a, err := GenStr(6)
	b, err := GenStr(6)
	c, err := GenStr(6)
	d, err := GenStr(6)
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
		t.Errorf("Key coalition expected 4 unique keys got %d", len(r))
	}
}

func TestGenStr2(t *testing.T) {
	// Check expected string length
	rs, _ := GenStr(32)
	if len(rs) != 32 {
		t.Errorf("Expected string with length of 32 got %d", len(rs))
	}
}

func TestGenStr3(t *testing.T) {
	// < 1 should generate a invalid length error
	_, err := GenStr(0)
	if err == nil {
		t.Error("Expected err 'Random string length must be greater than 0' got 'nil'")
	}
}

func TestGenStrIgnoreErr(t *testing.T) {
	// Expecting string of 32 char
	sl := len(GenStrIgnoreErr(32))
	if sl != 32 {
		t.Errorf("Expected string with length of 32 got %d", sl)
	}
}

func TestGenStrIgnoreErr2(t *testing.T) {
	// Check for empty string on invalid input
	s := GenStrIgnoreErr(0)
	if s != "" {
		t.Errorf("Expected '' got %s", s)
	}
}

/*
Performance tests
 */

// Store the bench result to package level so the compiler
// cannot eliminate the Benchmark itself.
var rGenInt64 int64
var rGenStr   string

func BenchmarkGenInt64(b *testing.B) {
	var r int64
	for n := 0; n < b.N; n++ {
		// Record the result of GenInt64 to prevent
		// the compiler eliminating the function call.
		r = GenInt64()
	}
	rGenInt64 = r
}

func BenchmarkGenStr(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r, _ = GenStr(16)
	}
	rGenStr = r
}
