package maps

func MapSquare(numbers []int, f func(int) int) []int {
	result := make([]int, len(numbers))

	for index, value := range numbers {
		result[index] = f(value)
	}

	return result
}
