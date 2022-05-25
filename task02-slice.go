package homework

func reverse(input []int64) (result []int64) {
	for i := range input {
		result = append(result, input[len(input)-i-1])
	}
	return
}
