package main

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

func SumInts(m map[string]int64) int64 {
	var s int64

	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// Generic Function

// SumIntsOrFloats sums the values of map
func SumIntsOrFloats[k comparable, V Number](m map[k]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Num interface {
	int64
}

func AddOddGrades[k comparable, V Num](m map[k]V) []V {
	var evenSum V
	var oddSum V
	for _, v := range m {
		if v%2 != 0 {
			oddSum += v
		} else {

			evenSum += v
		}
	}
	var ff []V
	ff = append(ff, oddSum)
	ff = append(ff, evenSum)

	return ff
}

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {

	ints := map[string]int64{
		"first":  10,
		"second": 12,
		"third":  13,
	}

	floats := map[string]float64{
		"first":  10.97,
		"second": 12.95,
		"third":  13.94,
	}

	moreInts := map[string]int64{
		"first":   10,
		"second":  12,
		"third":   13,
		"fourth":  13,
		"fifth":   13,
		"sixth":   13,
		"seventh": 13,
		"eighth":  13,
	}

	fmt.Printf("Non-generic sums: %v and %v\nGeneric sums: %v and %v\n", SumInts(ints), SumFloats(floats), SumIntsOrFloats(ints), SumIntsOrFloats(floats))
	fmt.Printf("Generic odd sums: %v Even sums: \n", AddOddGrades(moreInts))

	roonies := []rune{97, 20, 12, 12, 12, 13, 97, 10, 10, 1, 2, 3, 4, 4, 5}
	fmt.Println(Set(roonies))

	fmt.Println(stuff(10, 12, min))
	fmt.Println(stuff(10, 12, max))
	fmt.Println(stuff(10, 10, max))
	hhh := []int{1, 2, 3}
	jjj := []int{1, 2, 3, 4}
	oo, ii := twoLen(hhh, jjj)
	fmt.Println(stuff(oo, ii, max))

	fmt.Println(manyLen(jjj, hhh))
	fmt.Println(manyLen(moreInts, ints))
}

// Removes duplicates
func Set[T any](s []T) []T {
	keys := make(map[any]bool)
	list := []T{}
	for _, k := range s {
		if _, v := keys[k]; !v {
			keys[k] = true
			list = append(list, k)
		}
	}
	return list
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type rawr func(int, int) int

// func take two ints and rawr func
func stuff(a, b int, f rawr) int {
	return f(a, b)
}

type inters interface {
	[]uint | []int | []int8 | []int16 | []int64
}

type floaters interface {
	[]float32 | []float64
}

type mappers interface {
	map[int]int | map[string]string | map[string]int64
}

type thingers interface {
	[]rune | inters | floaters | mappers
}

// Return two lengths
func twoLen[T thingers](a T, b T) (int, int) {
	return len(a), len(b)
}

// Return a slice of lengths
func manyLen[T thingers](arrs ...T) []int {
	lengths := make([]int, 0)
	for _, arr := range arrs {
		arrLen := len(arr)
		lengths = append(lengths, arrLen)
	}
	return lengths
}
