package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	arrSize := 100000
	numCoresList := []int{4}

	fmt.Printf("Array size: %d\n", arrSize)

	slice := createSlice(arrSize)
	fmt.Println("\nUnsorted\n", slice)

	fmt.Println("\nConcurrent Quick Sort Execution Times:")
	for _, numCores := range numCoresList {
		sortedArr, execTime := measureQuickSortExecutionTime(slice, numCores)
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

func measureQuickSortExecutionTime(arr []int, numCores int) ([]int, time.Duration) {
	runtime.GOMAXPROCS(numCores)
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)

	start := time.Now()
	concurrentQuickSort(arrCopy)
	execTime := time.Since(start)

	return arrCopy, execTime
}

func concurrentQuickSort(a []int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)
	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		concurrentQuickSort(a[:left])
	}()
	concurrentQuickSort(a[left+1:])

	wg.Wait()
}
