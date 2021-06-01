package main

import "fmt"

// TopK问题 快速排序发解决方式

func QuickSearch(nums []int, k int) []int {
	length := len(nums)
	if length < 2 || k > length {
		return nums
	}
	quickSort(nums, 0, length-1, k-1)
	return nums[:k]
	//return nums
}

func quickSort(nums []int, left int, right int, k int) {
	if left >= right {
		return
	}
	pivot := partition(nums, left, right)
	fmt.Println(pivot)
	if pivot == k {
		fmt.Println("Return", k)
		return
	}
	if pivot < k {
		// 只需要分类 pivot+1 到 right就可以了
		quickSort(nums, pivot+1, right, k)
	} else {
		// 从最左边到pivot - 1 再次排序
		quickSort(nums, left, pivot-1, k)
	}
}

func partition(nums []int, left int, right int) int {
	mid := nums[left]
	for left < right {
		for left < right && mid <= nums[right] {
			right--
		}
		nums[left] = nums[right]
		fmt.Println("-", nums)
		for left < right && mid >= nums[left] {
			left++
		}
		nums[right] = nums[left]
		fmt.Println("=", nums)
	}
	nums[left] = mid
	return left
}

func main() {
	fmt.Println("vim-go")
	a := []int{1223, 4, 5, 12, 5, 6, 7, 7, 32, 125, 6, 7, 9}
	fmt.Println(a)
	b := QuickSearch(a, 5)
	fmt.Println(b)

	a1 := []int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}
	fmt.Println(a1)
	b1 := QuickSearch(a1, 5)
	fmt.Println(b1)
}
