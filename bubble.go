package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	arrSize := 1000
	numCoresList := []int{1, 2, 4, 8}

	fmt.Printf("Array size: %d\n", arrSize)

	slice := createSlice(arrSize)
	fmt.Println("\nUnsorted\n", slice)

	fmt.Println("\nConcurrent Bubble Sort Execution Times:")
	for _, numCores := range numCoresList {
		sortedArr, execTime := measureBubbleSortExecutionTime(slice, numCores)
		fmt.Printf("Num Cores: %d, Execution Time: %v\n", numCores, execTime)
		fmt.Println("Sorted:", sortedArr)
	}
}

func createSlice(size int) []int {
	slice := make([]int, size, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func measureBubbleSortExecutionTime(arr []int, numCores int) ([]int, time.Duration) {
	runtime.GOMAXPROCS(numCores)
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)

	start := time.Now()
	concurrentBubbleSort(arrCopy)
	execTime := time.Since(start)

	return arrCopy, execTime
}

func concurrentBubbleSort(a []int) {
	size := len(a)
	var wg sync.WaitGroup
	for i := 0; i < size-1; i++ {
		wg.Add(1)
		go func(pass int) {
			defer wg.Done()
			for j := 0; j < size-pass-1; j++ {
				if a[j] > a[j+1] {
					a[j], a[j+1] = a[j+1], a[j]
				}
			}
		}(i)
	}
	wg.Wait()
}
