package main

import "fmt"

// 选择排序
// 思路分析：在要排序的切片中，选出最小的一个元素与
// 第一个位置的元素交换。然后在剩下的元素当中再找最
// 小的与第二个位置的元素交换，如此循环到倒数第二个
// 元素和最后一个元素比较为止。

// 重复（元素个数-1）次
//     把第一个没有排序过的元素设置为最小值
//     遍历每个没有排序过的元素
//         如果元素 < 现在的最小值
//             将此元素设置成为新的最小值
//     将最小值和第一个没有排序过的位置交换

func selectSort(nums []int) []int {
	//双重循环完成，外层控制轮数，内层控制比较次数
	length := len(nums)
	for i := 0; i < length-1; i++ {
		//先假设最小的值的位置
		k := i
		for j := i + 1; j < length; j++ {
			//nums[k] 是当前已知的最小值
			if nums[k] > nums[j] {
				// 比较，发现更小的,记录下最小值的位置；
				// 并且在下次比较时采用已知的最小值进行比较。
				k = j
			}
		}
		// 已经确定了当前的最小值的位置，保存到 k 中。
		// 如果发现最小值的位置与当前假设的位置 i 不同，
		// 则位置互换即可。
		if k != i {
			nums[k], nums[i] = nums[i], nums[k]
		}
	}
	return nums
}

func selectSort1(nums []int) []int {
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

func main() {
	fmt.Println("vim-go")
	a := []int{2, 4, 1, 8, 3, 5, 10}
	fmt.Println(a)
	b := selectSort(a)
	fmt.Println(b)
}
