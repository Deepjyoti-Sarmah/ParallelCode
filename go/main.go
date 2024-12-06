package main

import (
	"fmt"
	"math/rand"
	"mergesort/mergeSort"
	parallelmergesort "mergesort/parallelMergeSort"
	"time"
)

func main() {
	const size = 10_000_000

	rand.Intn(int(time.Now().UnixNano()))

	nums := make([]int, size)
	nums1 := make([]int, size)

	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(1000)
		nums1[i] = nums[i]
	}

	//Regular Merge Sort
	mergeSort := mergesort.NewMergeSort(nums)
	start := time.Now()
	mergeSort.Sort()
	mergeSortDuration := time.Since(start)
	fmt.Printf("MergeSort time taken: %f seconds\n", mergeSortDuration.Seconds())

	//Parallel Merge Sort
	parallelMergeSort := parallelmergesort.NewParallelMergeSort(nums1)
	start = time.Now()
	parallelMergeSort.Sort()
	parallelMergeSortDuration := time.Since(start)
	fmt.Printf("parallelMergeSort time taken: %f seconds\n", parallelMergeSortDuration.Seconds())
}
