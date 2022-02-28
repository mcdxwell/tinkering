package main

import (
	"fmt"
	"math"
	"testing"
)

// idea for testing:
//
// provide input
// get output
// take the reverse function and pass the output to reproduce the input
// due to rounding, the best way to test accuracy would be to take the difference
// if ther difference is marginal => test pass

var tests = []struct {
	temp float64
}{
	{-236},
	{0},
	{273.15},
}

// Test:
// Celsius => Fahrenheit
// Fahrenheit => Celsius
func TestCFFC(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %f", tt.temp)
		t.Run(testname, func(t *testing.T) {
			output := celsius{tt.temp}.Fahrenheit()
			input := fahrenheit{output}.Celsius()
			diff := math.Abs(tt.temp - input)
			if diff > 1 {
				t.Errorf("Failed!\n")
			}
			//fmt.Printf("CFFC-Input: %f\n Ouput: %f\n Difference: %f\n", tt.temp, output, diff)
		})
	}
}

// Test:
// Kelvin => Fahrenheit
// Fahrenheit => Kelvin
func TestKFFK(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %f", tt.temp)
		t.Run(testname, func(t *testing.T) {
			output := kelvin{tt.temp}.Fahrenheit()
			input := fahrenheit{output}.Kelvin()
			diff := math.Abs(tt.temp - input)
			if diff > 1 {
				t.Errorf("Failed!\n")
			}
			//fmt.Printf("KFFK-Input: %f\n Ouput: %f\n Difference: %f\n", tt.temp, output, diff)
		})
	}
}

// Test:
// Celsius => Kelvin
// Kelvin => Celsius
func TestCKKC(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("Input: %f", tt.temp)
		t.Run(testname, func(t *testing.T) {
			output := celsius{tt.temp}.Kelvin()
			input := kelvin{output}.Celsius()
			diff := math.Abs(tt.temp - input)
			if diff > 1 {
				t.Errorf("Failed!\n")
			}
			//fmt.Printf("CKKC-Input: %f\n Ouput: %f\n Difference: %f\n", tt.temp, output, diff)
		})
	}
}
