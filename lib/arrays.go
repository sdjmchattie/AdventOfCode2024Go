package lib

import "strconv"

func MapArray[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
			result[i] = fn(t)
	}
	return result
}

func MapInts(strs []string) []int {
	result := make([]int, len(strs))

	for i, t := range strs {
		intVal, err := strconv.Atoi(t)

		if err != nil {
			panic(err)
		}

		result[i] = intVal
	}

	return result
}
