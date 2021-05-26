package main

import "fmt"

// 冒泡排序
// 思路分析：在要排序的切片中，对当前还未排好的序列，
// 从前往后对相邻的两个元素依次进行比较和调整，
// 让较大的数往下沉，较小的往上冒。即，每当两相邻的
// 元素比较后发现它们的排序与排序要求相反时，就将它们互换。

func bubbleSort(nums []int) []int {
	length := len(nums)
	// 盖层循环控制 需要冒泡的轮数
	for i := 0; i < length; i++ {
		//该层循环用来控制每轮 冒出一个数 需要比较的次数
		for j := 0; j < length-1; j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func bubbleSort1(nums []int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func bubbleSort2(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
	return nums
}

func main() {
	fmt.Println("vim-go")
	a := []int{2, 4, 1, 8, 3, 5, 10}
	fmt.Println(a)
	b := bubbleSort2(a)
	fmt.Println(b)
}
