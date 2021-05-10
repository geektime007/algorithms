package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indexs := []int{}
	if len(nums) == 0 {
		return nil
	}
	var indexsMap = map[int]int{}
	for i := 0; i < len(nums); i++ {
		anotherNum := target - nums[i]
		value, ok := indexsMap[anotherNum]
		if !ok {
			indexsMap[nums[i]] = i
		} else {
			indexs = append(indexs, i)
			indexs = append(indexs, value)
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
