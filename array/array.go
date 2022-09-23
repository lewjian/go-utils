package array

import (
	"github.com/lewjian/utils/collection"
)

// InArray check if data contains element item, index=-1 when not exists
func InArray[T comparable](item T, data []T) (exist bool, index int) {
	index = -1
	for i := 0; i < len(data); i++ {
		if data[i] == item {
			exist = true
			index = i
			return
		}
	}
	return
}

// Diff returns elements in src, but not in others...
func Diff[T comparable](src []T, others ...[]T) []T {
	if len(others) == 0 {
		return src
	}
	s := collection.NewSet[T]()
	for i := 0; i < len(others); i++ {
		s.Add(others[i]...)
	}
	diffs := make([]T, 0, len(src))
	for i := 0; i < len(src); i++ {
		if !s.Has(src[i]) {
			diffs = append(diffs, src[i])
		}
	}
	return diffs
}

// Intersect return a slice contains values all arrs has
func Intersect[T comparable](arrs ...[]T) []T {
	if len(arrs) == 0 {
		return nil
	}
	if len(arrs) == 0 {
		return arrs[0]
	}
	m := make(map[T]int, len(arrs[0]))
	for i := 0; i < len(arrs); i++ {
		if len(arrs[i]) == 0 {
			return nil
		}
		for j := 0; j < len(arrs[i]); j++ {
			m[arrs[i][j]]++
		}
	}
	results := make([]T, 0, len(m))
	for val, count := range m {
		if count == len(arrs) {
			results = append(results, val)
		}
	}
	return results
}

// Merge all slice into one
func Merge[T comparable](base []T, others ...[]T) []T {
	if len(others) == 0 {
		return base
	}
	// calculate element count
	count := len(base)
	for _, other := range others {
		count += len(other)
	}
	results := make([]T, 0, count)
	results = append(results, base...)
	for _, other := range others {
		results = append(results, other...)
	}
	return results
}

// Filter a slice use function f, element dropped when f get false
func Filter[T comparable](data []T, f func(i int) bool) []T {
	if len(data) == 0 {
		return data
	}
	results := make([]T, 0, len(data))
	for i, item := range data {
		if f(i) {
			results = append(results, item)
		}
	}
	return results
}

// Range same as for range, break when f got true
func Range[T comparable](data []T, f func(i int) (breakHere bool)) {
	for i, _ := range data {
		if f(i) {
			break
		}
	}
}

// Reverse slice on the input data slice
func Reverse[T comparable](data []T) {
	if len(data) == 0 {
		return
	}
	for i, j := 0, len(data)-1; i < len(data)/2; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// ReverseNew returns a new reverse slice
func ReverseNew[T comparable](data []T) []T {
	if len(data) == 0 {
		return data
	}
	newData := make([]T, len(data))
	for i, _ := range data {
		newData[len(data)-i-1] = data[i]
	}
	return newData
}

// Remove index element of data
func Remove[T comparable](data []T, index int) []T {
	if index < 0 {
		return data
	}
	if index >= len(data) {
		return data
	}
	results := make([]T, 0, len(data)-1)
	results = append(results, data[:index]...)
	if index < len(data)-1 {
		results = append(results, data[index+1:]...)
	}
	return results
}

// RemoveElement remove element of data, only remove first count elements
// if count < 0, remove all elements
// 删除切片中出现的前count次元素，如果count小于0，则删除所有切片中的element
func RemoveElement[T comparable](data []T, item T, count int) []T {
	if len(data) == 0 {
		return data
	}
	if count == 0 {
		return data
	}
	deletedCount := 0
	results := make([]T, 0, len(data))
	for i := 0; i < len(data); i++ {
		if data[i] != item {
			// 不需要删除
			results = append(results, data[i])
			continue
		}
		if count < 0 {
			// 删除所有
			continue
		}
		deletedCount++
		if deletedCount >= count && i+1 < len(data) {
			results = append(results, data[i+1:]...)
			break
		}
	}
	return results
}

// RemoveAll remove all element item from data slice
func RemoveAll[T comparable](data []T, item T) []T {
	return RemoveElement(data, item, -1)
}
