package utils

import (
	"math"
)

// SeparationSlice Разделение слайса на части
func SeparationSlice(originSlice []string, size int) [][]string {

	if size <= 0 {
		return nil
	}
	// Определяем capacity нового слайса
	dif := float64(len(originSlice)) / float64(size)
	caps := int(math.Ceil(dif))
	newSlice := make([][]string, 0, caps)

	var low int
	var high int
	for i := 0; i < caps; i++ {
		low = i * size
		high = i*size + size
		if high > len(originSlice) {
			high = len(originSlice)
		}
		newSlice = append(newSlice, originSlice[low:high])
	}
	return newSlice
}

// ReverseKeyView - Обратный ключ
// Конвертирует отображения слайса (“ключ-значение“) в map (“значение-ключ“)
func ReverseKeyView(origin map[int]string) map[string]int {
	revert := make(map[string]int, len(origin))
	for key, value := range origin {
		revert[value] = key
	}

	return revert
}

// FilterSlice Фильтрация по захардкоженному списку
func FilterSlice(origin []string) []string {
	filterList := []string{
		"rebus",
		"conus",
		"globus",
	}
	filtered := make([]string, 0, len(origin))

	filter := func(val string, array []string) bool {
		for i := range array {
			if array[i] == val {
				return false
			}
		}
		return true
	}

	for _, value := range origin {
		if filter(value, filterList) {
			filtered = append(filtered, value)
		}
	}

	return filtered
}
