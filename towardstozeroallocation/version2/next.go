package version2

import (
	"math/rand"
	"time"
)

const (
	maxBase62 = uint64(62)
	mapping   = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func Next() string {
	t := now()
	s := random()
	n := node()

	var id [16]byte
	copy(id[0:8], toBase62WithPaddingZeros(t, 8))
	copy(id[8:12], toBase62WithPaddingZeros(s, 4))
	copy(id[12:], toBase62WithPaddingZeros(n, 4))
	return string(id[:])
}

func random() uint64 {
	return uint64(rand.Intn(100_000))
}

func now() uint64 {
	t := time.Now().UnixNano() / int64(time.Millisecond)
	return uint64(t)
}

func node() uint64 {
	return 1
}

// toBase62WithPaddingZeros converts int types to Base62 encoded byte array
// with padding zeros
func toBase62WithPaddingZeros(u uint64, length int) []byte {
	const size = 11
	var a [size]byte
	i := size
	for u >= maxBase62 {
		i--
		// Avoid using r = a%b in addition to q = a/maxBase62
		// since 64bit division and modulo operations
		// are calculated by runtime functions on 32bit machines.
		q := u / maxBase62
		a[i] = mapping[u-q*maxBase62]
		u = q
	}
	// when u < maxBase62
	i--
	a[i] = mapping[u]
	for i > size-length {
		i--
		a[i] = mapping[0]
	}
	return a[i:]
}
