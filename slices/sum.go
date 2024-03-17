package main

// Arrays have a fixed size, slices do not.
// To initialize a slice, use the make function.
// make([]int, 5) creates a slice with a length of 5.
// make([]type, size)

// To initialize an array, use the [size] syntax.
// [5]int{1, 2, 3, 4, 5} creates an array with a length of 5.
// [N]type{value1, value2, ..., valueN}
func Sum(numbers []int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))

	}
	return sums
}
