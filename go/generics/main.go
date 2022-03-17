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
	int
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

	moreInts := map[string]int{
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
