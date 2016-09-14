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
var mutex sync.Mutex

// pseudo-random int64 values in the range [0, 1<<63).
func int63() int64 {
	mutex.Lock()
	v := mathrand.Int63()
	mutex.Unlock()
	return v
}

// Generate func for creating a pseudo random int64 using crypto/rand
// This function is designed especially for the seeding rand.Seed()
func GenerateSeed() int64 {
	var s int64
	err := binary.Read(rand.Reader, binary.LittleEndian, &s)
	if err != nil {
		panic(fmt.Sprintf("Can not read crypto/rand lib: %s", err.Error()))
	}
	seeded = true // Set math.Seed status
	return s
}

// Generate a url safe pseudo random alphabetic string of N length
// Credits goes to (icza) http://stackoverflow.com/a/31832326/5315198
func GenerateString(n int) (string, error) {
	if n < 1 {
		return "", errors.New("Random string length must be greater than 0")
	}
	// Seed math if not already done
	if !seeded {
		mathrand.Seed(GenerateSeed())
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
func GenerateStringIgnErr(n int) string {
	s, _ := GenerateString(n)
	return s
}
