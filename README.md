# Sorting Algorithms in Golang
This repository contains Golang implementations of various sorting algorithms. Each algorithm is implemented in a separate file for clarity.

# Sorting Algorithms Included


Bubble Sort:

File: bubble_sort.go
Description: Simple sorting algorithm that repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order.


Insertion Sort:

File: insertion_sort.go
Description: Builds the sorted array one item at a time by repeatedly taking an element from the unsorted part and inserting it into its correct position in the sorted part.


Quick Sort:

File: quick_sort.go
Description: Efficient, divide-and-conquer sorting algorithm that works by selecting a 'pivot' element from the array and partitioning the other elements into two sub-arrays according to whether they are less than or greater than the pivot.


Heap Sort:

File: heap_sort.go
Description: Heapsort uses a binary heap data structure to build a partially sorted array and efficiently maintains the heap property during the sorting process.


Selection Sort:

File: selection_sort.go
Description: Simple sorting algorithm that divides the input array into a sorted and an unsorted region. It repeatedly selects the smallest (or largest) element from the unsorted region and swaps it with the first element of the unsorted region.


Merge Sort:

File: merge_sort.go
Description: Divide-and-conquer sorting algorithm that divides the unsorted list into n sub-lists, each containing one element, and then repeatedly merges sub-lists to produce new sorted sub-lists until there is only one sub-list remaining.


Radix Sort:

File: radix_sort.go
Description: Non-comparative sorting algorithm that sorts integers by processing individual digits. It distributes elements into buckets based on their individual digits, then concatenates the buckets to form the sorted array.


Bucket Sort:

File: bucket_sort.go
Description: Sorting algorithm that distributes elements into a number of buckets, each of which is then sorted individually. The sorted buckets are then concatenated to produce the final sorted array.