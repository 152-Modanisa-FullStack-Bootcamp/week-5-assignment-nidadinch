package assignment

import "math"

func AddUint32(x, y uint32) (uint32, bool) {
	z := x + y
	if x != 0 && y != 0 && x > math.MaxUint32-y {
		return z, true
	}
	return z, false
}

func CeilNumber(f float64) float64 {
	return 0
}

func AlphabetSoup(s string) string {
	return ""
}

func StringMask(s string, n uint) string {
	return ""
}

func WordSplit(arr [2]string) string {
	return ""
}

func VariadicSet(i ...interface{}) []interface{} {
	return nil
}
