package utils

import (
	"reflect"
	"strconv"
	"strings"
)

type Int interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// Ints2String convert int based array to sep joined string
func Ints2String[T Int](ints []T, sep string) string {
	if len(ints) == 0 {
		return ""
	}
	var sb strings.Builder
	t := reflect.TypeOf(ints[0])
	sb.Grow(len(ints)*int(t.Size()) + len(sep)*len(ints))
	for i, t2 := range ints {
		if t2 < 0 {
			s := strconv.FormatInt(int64(t2), 10)
			sb.WriteString(s)
		} else {
			s := strconv.FormatUint(uint64(t2), 10)
			sb.WriteString(s)
		}
		if i != len(ints)-1 {
			sb.WriteString(sep)
		}
	}
	return sb.String()
}
