package main

import "fmt"

func insertSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		tmp := nums[i]
		for j := i - 1; j >= 0; j-- {
			if tmp < nums[j] {
				nums[j+1], nums[j] = nums[j], tmp
			} else {
				break
			}
		}
	}
	return nums
}

func main() {
	fmt.Println("vim-go")
	a := []int{2, 4, 1, 8, 3, 5, 10}
	fmt.Println(a)
	b := insertSort(a)
	fmt.Println(b)
}
