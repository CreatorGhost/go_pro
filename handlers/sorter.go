package handlers

import (
	"sort"
	"sync"
	"time"
)

// sortArray is a helper function that sorts a single array of integers.
func sortArray(arr []int) []int {
	sort.Ints(arr)
	return arr
}

// processSingle sorts each sub-array sequentially and measures the time taken.
func processSingle(arrays [][]int) ([][]int, int64) {
	start := time.Now()

	sortedArrays := make([][]int, len(arrays))
	for i, array := range arrays {
		sortedArrays[i] = sortArray(array)
	}

	elapsed := time.Since(start)
	return sortedArrays, elapsed.Nanoseconds()
}

// processConcurrent sorts each sub-array concurrently and measures the time taken.
func processConcurrent(arrays [][]int) ([][]int, int64) {
	start := time.Now()

	var wg sync.WaitGroup
	sortedArrays := make([][]int, len(arrays))
	wg.Add(len(arrays))

	for i, array := range arrays {
		go func(i int, arr []int) {
			defer wg.Done()
			sortedArrays[i] = sortArray(arr)
		}(i, array)
	}

	wg.Wait()

	elapsed := time.Since(start)
	return sortedArrays, elapsed.Nanoseconds()
}
