package main

import "fmt"

// 快速排序
// 思路分析：
// 选择一格基准元素,
// 通常选择第一个元素或者最后一个元素。通过一趟扫描，
// 将待排序列分成两部分，一部分比基准元素小，一部分
// 大于基准元素。此时基准元素在其排好序后的正确位置，
// 然后再用同样的方法递归地排序划分的两部分。

func quickSort(nums []int) []int {
	// 先判断是否需要继续进行
	length := len(nums)
	if length <= 1 {
		return nums
	}

	// 选择第一个元素作为基准
	baseNum := nums[0]
	// 遍历除了标尺外的所有元素，按照大小关系放入左右两个切片内
	// 初始化左右两个切片
	leftNums := []int{}  // 小雨基准的
	rightNums := []int{} // 大于基准的
	for i := 1; i < length; i++ {
		if baseNum > nums[i] {
			// 放入左边切片
			leftNums = append(leftNums, nums[i])
		} else {
			// 放入右边切片
			rightNums = append(rightNums, nums[i])
		}
	}

	// 再分别对左边和右边切片进行相同的处理方式递归调用这个函数
	leftNums = quickSort(leftNums)
	rightNums = quickSort(rightNums)
	leftNums = append(leftNums, baseNum)
	return append(leftNums, rightNums...)
}

func main() {
	fmt.Println("vim-go")
	a := []int{2, 4, 1, 8, 3, 5, 10}
	fmt.Println(a)
	b := quickSort(a)
	fmt.Println(b)
}
