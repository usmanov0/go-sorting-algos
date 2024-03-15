package main

import (
	"fmt"
	"sort"
)

func bucketSort(arr []float64) []float64 {
	if len(arr) <= 1 {
		return arr
	}
	// Determine the number of buckets (assumed to be 10 in this example)
	numBuckets := 10
	buckets := make([][]float64, numBuckets)

	// Distribute elements into buckets
	for _, value := range arr {
		index := int(value * float64(numBuckets))
		buckets[index] = append(buckets[index], value)
	}

	// Sort each bucket using a simple sorting algorithm (e.g., insertion sort)
	for i := 0; i < numBuckets; i++ {
		sort.Float64s(buckets[i])
	}

	// Concatenate the sorted buckets to get the final sorted array
	sortedArr := make([]float64, 0, len(arr))
	for i := 0; i < numBuckets; i++ {
		sortedArr = append(sortedArr, buckets[i]...)
	}

	return sortedArr
}

func main() {
	arr := []float64{0.64, 0.34, 0.25, 0.12, 0.22, 0.11, 0.90}
	fmt.Println("Unsorted array:", arr)
	arr = bucketSort(arr)
	fmt.Println("Sorted array:", arr)
}
