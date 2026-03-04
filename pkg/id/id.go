// Package id provides UUID generation helpers.
package id

import "crypto/rand"

// New generates a new random UUID v4 string.
func New() string {
	var uuid [16]byte
	_, _ = rand.Read(uuid[:])
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // variant 2

	var buf [36]byte
	hexEncode(buf[:], uuid[:])
	return string(buf[:])
}

func hexEncode(dst []byte, src []byte) {
	const hextable = "0123456789abcdef"
	pos := 0
	for i, b := range src {
		if i == 4 || i == 6 || i == 8 || i == 10 {
			dst[pos] = '-'
			pos++
		}
		dst[pos] = hextable[b>>4]
		dst[pos+1] = hextable[b&0x0f]
		pos += 2
	}
}
