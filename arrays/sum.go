package main

/*
Sum of numbers in an arrays
*/
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
