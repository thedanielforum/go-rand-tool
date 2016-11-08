![godoc](https://travis-ci.org/thedanielforum/randtool.svg?branch=master)

GO package with some easy tools for generating common types of pseudo
random data.

**Usage Examples**
```cmd
go test
```
```go
// Generate a random 16 char string
str, err := randtool.GenStr(16)
```

**Benchmarks**
```cmd
go test -bench=.
```
```go
BenchmarkGenInt64	 2000000	       715 ns/op	      16 B/op	       2 allocs/op
BenchmarkGenStr  	10000000	       235 ns/op	      32 B/op	       2 allocs/op
```
