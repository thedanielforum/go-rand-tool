![godoc](https://travis-ci.org/thedanielforum/randtool.svg?branch=master) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/676a9e1aabe14027ae63a941312c399d)](https://www.codacy.com/app/thedanielforum/randtool?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=thedanielforum/randtool&amp;utm_campaign=Badge_Grade)

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

Avalible integers
```go
GenInt64() int64
GenInt32() int32
GenInt16() int16
GenInt8() int8
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
