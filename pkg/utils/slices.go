package utils

func Flatten[T any](input [][]T) []T {
	result := make([]T, 0)

	for _, slice := range input {
		result = append(result, slice...)
	}

	return result
}
