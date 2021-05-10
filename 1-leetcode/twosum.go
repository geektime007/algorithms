package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indexs := []int{}
	if len(nums) == 0 {
		return nil
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				indexs = append(indexs, i)
				indexs = append(indexs, j)
			}
		}
	}
	return indexs
}

func main() {
	fmt.Println("vim-go")
	target := 15
	nums := []int{2, 4, 7, 8, 10, 11, 13, 17, 15}
	ret := twoSum(nums, target)
	fmt.Println("Result:", ret)
}
