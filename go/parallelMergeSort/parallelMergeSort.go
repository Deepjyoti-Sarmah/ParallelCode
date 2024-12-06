package parallelmergesort

import (
	"runtime"
	"sort"
	"sync"
)

type ParallelMergeSort struct {
	nums     []int
	maxDepth int
}

func NewParallelMergeSort(nums []int) *ParallelMergeSort {
	return &ParallelMergeSort{
		nums:     nums,
		maxDepth: calculateMaxDepth(),
	}
}

func calculateMaxDepth() int {
	numCPU := runtime.NumCPU()
	return int(1 + numCPU)
}

func (p *ParallelMergeSort) Sort() {
	if len(p.nums) == 0 {
		return
	}

	p.recursiveSort(0, len(p.nums)-1, 0)
}

func (p *ParallelMergeSort) recursiveSort(left, right, depth int) {

	const threshold = 10000
	if right-left < threshold || depth >= p.maxDepth {
		sort.Ints(p.nums[left : right+1])
		return
	}

	if left >= right {
		return
	}

	mid := left + (right-left)/2

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		p.recursiveSort(left, mid, depth+1)
	}()

	go func() {
		defer wg.Done()
		p.recursiveSort(mid+1, right, depth+1)
	}()

	wg.Wait()

	p.merge(left, mid, right)
}

func (p *ParallelMergeSort) merge(left, mid, right int) {

	if p.nums[mid] <= p.nums[mid+1] {
		return
	}

	result := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	for i <= mid && j <= right {
		if p.nums[i] <= p.nums[j] {
			result[k] = p.nums[i]
			i++
		} else {
			result[k] = p.nums[j]
			j++
		}
		k++
	}

	for i <= mid {
		result[k] = p.nums[i]
		i++
		k++
	}

	for j <= right {
		result[k] = p.nums[j]
		j++
		k++
	}

	copy(p.nums[left:right+1], result)
}
