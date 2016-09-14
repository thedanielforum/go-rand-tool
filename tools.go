// Package randtool is for generating common types of pseudo random data.
package randtool

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	mathrand "math/rand"
	"sync"
)

const (
	// The available alphabet for GenerateAlphaString()
	chars   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// 6 bits to represent a letter index
	letterIdxBits = 6
	// All 1-bits, as many as letterIdxBits
	letterIdxMask = 1<<letterIdxBits - 1
	// # of letter indices fitting in 63 bits
	letterIdxMax  = 63 / letterIdxBits
)

var seeded bool = false
var mutex  sync.Mutex

// Pseudo-random int64 values in the range [0, 1<<63).
func int63() int64 {
	mutex.Lock()
	v := mathrand.Int63()
	mutex.Unlock()
	return v
}

// Generate func for creating a pseudo random int64 using crypto/rand
// This function is designed especially for the seeding rand.Seed()
func GenInt64() int64 {
	var i int64
	// On Unix-like systems, Reader reads from /dev/urandom.
	// On Linux, Reader uses getrandom(2) if available, /dev/urandom otherwise.
	// On Windows systems, Reader uses the CryptGenRandom API.
	err := binary.Read(rand.Reader, binary.LittleEndian, &i)
	if err != nil {
		panic(fmt.Sprintf("Can not read crypto/rand lib: %s", err.Error()))
	}
	return i
}

// Seed math/rand using a pseudo random int64 value
func SeedMathRand() {
	mathrand.Seed(GenInt64())
	// Set math.Seed status
	seeded = true
}

// Generate a url safe pseudo random alphabetic string of N length
// Credits goes to (icza) http://stackoverflow.com/a/31832326/5315198
func GenStr(n int) (string, error) {
	if n < 1 {
		return "", errors.New("Random string length must be greater than 0")
	}
	// Don't seed on every iteration to increase performance
	// and reduce chance for reseeding with same int64
	if !seeded {
		SeedMathRand()
	}

	b := make([]byte, n)
	// math_rand.Int63() generates 63 random bits, enough for letterIdxMax characters
	for i, cache, remain := n-1, int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(chars) {
			b[i] = chars[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b), nil
}

// Returns the value of GenerateAlpha with the error ignored
// Use with caution
func GenStrIgnoreErr(n int) string {
	s, _ := GenStr(n)
	return s
}
