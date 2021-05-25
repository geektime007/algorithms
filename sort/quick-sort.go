package main

import "fmt"

// 快速排序
// 思路分析：
// 选择一格基准元素,
// 通常选择第一个元素或者最后一个元素。通过一趟扫描，
// 将待排序列分成两部分，一部分比基准元素小，一部分
// 大于基准元素。此时基准元素在其排好序后的正确位置，
// 然后再用同样的方法递归地排序划分的两部分。

// 使用了额外的内存空间
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

// 不实用额外的内存空间  原地排序
func QuickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	_QuickSort(nums, 0, len(nums)-1)
	return nums
}

func _QuickSort(nums []int, left, right int) {
	if left > right {
		return
	}
	// 位置划分
	pivot := partition(nums, left, right)
	// 左边部分排序
	_QuickSort(nums, left, pivot-1)
	// 右边部分排序
	_QuickSort(nums, pivot+1, right)
}

func partition(nums []int, left, right int) int {
	// 导致 low 位置值为空
	pivot := nums[left]
	for left < right {
		// right 指针值 >= pivot right 左移
		for left < right && pivot <= nums[right] {
			right--
		}
		// 填补 low 位置空值
		// right 指针值 < pivot right值 移动到 low 位置
		// right 位置值空
		nums[left] = nums[right]

		for left < right && pivot >= nums[left] {
			left++
		}

		// 填补 right 位置空值
		// low 指针值 > pivot low 值 移动到 right 位置
		// low 位置值空
		nums[right] = nums[left]
	}

	// pivot 填补 low 位置的空值
	nums[left] = pivot
	return left
}

func main() {
	fmt.Println("vim-go")
	a := []int{2, 4, 1, 8, 3, 5, 10}
	fmt.Println(a)
	//b := quickSort(a)
	b := QuickSort(a)
	fmt.Println(b)
}
