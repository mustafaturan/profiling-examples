package version1

import (
	"fmt"
	"math/rand"
	"strconv"
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
	return toBase62WithPaddingZeros(t, 8) + toBase62WithPaddingZeros(s, 4) + toBase62WithPaddingZeros(n, 4)
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

// toBase62 converts int types to Base62 encoded string
func toBase62(u uint64) string {
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
	return string(a[i:])
}

// toBase62WithPaddingZeros converts int types to Base62 encoded string with
// padding zeros
func toBase62WithPaddingZeros(u uint64, padding int64) string {
	formatter := "%+0" + strconv.FormatInt(padding, 10) + "s"
	return fmt.Sprintf(formatter, toBase62(u))
}
