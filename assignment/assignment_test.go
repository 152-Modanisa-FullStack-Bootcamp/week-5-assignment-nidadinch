package assignment

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {

	testCases := []struct {
		name     string
		first    uint32
		second   uint32
		sum      uint32
		overflow bool
	}{
		/*
			Sum uint32 numbers, return uint32 sum value and boolean overflow flag
			cases need to pass:
				math.MaxUint32, 1 => 0, true
				1, 1 => 2, false
				42, 2701 => 2743, false
				42, math.MaxUint32 => 41, true
				4294967290, 5 => 4294967295, false
				4294967290, 6 => 0, true
				4294967290, 10 => 4, true
		*/
		{"testCase1", math.MaxUint32, 1, 0, true},
		{"testCase2", 1, 1, 2, false},
		{"testCase3", 42, 2701, 2743, false},
		{"testCase4", 42, math.MaxUint32, 41, true},
		{"testCase5", 4294967290, 5, 4294967295, false},
		{"testCase6", 4294967290, 6, 0, true},
		{"testCase7", 4294967290, 10, 4, true},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			sum, overflow := AddUint32(tt.first, tt.second)
			assert.Equal(t, tt.sum, sum)
			assert.Equal(t, tt.overflow, overflow)
		})
	}
}

func TestCeilNumber(t *testing.T) {

	tests := []struct {
		name string
		f    float64
		want float64
	}{
		/*
			Ceil the number within 0.25
			cases need to pass:
				42.42 => 42.50
				42 => 42
				42.01 => 42.25
				42.24 => 42.25
				42.25 => 42.25
				42.26 => 42.50
				42.55 => 42.75
				42.75 => 42.75
				42.76 => 43
				42.99 => 43
				43.13 => 43.25
		*/
		{"testCase01", 42.42, 42.50},
		{"testCase02", 42, 42},
		{"testCase03", 42.01, 42.25},
		{"testCase04", 42.24, 42.25},
		{"testCase05", 42.25, 42.25},
		{"testCase06", 42.26, 42.50},
		{"testCase07", 42.55, 42.75},
		{"testCase08", 42.75, 42.75},
		{"testCase09", 42.76, 43},
		{"testCase10", 42.99, 43},
		{"testCase11", 43.13, 43.25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			point := CeilNumber(tt.f)
			assert.Equal(t, tt.want, point)
		})
	}
}

func TestAlphabetSoup(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want string
	}{
		/*
			String with the letters in alphabetical order.
			cases need to pass:
			 	"hello" => "ehllo"
				"" => ""
				"h" => "h"
				"ab" => "ab"
				"ba" => "ab"
				"bac" => "abc"
				"cba" => "abc"
		*/
		{"testCase01", "hello", "ehllo"},
		{"testCase02", "", ""},
		{"testCase03", "h", "h"},
		{"testCase04", "ab", "ab"},
		{"testCase05", "ba", "ab"},
		{"testCase06", "bac", "abc"},
		{"testCase07", "cba", "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AlphabetSoup(tt.s)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestStringMask(t *testing.T) {

	tests := []struct {
		name  string
		str   string
		count int
		want  string
	}{
		/*
			Replace after n(uint) character of string with '*' character.
			cases need to pass:
				"!mysecret*", 2 => "!m********"
				"", n(any positive number) => "*"
				"a", 1 => "*"
				"string", 0 => "******"
				"string", 3 => "str***"
				"string", 5 => "strin*"
				"string", 6 => "******"
				"string", 7(bigger than len of "string") => "******"
				"s*r*n*", 3 => "s*r***"
		*/
		{"testCase01", "!mysecret*", 2, "!m********"},
		{"testCase02", "", 1, "*"},
		{"testCase03", "a", 1, "*"},
		{"testCase04", "string", 0, "******"},
		{"testCase05", "string", 3, "str***"},
		{"testCase06", "string", 5, "strin*"},
		{"testCase07", "string", 7, "******"},
		{"testCase08", "s*r*n*", 3, "s*r***"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringMask(tt.str, uint(tt.count))
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"

	tests := []struct {
		name     string
		givenStr string
		wantStr  string
	}{
		/*
			Your goal is to determine if the first element in the array can be split into two words,
			where both words exist in the dictionary(words variable) that is provided in the second element of array.
			cases need to pass:
				[2]string{"hellocat",words} => hello,cat
				[2]string{"catbat",words} => cat,bat
				[2]string{"yellowapple",words} => yellow,apple
				[2]string{"",words} => not possible
				[2]string{"notcat",words} => not possible
				[2]string{"bootcamprocks!",words} => not possible
		*/
		{"testCase01", "hellocat", "hello,cat"},
		{"testCase02", "catbat", "cat,bat"},
		{"testCase03", "yellowapple", "yellow,apple"},
		{"testCase04", "", "not possible"},
		{"testCase05", "notcat", "not possible"},
		{"testCase06", "bootcamprocks!", "not possible"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WordSplit([2]string{tt.givenStr, words})
			assert.Equal(t, tt.wantStr, result)
		})
	}
}

func TestVariadicSet(t *testing.T) {

	tests := []struct {
		name string
		list []interface{}
		want []interface{}
	}{
		/*
			FINAL BOSS ALERT :)
			Tip: Learn and apply golang variadic functions(search engine -> "golang variadic function" -> WOW You can really dance! )
			Convert inputs to set(no duplicate element)
			cases need to pass:
				4,2,5,4,2,4 => []interface{4,2,5}
				"bootcamp","rocks!","really","rocks!"=> []interface{"bootcamp","rocks!","really"}
				1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first" => []interface{1,uint32(1),"first",2,uint32(2),"second"}
		*/
		{"TestCase01", []interface{}{4, 2, 5, 4, 2, 4}, []interface{}{4, 2, 5}},
		{"TestCase02", []interface{}{"bootcamp", "rocks!", "really", "rocks!"}, []interface{}{"bootcamp", "rocks!", "really"}},
		{"TestCase03", []interface{}{1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"}, []interface{}{1, uint32(1), "first", 2, uint32(2), "second"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := VariadicSet(tt.list...)
			assert.Equal(t, tt.want, set)
		})
		t.Run(tt.name, func(t *testing.T) {
			set := VariadicSet2(tt.list...)
			assert.Equal(t, tt.want, set)
		})
	}
}

//BenchmarkVariadicSet-4           3125035               393.6 ns/op           240 B/op          4 allocs/op
func BenchmarkVariadicSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		VariadicSet(1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first")
	}
}

//BenchmarkVariadicSet2-4          1884538               576.6 ns/op           408 B/op          4 allocs/opfunc BenchmarkVariadicSet2(b *testing.B) {
func BenchmarkVariadicSet2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		VariadicSet2(1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first")
	}
}
