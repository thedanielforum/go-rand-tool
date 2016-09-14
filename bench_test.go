package randtool

import "testing"

// Store the bench result to package level so the compiler
// cannot eliminate the Benchmark itself.
var rGenInt64 int64
var rGenStr   string

func BenchmarkGenInt64(b *testing.B) {
	var r int64
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		// Record the result of GenInt64 to prevent
		// the compiler eliminating the function call.
		r = GenInt64()
	}
	rGenInt64 = r
}

func BenchmarkGenStr(b *testing.B) {
	var r string
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		r, _ = GenStr(16)
	}
	rGenStr = r
}
