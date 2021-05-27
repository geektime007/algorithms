package main

import "fmt"

// 方法一
func binarySearch(nums []int, val int) int {
	length := len(nums)
	if length <= 0 {
		return -1
	}
	left := 0
	right := length - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < val {
			left = mid + 1
		} else if nums[mid] > val {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// 方法二
func BinarySearch(nums []int, val int) int {
	length := len(nums)
	if length <= 0 {
		return -1
	}
	left := 0
	right := length
	for left < right {
		mid := left + ((right - left) >> 1)
		if nums[mid] < val {
			left = mid + 1
		} else if nums[mid] > val {
			right = mid
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	fmt.Println("vim-go")
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	b := binarySearch(a, 5)
	fmt.Println(b)

	c := BinarySearch(a, 6)
	fmt.Println(c)
}
