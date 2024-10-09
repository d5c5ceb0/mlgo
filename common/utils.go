package common

import (
	"math"
	"os"
	"unsafe"

	"github.com/kshard/float8"
)

// NB! INT = 32 bits
func ReadInt32FromFile(file *os.File) uint32 {
	buf := make([]byte, 4)
	if count, err := file.Read(buf); err != nil || count != 4 {
		return 0
	}
	return uint32(buf[3])<<24 | uint32(buf[2])<<16 | uint32(buf[1])<<8 | uint32(buf[0])
}

func ReadStringFromFile(file *os.File, len uint32) string {
	buf := make([]byte, len)
	if count, err := file.Read(buf); err != nil || count != int(len) {
		return ""
	}
	return string(buf)
}

func ReadFP32FromFile(file *os.File) float32 {
	buf := make([]byte, 4)
	if count, err := file.Read(buf); err != nil || count != 4 {
		return 0.0
	}
	bits := uint32(buf[3])<<24 | uint32(buf[2])<<16 | uint32(buf[1])<<8 | uint32(buf[0])
	return math.Float32frombits(bits)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func DecodeFloat32List(bs []byte) []float8.Float8 {
	for n := 0; n < len(bs)/4; n++ {
		bits := uint32(bs[n*4+3])<<24 | uint32(bs[n*4])<<16 | uint32(bs[n*4+1])<<8 | uint32(bs[n*4])
		bs[n] = float8.ToFloat8(math.Float32frombits(bits))
	}

	return bs
}

func EncodeFloat32List(fs []float32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&fs[0])), len(fs)*4)
}
