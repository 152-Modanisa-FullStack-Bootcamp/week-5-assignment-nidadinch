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
