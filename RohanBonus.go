package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	numbers, err := readNumbers("in.csv")
	if err != nil {
		log.Fatalf("Error reading numbers: %v", err)
	}

	start := time.Now()

	if len(numbers) < 1000 {
		radixSort(numbers)
	} else {
		quickSort(numbers)
	}

	elapsed := time.Since(start)

	err = writeNumbers("out(20086377).csv", numbers)
	if err != nil {
		log.Fatalf("Error writing numbers: %v", err)
	}

	fmt.Printf("Sorted %d numbers in %s.\n", len(numbers), elapsed)

	if isSorted(numbers) {
		fmt.Println("The numbers are sorted correctly.")
	} else {
		fmt.Println("The numbers are not sorted correctly.")
	}
}

func getMax(arr []int) int {
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

func radixSort(a []int) {
	max := getMax(a)

	for exp := 1; max/exp > 0; exp *= 10 {
		count := make([]int, 10)
		output := make([]int, len(a))

		for _, num := range a {
			count[(num/exp)%10]++
		}

		for i := 1; i < 10; i++ {
			count[i] += count[i-1]
		}

		for i := len(a) - 1; i >= 0; i-- {
			digit := (a[i] / exp) % 10
			output[count[digit]-1] = a[i]
			count[digit]--
		}

		for i := range a {
			a[i] = output[i]
		}
	}
}

func quickSort(a []int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1

	pivot := a[right]
	for i := range a {
		if a[i] < pivot {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])
}

func readNumbers(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	numbers := []int{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		for _, value := range record {
			number, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, number)
		}
	}

	return numbers, nil
}

func writeNumbers(filename string, numbers []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, number := range numbers {
		err := writer.Write([]string{strconv.Itoa(number)})
		if err != nil {
			return err
		}
	}

	return nil
}

func isSorted(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] > numbers[i] {
			return false
		}
	}
	return true
}
