package util

import (
	"bytes"
	"fmt"
	"strconv"
)

func NumberToString(n int, sep rune) string {

	s := strconv.Itoa(n)

	startOffset := 0
	var buff bytes.Buffer

	if n < 0 {
		startOffset = 1
		buff.WriteByte('-')
	}

	l := len(s)

	commaIndex := 3 - ((l - startOffset) % 3)

	if commaIndex == 3 {
		commaIndex = 0
	}

	for i := startOffset; i < l; i++ {

		if commaIndex == 3 {
			buff.WriteRune(sep)
			commaIndex = 0
		}
		commaIndex++

		buff.WriteByte(s[i])
	}

	return buff.String()
}

func SecondsToMinutes(inSeconds int) string {
	minutes := inSeconds / 60
	seconds := inSeconds % 60

	var s string
	if seconds < 10 {
		s = fmt.Sprintf("%d%d", 0, seconds)
	} else {
		s = fmt.Sprintf("%d", seconds)
	}

	str := fmt.Sprintf("%d:%s", minutes, s)
	return str
}
