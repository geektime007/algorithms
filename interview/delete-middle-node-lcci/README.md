# 面试题 02.03 删除中间节点

### 难度: 简单

## 题目
------

https://leetcode-cn.com/problems/delete-middle-node-lcci/
------

## 原题
------
![](img/leetcode-0203.png)

## 解法一
------

```golang
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
```

## 解法二
------

```golang
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
```

