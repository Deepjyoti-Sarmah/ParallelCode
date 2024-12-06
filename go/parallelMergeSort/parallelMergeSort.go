package parallelmergesort

import (
	"sort"
	"sync"
)

type ParallelMergeSort struct {
	nums []int
}

func NewParallelMergeSort(nums []int) *ParallelMergeSort {
	return &ParallelMergeSort{nums: nums}
}

func (p *ParallelMergeSort) Sort() {
	if len(p.nums) == 0 {
		return
	}

	p.recursiveSort(0, len(p.nums)-1)
}

func (p *ParallelMergeSort) recursiveSort(left, right int) {
	const threshold = 5000

	if right-left < threshold {
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
		p.recursiveSort(left, mid)
	}()

	go func() {
		defer wg.Done()
		p.recursiveSort(mid+1, right)
	}()

	wg.Wait()

	result := make([]int, 0, right-left+1)
	i, j := left, mid+1

	for i <= mid && j <= right {
		if p.nums[i] <= p.nums[j] {
			result = append(result, p.nums[i])
			i++
		} else {
			result = append(result, p.nums[j])
			j++
		}
	}

	for i <= mid {
		result = append(result, p.nums[i])
		i++
	}

	for j <= right {
		result = append(result, p.nums[j])
		j++
	}

	copy(p.nums[left:right+1], result)
}
