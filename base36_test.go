package base36

import (
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var raw = []uint64{0, 50, 100, 999, 1000, 1111, 5959, 99999,
	123456789, 5481594952936519619, math.MaxInt64 / 2048, math.MaxInt64 / 512,
	math.MaxInt64, math.MaxUint64}

var encoded = []string{"0", "1E", "2S", "RR", "RS", "UV", "4LJ", "255R",
	"21I3V9", "15N9Z8L3AU4EB", "18CE53UN18F", "4XDKKFEK4XR",
	"1Y2P0IJ32E8E7", "3W5E11264SGSF"}

func TestEncode(t *testing.T) {
	for i, v := range raw {
		assert.Equal(t, encoded[i], Encode(v))
	}
}

func TestDecode(t *testing.T) {
	for i, v := range encoded {
		assert.Equal(t, raw[i], Decode(v))
		assert.Equal(t, raw[i], Decode(strings.ToLower(v)))
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(5481594952936519619)
	}
}

func BenchmarkEncodeBytesAsBytes(b *testing.B) {
	data := []byte{
		0x86, 0x4F, 0xD2, 0x6F, 0xB5, 0x59, 0xF7, 0x5B, 0x48, 0x4F, 0x2A, 0x48, 0x4F, 0x2A,
	}
	for i := 0; i < b.N; i++ {
		EncodeBytesAsBytes(data)
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode("1Y2P0IJ32E8E7")
	}
}
