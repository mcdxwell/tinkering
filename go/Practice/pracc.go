package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// We'll use l1 to write the results in, so we need to store the head
	carry, head := 0, l1
	for {
		// Add the two nodes' values and the carry, write to l1
		l1.Val += l2.Val + carry

		// Check for carry and update l1's value
		carry = int(float64(l1.Val) * 0.1)
		l1.Val = l1.Val % 10 // or l1.Next.Val - (carry * 10)

		// As soon as either of the  runs out we're done with this loop
		if l2.Next == nil {
			break
		} else if l1.Next == nil {
			l1.Next = l2.Next
			break
		}

		// Advance!
		l1 = l1.Next
		l2 = l2.Next
	}

	// We'll now apply the carry to l1, adding nodes as needed
	for carry != 0 {
		// If there is a carry and l1.Next is null, add another node to the list
		if l1.Next == nil {
			l1.Next = &ListNode{0, nil}
		}

		// Same as the first loop, just using one list and the carry
		l1.Next.Val += carry

		carry = l1.Next.Val / 10
		l1.Next.Val = l1.Next.Val % 10

		// Advance!
		l1 = l1.Next
	}

	return head
}

// source: leet code - > fast
func longestSubStr(s string) int {

	chars := []rune(s)
	output, start := 0, 0
	chMap := make(map[rune]int)
	for i, ch := range chars {
		if value, exist := chMap[ch]; exist && value >= start {
			//chMap = make(map[rune]int)
			chMap[ch] = i
			start = value + 1
		} else {
			chMap[ch] = i
			length := i - start + 1
			if output < length {
				output = length
			}
		}
	}
	fmt.Println(output)
	return output
}

// source: leetcode -> insanely fast and memory efficient
func longestSubStrr(s string) int {
	freqarr := make([]int, 95)
	maxlen := 0
	currlen := 0
	startidx := 1
	roonies := make([]rune, 0)
	intoonies := make([]int, 0)
	for i, r := range s {
		fmt.Println(i)
		roonies = append(roonies, r)
		intoonies = append(intoonies, int(r))
		idx := int(r) - 32

		if freqarr[idx] >= startidx {
			startidx = freqarr[idx] + 1
		}
		fmt.Println("Freqarr before:", freqarr[idx])
		freqarr[idx] = i + 1
		fmt.Println("Freqarr after:", freqarr[idx])
		currlen = i + 1 - startidx + 1
		if currlen > maxlen {
			maxlen = currlen
		}
	}

	fmt.Println(roonies)
	fmt.Println(intoonies)
	fmt.Println(freqarr)
	fmt.Println(maxlen)
	return maxlen
}

// 1338 reduce array size to half
// return minimum size set so at least half...
func halfArray(arr []int) int {

	/* myMap := make(map[int]int)

	for _, v := range arr {
		// if a key in my map does or does not have a count
		// 		increment count by 1
		if _, count := myMap[v]; !count || count {
			myMap[v] += 1
		}
	}

	//i := 0
	//limit := len(arr) / 2
	numbers := make([]int, 0)
	counts := make([]int, 0)
	for key, val := range myMap {
		numbers = append(numbers, key)
		counts = append(counts, val)
	}
	*/

	setSize := 0 // the elements removed from the array create the set

	myMap := make(map[int]int)
	arrSize := len(arr)
	halfArr := arrSize / 2

	for _, v := range arr {
		myMap[v]++
	}

	elmntCounts := make([]int, 0)
	for _, v := range myMap {
		elmntCounts = append(elmntCounts, v)
	}

	sort.Ints(elmntCounts)
	sumOfElements := 0
	for i := len(elmntCounts) - 1; i >= 0; i-- {
		if sumOfElements >= halfArr {
			return setSize
		}
		sumOfElements += elmntCounts[i]
		setSize++
	}
	return setSize
}

func minSetSize(arr []int) int {
	counts := getCounts(arr)
	target := len(arr) / 2
	totalRemoved := 0

	for i := range counts {
		totalRemoved += counts[i]
		if totalRemoved >= target {
			return i + 1
		}
	}
	return 0
}

func getCounts(arr []int) []int {
	myMap := make(map[int]int)
	for _, v := range arr {
		myMap[v]++
	}

	counts := make([]int, 0)

	for _, count := range myMap {
		counts = append(counts, count)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	return counts
}

// basic
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	mergedArray := make([]int, 0)
	for i := range nums1 {
		mergedArray = append(mergedArray, nums1[i])
	}
	for i := range nums2 {
		mergedArray = append(mergedArray, nums2[i])
	}
	sort.Ints(mergedArray)
	fmt.Println(mergedArray)

	arrSize := len(mergedArray)
	if arrSize%2 == 0 {
		nIndex := arrSize / 2
		mIndex := nIndex - 1
		adding := mergedArray[nIndex] + mergedArray[mIndex]
		fmt.Println(adding)
		median := float64(adding) * 0.5
		fmt.Println("Even: ", median)
		return float64(median)
	} else {
		index := int(math.Floor(float64(arrSize) / 2))
		median := float64(mergedArray[index])
		fmt.Println("Odd: ", median)
		return median
	}
}

func deleteAndEarn(nums []int) int {
	m := nums[0]
	for _, v := range nums {
		if v > m {
			m = v
		}
	}

	n := m + 1
	values := make([]int, n)
	fmt.Println("Values: ", values)
	for _, num := range nums {
		fmt.Println("Num:", num)
		values[num] += num
		fmt.Println("values[num]: ", values[num])
	}
	fmt.Println(values)
	var take, skip int
	for i := 0; i < n; i++ {
		fmt.Println("First: ", take, skip)
		takei := skip + values[i]
		skipi := max(skip, take)
		take = takei
		skip = skipi
		fmt.Println("Second: ", take, skip)
	}
	fmt.Println(max(take, skip))
	return max(take, skip)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	left, max

	ss := []rune(s)
	i, j := 0, 0

	for range s {
		ss[i]
	}
}

func isPalindrome(s []rune) bool {

	roons := make([]rune, 0)
	for i := len(s) - 1; i >= 0; i-- {
		roons = append(roons, s[i])
	}

	if string(s) == string(roons) {
		fmt.Println("true")
		return true
	} else {
		fmt.Println("false")
		return false
	}
}

func main() {

	/* arr := []int{-2, -7, -2, 12, 10, 4, 2, 18, 11, 4}
	computeClosest(arr)
	removeDupes(arr)
	s, t := "aaa", "ahbgdc"
	isSubsequence(s, t)
	nums := []int{1, 2, 3, 4}
	numOfArithSlices(nums)
	l1 := new(ListNode)
	l2 := new(ListNode)
	for i := 0; i <= 10; i++ {
		l1.Val = i
		l2.Val = i
		l1 = l1.Next
		l2 = l2.Next
	}
	addTwoNumbers(l1, l2)
	duo := 1001
	trio := 123
	twoDigits(duo)
	twoDigits(trio) */

	//longestSubStr(s)
	//longestSubStrr(ss)
	//longestSubStrr(s)

	//array := []int{1, 2}
	//array2 := []int{3, 4}
	//findMedianSortedArrays(array, array2)
	//halfArray(array)
	//halfArray(array)
	//array3 := []int{3, 4, 2}
	// array4 := []int{2, 2, 3, 3, 3, 4}
	// array5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// array6 := []int{5, 5, 5, 5, 5, 4, 6, 7}

	//deleteAndEarn(array3)
	// deleteAndEarn(array4)
	// deleteAndEarn(array5)
	// deleteAndEarn(array6)
	s := "babad"
	s = "bab"
	isPalindrome([]rune(s))
}
