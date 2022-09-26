package utils

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntArray2String(t *testing.T) {
	// int
	ints1 := []int{1, 2, 3, 4}
	sep1 := ","
	want1 := "1,2,3,4"
	got1 := Ints2String(ints1, sep1)
	assert.Equal(t, want1, got1)

	// uint
	ints2 := []uint{1, 2, 3, 4}
	sep2 := "||"
	expect2 := "1||2||3||4"
	got2 := Ints2String(ints2, sep2)
	assert.Equal(t, expect2, got2)

	// int64
	ints3 := []int64{1, 2, 3, 4}
	sep3 := "~"
	expect3 := "1~2~3~4"
	got3 := Ints2String(ints3, sep3)
	assert.Equal(t, expect3, got3)

	// empty
	var ints4 []uint
	expect4 := ""
	assert.Equal(t, expect4, Ints2String(ints4, sep3))

	maxInts := []int64{math.MaxInt64, math.MinInt64}
	expectMaxInts := fmt.Sprintf("%d,%d", math.MaxInt64, math.MinInt64)
	assert.Equal(t, expectMaxInts, Ints2String(maxInts, ","))
}

func BenchmarkIntArray2String(b *testing.B) {
	ints := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < b.N; i++ {
		Ints2String(ints, ",")
	}
}
