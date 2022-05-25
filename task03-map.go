package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))

	for i := range input {
		keys = append(keys, i)
	}

	sort.Ints(keys)

	for _, i := range keys {
		result = append(result, input[i])
	}

	return
}
