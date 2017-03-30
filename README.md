![godoc](https://travis-ci.org/thedanielforum/randtool.svg?branch=master)

GO package with some easy tools for generating common types of pseudo
random data.

**Installation**
```go
go get github.com/thedanielforum/randtool
```

**Usage Examples**
Generate a random string of n length.
```go
// Generate a random 16 char string
str, err := randtool.GenStr(16)
// Generate a new string but ignore any potential errors
str := randtool.GenStrIgnoreErr(16)
```

Generate random number within specified range
```go
// Good for mobile verification challenges
i := randtool.GenIntRange(1000, 9999)
```


**Benchmarks**
```cmd
go test -bench=.
```
```go
BenchmarkGenInt64	 2000000	       715 ns/op	      16 B/op	       2 allocs/op
BenchmarkGenStr  	10000000	       235 ns/op	      32 B/op	       2 allocs/op
```
