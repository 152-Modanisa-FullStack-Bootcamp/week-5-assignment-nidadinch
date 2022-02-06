package assignment

import (
	"math"
	"reflect"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	z := x + y
	if x != 0 && y != 0 && x > math.MaxUint32-y {
		return z, true
	}
	return z, false
}

func CeilNumber(f float64) float64 {
	x := math.Ceil(f*4) / 4
	return x
}

func AlphabetSoup(s string) string {
	var resultStr string
	strRune := []rune(s)
	// bubble sort
	for i := 0; i < len(strRune); i++ {
		for j := 1; j < len(strRune)-i; j++ {
			if strRune[j] < strRune[j-1] {
				strRune[j], strRune[j-1] = strRune[j-1], strRune[j]
			}
		}
	}
	for _, v := range strRune {
		resultStr += string(v)
	}
	return resultStr
}

func StringMask(s string, n uint) string {
	var resultStr string
	str1 := s[:n]
	str2 := s[n:]

	if len(s) == 0 {
		return "*"
	}

	if int(n) == len(s) || int(n) > len(s) {
		for i := 0; i < len(s); i++ {
			resultStr += "*"
		}
		return resultStr
	}

	for i := 0; i < len(str2); i++ {
		resultStr += "*"
	}

	return str1 + resultStr
}

func WordSplit(arr [2]string) string {
	firstStr := arr[0]
	words := strings.Split(arr[1], ",")
	resultStr := ""

	for _, s := range words {
		if index := strings.Index(firstStr, s); index != -1 {
			splitted1 := firstStr[:index]
			splitted2 := firstStr[index:]

			if strings.Contains(arr[1], splitted1) {
				resultStr = splitted1 + "," + splitted2
				return resultStr
			}
		}
	}
	return "not possible"
}

// Implementation for Final Boss using Reflect
func VariadicSet2(i ...interface{}) []interface{} {
	count := 0
	s := reflect.ValueOf(i)
	result := make([]interface{}, s.Len())

	for j := 0; j < s.Len(); j++ {
		if !Contains(result, s.Index(j).Interface()) {
			result[j] = s.Index(j).Interface()
			count++
		}
	}
	// I have created another slice to get rid of nil values
	// Because 'result' contains nil values
	finalRes := make([]interface{}, count)
	copy(finalRes, result)
	return finalRes
}

// Implementation for Final Boss using traditional method
func VariadicSet(i ...interface{}) []interface{} {
	return nil
}
