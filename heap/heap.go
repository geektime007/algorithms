package main

import "fmt"

type Heap struct {
	// 数组，从下标1开始存储数据
	array []int
	// 堆可以存储的最大数据个数
	Cap int
	// 堆中已经存储的数据个数
	count int
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		array: make([]int, capacity+1),
		Cap:   capacity,
		count: 0,
	}
}

func (h *Heap) String() string {
	return fmt.Sprintf("%v\n", h.array)
}

// 大顶堆方式 插入数据
func (h *Heap) Insert(val int) {
	if h.count >= h.Cap {
		return
	}
	h.count++
	h.array[h.count] = val
	i := h.count
	// 自下往上
	for i/2 > 0 && h.array[i] > h.array[i/2] {
		h.array[i], h.array[i/2] = h.array[i/2], h.array[i]
		i = i / 2
	}
}

// 小顶堆方式 插入数据
func (h *Heap) Insert1(val int) {
	if h.count >= h.Cap {
		return
	}
	h.count++
	h.array[h.count] = val
	i := h.count
	// 自下往上
	for i/2 > 0 && h.array[i] < h.array[i/2] {
		h.array[i], h.array[i/2] = h.array[i/2], h.array[i]
		i = i / 2
	}
}

// 删除堆顶最大元素
func (h *Heap) RemoveMax() int {
	if h.count == 0 {
		return -1
	}
	tmp := h.array[h.count]

	h.array[1] = h.array[h.count]
	h.array[h.count] = 0 // 抹除最后一位下标的元素值
	h.count--
	h.heapify1(1) // 大顶堆堆化
	return tmp
}

// 删除堆顶最小元素
func (h *Heap) RemoveMin() int {
	if h.count == 0 {
		return -1
	}
	tmp := h.array[h.count]
	h.array[h.count] = 0 // 抹除最后一位下标的元素值
	h.count--
	// 小顶堆 堆化
	h.heapify3(h.count, 1)
	return tmp
}

// 建堆
func (h *Heap) BuildBigHeap() {
	for i := h.count / 2; i >= 1; i-- {
		h.heapify1(i) // 大顶堆 堆化
	}
}

// 建堆
func (h *Heap) BuildLittletHeap() {
	for i := h.count / 2; i >= 1; i-- {
		h.heapify3(h.count, i) // 小顶堆 堆化
	}
}

// 从下往上堆化  大顶堆
func (h *Heap) heapify1(pos int) {
	for {
		maxPos := pos
		// 左子结点
		if pos*2 <= h.count && h.array[pos] < h.array[pos*2] {
			maxPos = pos * 2
		}
		// 右子结点
		if pos*2+1 <= h.count && h.array[maxPos] < h.array[pos*2+1] {
			maxPos = pos*2 + 1
		}
		if maxPos == pos {
			break
		}
		h.array[pos], h.array[maxPos] = h.array[maxPos], h.array[pos]
		pos = maxPos
	}
}

// 从下往上堆化  大顶堆
func (h *Heap) heapify2(n int, pos int) {
	for {
		maxPos := pos
		if pos*2 <= n && h.array[pos] < h.array[pos*2] {
			maxPos = pos * 2
		}
		if pos*2+1 <= n && h.array[maxPos] < h.array[pos*2+1] {
			maxPos = pos*2 + 1
		}
		if maxPos == pos {
			break
		}
		h.array[pos], h.array[maxPos] = h.array[maxPos], h.array[pos]
		pos = maxPos
	}
}

// 从下往上堆化  小顶堆
func (h *Heap) heapify3(n int, pos int) {
	for {
		minPos := pos
		// 左子结点下标小于数组数据个数
		if pos*2 <= n && h.array[pos] > h.array[pos*2] {
			minPos = pos * 2
		}
		// 右子节点下标小于数组数据个数
		if pos*2+1 <= n && h.array[minPos] > h.array[pos*2+1] {
			minPos = pos*2 + 1
		}
		if minPos == pos {
			break
		}
		h.array[pos], h.array[minPos] = h.array[minPos], h.array[pos]
		pos = minPos
	}
}

// 堆排序 大顶堆
func (h *Heap) Sort() {
	h.BuildBigHeap()
	k := h.count
	for k > 1 {
		h.array[1], h.array[k] = h.array[k], h.array[1]
		k--
		h.heapify2(k, 1)
	}
}

// 堆排序 小顶堆
func (h *Heap) Sort1() {
	h.BuildLittletHeap()
	k := h.count
	for k > 1 {
		h.array[1], h.array[k] = h.array[k], h.array[1]
		k--
		h.heapify3(k, 1)
	}
}

// TopK  大顶堆方式
func (h *Heap) BigTopK(n int) []int {
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, h.RemoveMax())
	}
	return res
}

// TopK 小顶堆方式
func (h *Heap) LittltTopK(n int) []int {
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, h.RemoveMin())
	}
	return res
}

func main() {
	fmt.Println("vim-go")
	heap := NewHeap(10)
	for i := 10; i <= 20; i++ {
		heap.Insert(i)
		fmt.Println("打印：", heap)
	}
	heap.Sort()
	fmt.Println("大顶堆排序后：", heap)
	heap.BuildBigHeap()
	fmt.Println("大顶堆堆化：", heap)
	heap.Sort1()
	fmt.Println("小顶堆排序后：", heap)
	heap.BuildLittletHeap()
	fmt.Println("小顶堆堆化：", heap)
	for i := 0; i < 3; i++ {
		heap.RemoveMax()
		fmt.Println("删除堆顶元素：", heap)
	}
	fmt.Println("取出最大的三个数：", heap.BigTopK(3))

	// ================================================
	fmt.Println("==============================")

	h1 := NewHeap(5)
	for i := 1; i <= 5; i++ {
		h1.Insert1(i)
		fmt.Println("插入后打印：", h1)
	}
	h1.Sort1()
	fmt.Println("小顶堆排序后：", h1)

	fmt.Println("取出最小的三个数：", h1.LittltTopK(3))
}
