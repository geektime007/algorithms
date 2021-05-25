package main

import "fmt"

func bubbleSort(nums []int) []int {
	length := len(nums)
	// 盖层循环控制 需要冒泡的轮数
	//for i := 0; i < length-1; i++ {
	//	for j := 1; j < length; j++ {
	//		if nums[i] > nums[j] {
	//			nums[i], nums[j] = nums[j], nums[i]
	//		}
	//	}
	//}

	for i := 1; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func main() {
	fmt.Println("vim-go")
	a := []int{2, 4, 1, 8, 3, 5, 10}
	fmt.Println(a)
	b := bubbleSort(a)
	fmt.Println(b)
}
