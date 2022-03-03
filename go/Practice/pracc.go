package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(arr)

	// Move the pivot to the right
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Pile elemenarr smaller than the pivot on the left
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot after the lrrayast smaller element
	arr[left], arr[right] = arr[right], arr[left]

	// Go down the rabbit hole
	quickSort(arr[:left])
	quickSort(arr[left+1:])
	//fmt.Println(arr)
	return arr

}

// compute closest number to 0
func computeClosest(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	s := make([]int, 0)
	for _, v := range arr {
		s = append(s, int(math.Abs(float64(v))))
	}
	sortedS := quickSort(s)

	if sortedS[0] >= 0 {
		fmt.Println(sortedS[0])
		return sortedS[0]

	}
	return -1
}

func removeDupes(arr []int) []int {

	set := make(map[int]bool)
	newSet := make([]int, 0)
	for _, num := range arr {
		if _, value := set[num]; !value {
			set[num] = true
			newSet = append(newSet, num)
		}
	}
	return newSet
}

func isSubsequence(s, t string) bool {

	i, j := 0, 0
	for i < len(s) && j < len(t) {
		for ; j < len(t); j++ {
			if s[i] == t[j] {
				i++
				j++
				break
			}
		}
	}
	return i == len(s) && j <= len(t)

}

func twoDigits(nums int) bool {

	newNums := strconv.Itoa(nums)
	myRunes := make([]int, 0)
	for i, num := range newNums {
		fmt.Printf("Loop %d: num: %v\n", i, num)
		myRunes = append(myRunes, int(num))
	}
	fmt.Println("MyRunes:", myRunes)
	newSet := removeDupes(myRunes)

	fmt.Println("New Set: ", newSet)
	if len(newSet) > 2 {
		fmt.Println("False")
		return false
	}
	fmt.Println("True")
	return true
}

func numOfArithSlices(nums []int) int {

	if len(nums) < 3 {
		return 0
	}
	dp, result := 0, 0
	for i := 2; i < len(nums); i++ {
		diff := (nums[i] - nums[i-1]) - (nums[i-1] - nums[i-2])
		if diff == 0 {
			dp++
			result += dp
		} else {
			dp = 0
		}
	}
	return result
}

func main() {

	/* arr := []int{-2, -7, -2, 12, 10, 4, 2, 18, 11, 4}
	computeClosest(arr)
	removeDupes(arr) */
	s, t := "aaa", "ahbgdc"
	isSubsequence(s, t)
	nums := []int{1, 2, 3, 4}
	numOfArithSlices(nums)
	/* duo := 1001
	trio := 123
	twoDigits(duo)
	twoDigits(trio) */
}
