package slices

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

func ToMap[T any, R comparable](items []T, keyFunc func(item T) R) map[R]any {
	m := make(map[R]any, 0)
	for _, v := range items {
		m[keyFunc(v)] = v
	}
	return m
}

// Filter return all elements match predicate
func Fliter[T any](items []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range items {
		v := v
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Distinct delete repeated element, keep origin order
func Distinct[T comparable](items []T) []T {
	result := make([]T, 0)
	m := make(map[T]struct{}, 0)
	for _, item := range items {
		if _, ok := m[item]; !ok {
			m[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// Distinct delete repeated element judge by distinctFunc, keep origin order
func DistinctI[T any, R comparable](items []T, distinctFunc func(T) R) []T {
	result := make([]T, 0)
	m := make(map[R]struct{}, 0)
	for _, item := range items {
		key := distinctFunc(item)
		if _, ok := m[key]; !ok {
			m[key] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// Search return the first index of target
// If not exist, return -1
func Search[T comparable](items []T, target T) int {
	for i, v := range items {
		if v == target {
			return i
		}
	}
	return -1
}
