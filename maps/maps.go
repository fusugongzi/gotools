package maps

import "gotools/sets"

func KeySet[T comparable, R any](m map[T]R) []T {
	result := make([]T, 0)
	for k, _ := range m {
		result = append(result, k)
	}
	return result
}

func ValueSet[T comparable, R any](m map[T]R, distinctFunc func(R) T) []R {
	set := sets.NewSet[T, R](distinctFunc)
	for _, v := range m {
		set.Add(v)
	}
	return set.Values()
}
