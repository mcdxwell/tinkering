package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

func quickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	left, right := 0, len(array)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(array)

	// Move the pivot to the right
	array[pivotIndex], array[right] = array[right], array[pivotIndex]

	// Pile elemenarray smaller than the pivot on the left
	for i := range array {
		if array[i] < array[right] {
			array[i], array[left] = array[left], array[i]
			left++
		}
	}

	// Place the pivot after the lrrayast smaller element
	array[left], array[right] = array[right], array[left]

	// Go down the rabbit hole
	quickSort(array[:left])
	quickSort(array[left+1:])
	//fmt.Println(array)
	return array

}

// compute closest number to 0
func ComputeClosest(array []int) int {

	if len(array) == 0 {
		return 0
	}

	if len(array) == 1 {
		return array[0]
	}

	s := make([]int, 0)
	for _, v := range array {
		s = append(s, int(math.Abs(float64(v))))
	}
	sortedS := quickSort(s)

	if sortedS[0] >= 0 {
		fmt.Println(sortedS[0])
		return sortedS[0]

	}
	return -1
}

func removeDupes(array []int) []int {

	keys := make(map[int]bool)
	list := make([]int, 0)

	for _, num := range array {
		if _, value := keys[num]; !value {
			keys[num] = true
			list = append(list, num)
		}
	}
	fmt.Println(list)
	return list
}

func isSubsequence(s, t string) bool {

	//myChars := make([]string, 0)

	for i := 0; i < len(t); i++ {
		fmt.Println(t[i])
	}

	return strings.ContainsAny(s, t)
}

func twoDigits(nums int) bool {

	mynums := strconv.Itoa(nums)
	checkers := make([]int, 0)
	for i, char := range mynums {
		fmt.Printf("Loop %d: char%v\n", i, char)
		checkers = append(checkers, int(char))
		//check_nums := strconv.Atoi()
	}

	set := make(map[int]bool)
	newSet := make([]int, 0)
	for _, num := range checkers {
		if _, value := set[num]; !value {
			set[num] = true
			newSet = append(newSet, num)
		}
	}
	fmt.Println("New Set: ", newSet)
	if len(newSet) > 2 {
		fmt.Println("False")
		return false
	}
	fmt.Println("True")
	return true
}

func main() {

	array := []int{-2, -7, -2, 12, 10, 4, 2, 18, 11, 4}
	ComputeClosest(array)
	removeDupes(array)
	s, t := "cba", "ahbgdc"
	isSubsequence(s, t)

	duo := 1001
	trio := 123
	twoDigits(duo)
	twoDigits(trio)
}
