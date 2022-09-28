package number

import (
	"fmt"
	"math"
	"strconv"
)

// RoundInt 将一个浮点数进行四舍五入
func RoundInt(num float64) int {
	return int(math.Trunc(num + 0.5))
}

// Round 按照精度四舍五入
func Round(num float64, precision uint) float64 {
	s := fmt.Sprintf(fmt.Sprintf("%%.%df", precision), num)
	f, _ := strconv.ParseFloat(s, 10)
	return f
}
