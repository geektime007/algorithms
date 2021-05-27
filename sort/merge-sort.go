package main

import "fmt"

func MergeSort(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}
	tmp := make([]int, length)
	mergeSort(nums, tmp, 0, length-1)
	return nums
}

func mergeSort(nums []int, tmp []int, start int, end int) {
	if start >= end {
		return
	}
	start1 := start
	end1 := start + (end-start)/2

	start2 := end1 + 1
	end2 := end

	mergeSort(nums, tmp, start1, end1)
	mergeSort(nums, tmp, start2, end2)

	MergeSortedArray(nums, start1, end1, nums, start2, end2, tmp, start)

	for i := start; i <= end; i++ {
		nums[i] = tmp[i]
	}
}

func MergeSortedArray(nums1 []int, start1 int, end1 int,
	nums2 []int, start2 int, end2 int, tmp []int, start int) {
	i := start1
	j := start2
	k := start
	for i <= end1 && j <= end2 {
		if nums1[i] < nums2[j] {
			tmp[k] = nums1[i]
			i++
		} else {
			tmp[k] = nums2[j]
			j++
		}
		k++
	}

	for i <= end1 {
		tmp[k] = nums1[i]
		k++
		i++
	}
	for j <= end2 {
		tmp[k] = nums1[j]
		k++
		j++
	}
}

func main() {
	a := []int{9, 2, 3, 4, 1, 2, 32, 12, 7, 6, 9, 11}
	fmt.Println(a)
	b := MergeSort(a)
	fmt.Println(b)
}
