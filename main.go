package main

import "fmt"

func filter(numbers []int, f func(int) bool) []int {

	var result []int
	for _, value := range numbers {

		if f(value) {
			result = append(result, value)
		}
	}

	return result
}

func mapSquare(numbers []int, f func(int) int) []int {
	result := make([]int, len(numbers))

	for index, value := range numbers {
		result[index] = f(value)
	}

	return result
}

func main() {
	numbers := []int{10, 3, 4, 5, 6, 7}

	squares := mapSquare(numbers, func(x int) int { return x * x })
	fmt.Println(squares)

}
