package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func randInt32Arr(size int, max int32) []int32 {
	if size <= 0 {
		return nil
	}
	rand.NewSource(time.Now().UnixNano())
	m := make(map[int32]bool)
	for len(m) < size {
		n := rand.Int31n(max)
		if n == 0 {
			continue
		}
		// unique number
		if !m[n] {
			m[n] = true
		}
	}

	arr := make([]int32, 0)
	for i := range m {
		arr = append(arr, i)
	}

	// sort
	sort.Slice(arr, func(x, y int) bool {
		return arr[x] < arr[y]
	})
	return arr
}

func main() {
	arr := randInt32Arr(int(math.Max(float64(5), float64(rand.Intn(15)))), 20)
	fmt.Println(arr)

	arr2 := GetRandomArray(int(math.Max(float64(5), float64(rand.Intn(15)))), 20)
	fmt.Println(arr2)
}

func GetRandomArray(size int, max int) []int {
	// Create a slice to store the random numbers.
	numbers := make([]int, size)

	// Generate random numbers and store them in the slice.
	for i := 0; i < size; i++ {
		numbers[i] = rand.Intn(max)
	}

	// Sort the numbers in ascending order.
	sort.Ints(numbers)
	//sort.Slice(numbers, func(x, y int) bool {
	//	return numbers[x] < numbers[y]
	//})

	// Remove any duplicate numbers from the slice.
	for i := 0; i < size-1; i++ {
		if numbers[i] == numbers[i+1] {
			numbers = append(numbers[:i], numbers[i+2:]...)
			i--
		}
	}

	// Return the slice of random numbers.
	return numbers
}
