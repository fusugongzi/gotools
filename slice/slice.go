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

func Index[T rune | string](slice []T, item T) int {
	for idx, v := range slice {
		if v == item {
			return idx
		}
	}
	return -1
}
