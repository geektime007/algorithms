package main

import "fmt"

type Heap struct {
	array []int
	count int
	Cap   int
}

func NewHeap(Cap int) *Heap {
	return &Heap{
		array: make([]int, Cap+1),
		count: 0,
		Cap:   Cap,
	}
}

func (h *Heap) String() string {
	return fmt.Sprintf("%v\n", h.array)
}

// 大顶堆插入一个元素
func (h *Heap) Insert(val int) bool {
	if h.count >= h.Cap {
		return false
	}
	h.count++
	h.array[h.count] = val
	i := h.count
	for i/2 > 0 && h.array[i] > h.array[i/2] {
		h.array[i], h.array[i/2] = h.array[i/2], h.array[i]
		i = i / 2
	}
	return true
}

// 小顶堆插入 一个元素
func (h *Heap) Insert1(val int) bool {
	if h.count >= h.Cap {
		return false
	}
	h.count++
	h.array[h.count] = val
	i := h.count
	for i/2 > 0 && h.array[i] < h.array[i/2] {
		h.array[i], h.array[i/2] = h.array[i/2], h.array[i]
		i = i / 2
	}
	return true
}

// 大顶堆 TopK
func (h *Heap) BigTopK(n int) []int {
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, h.BigHeapRemoveMax())
	}
	return res
}

// 大顶堆 删除堆顶元素
func (h *Heap) BigHeapRemoveMax() int {
	if h.count <= 0 {
		return -1
	}
	tmp := h.array[1]
	h.array[1] = h.array[h.count]
	h.array[h.count] = 0
	h.count--
	h.heapify1(1)
	return tmp
}

// 大顶堆 堆化
func (h *Heap) heapify1(pos int) {
	for {
		maxPos := pos
		if 2*pos <= h.count && h.array[pos] < h.array[2*pos] {
			maxPos = 2 * pos
		}
		if 2*pos+1 <= h.count && h.array[maxPos] < h.array[2*pos+1] {
			maxPos = 2*pos + 1
		}
		if maxPos == pos {
			break
		}
		h.array[pos], h.array[maxPos] = h.array[maxPos], h.array[pos]
		pos = maxPos
	}
}

// 大顶堆 堆化
func (h *Heap) heapify2(n int, pos int) {
	for {
		maxPos := pos
		if 2*pos <= n && h.array[pos] < h.array[2*pos] {
			maxPos = 2 * pos
		}
		if 2*pos+1 <= n && h.array[maxPos] < h.array[2*pos+1] {
			maxPos = 2*pos + 1
		}
		if maxPos == pos {
			break
		}
		h.array[pos], h.array[maxPos] = h.array[maxPos], h.array[pos]
		pos = maxPos
	}
}

// 大顶堆排序
func (h *Heap) BigHeapSort() {
	h.BuildBigHeap()
	k := h.count
	for k > 1 {
		h.array[1], h.array[k] = h.array[k], h.array[1]
		k--
		h.heapify2(k, 1)
	}
}

// 大顶堆 建堆
func (h *Heap) BuildBigHeap() {
	for i := h.count / 2; i >= 1; i-- {
		h.heapify1(i)
	}
}

// 小顶堆排序
func (h *Heap) LittleHeapSort() {
	h.BuildLittletHeap()
	k := h.count
	for k > 1 {
		h.array[1], h.array[k] = h.array[k], h.array[1]
		k--
		h.heapify3(k, 1)
	}
}

// 小顶堆 建堆
func (h *Heap) BuildLittletHeap() {
	for i := h.count / 2; i > 0; i-- {
		h.heapify3(h.count, i)
	}
}

// 小顶堆 堆化
func (h *Heap) heapify3(n int, pos int) {
	for {
		minPos := pos
		if 2*pos <= n && h.array[pos] > h.array[2*pos] {
			minPos = 2 * pos
		}
		if 2*pos+1 <= n && h.array[minPos] > h.array[2*pos+1] {
			minPos = 2*pos + 1
		}
		if minPos == pos {
			break
		}
		h.array[minPos], h.array[pos] = h.array[pos], h.array[minPos]
		pos = minPos
	}
}

// 小顶堆 TopK
func (h *Heap) LittltTopK(n int) []int {
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, h.LittleHeapRemoveMin())
	}
	return res
}

// 小顶堆删除堆顶元素
func (h *Heap) LittleHeapRemoveMin() int {
	if h.count <= 0 {
		return -1
	}
	tmp := h.array[h.count]
	h.array[1] = h.array[h.count]
	h.array[h.count] = 0
	h.count--
	h.heapify3(h.count, 1)
	return tmp
}

func main() {
	fmt.Println("vim-go")
	h := NewHeap(5)
	for i := 1; i <= 5; i++ {
		h.Insert(i)
		fmt.Println("插入后打印：", h)
	}
	h.BigHeapSort()
	fmt.Println("大顶堆排序后：", h)
	h.BuildBigHeap()
	fmt.Println("大顶堆堆化后：", h)
	fmt.Println("大顶堆取TopK: ", h.BigTopK(3))
	h.LittleHeapSort()
	fmt.Println("小顶堆排序后：", h)
	h.Insert1(3)
	h.Insert1(4)
	h.Insert1(5)
	h.Insert1(6)
	h.LittleHeapSort()
	fmt.Println("小顶堆排序后：", h)
	h.BuildLittletHeap()
	fmt.Println("小顶堆堆化后：", h)
	h.LittleHeapSort()
	fmt.Println("小顶堆排序后：", h)

	fmt.Println("小顶堆取TopK: ", h.LittltTopK(3))
}
