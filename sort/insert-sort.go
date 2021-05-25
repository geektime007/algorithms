package main

import "fmt"

// 插入排序
// 在要排序的切片中，假设前面的元素已经是排好顺序的
// ，现在要把第n个元素插到前面的有序切片中，使得这n个
// 元素也是排好顺序的。如此反复循环，直到全部排好顺序。

// 将第一个元素标记为已排序
// 遍历每个没有排序过的元素
//     “提取”元素 X
//     i = 最后排序过元素的指数 到 0 的遍历
//       如果现在排序过的元素 > 提取的元素
//          将排序过的元素向右移动一格
//       否则：插入提取的元素

func insertSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		tmp := nums[i]
		for j := i - 1; j >= 0; j-- {
			if tmp < nums[j] {
				// 这一层 for 循环是将小的元素往前挪
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
