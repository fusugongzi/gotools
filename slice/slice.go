package slice

import "strings"

func Contains[T int | string](items []T, item T) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func ContainsIgnoreCase(items []string, item string) bool {
	itemLower := strings.ToLower(item)
	for _, v := range items {
		vLower := strings.ToLower(v)
		if vLower == itemLower {
			return true
		}
	}
	return false
}

func Index[T rune | string](items []T, item T) int {
	for idx, v := range items {
		if v == item {
			return idx
		}
	}
	return -1
}

type any = interface{}

func ToMap[T any, R int | int32 | int64 | string](items []T, keyFunc func(item T) R) map[R]any {
	m := make(map[R]any, 0)
	for _, v := range items {
		m[keyFunc(v)] = v
	}
	return m
}
