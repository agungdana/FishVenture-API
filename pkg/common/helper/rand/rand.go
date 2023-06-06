package rand

import (
	"math/rand"
	"strings"
	"time"
)

const byteList = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

var gen = rand.NewSource(time.Now().UnixNano())

func RandCode(length int) string {
	sbuilder := strings.Builder{}
	sbuilder.Grow(length)
	for value, max := gen.Int63(), 0; length > 0; length-- {
		if max == 12 {
			value, max = gen.Int63(), 0
		}
		sbuilder.WriteByte(byteList[int(value&0b111111)])
		value >>= 5
		max++
	}

	return sbuilder.String()
}
