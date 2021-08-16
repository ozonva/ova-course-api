package utils

import (
	"errors"

	"github.com/ozonva/ova-course-api/internal/course"
)

// SeparationSlice Разделение слайса на части
func SeparationSlice(originSlice []string, size int) [][]string {

	if size <= 0 {
		return [][]string{}
	}
	// Определяем capacity нового слайса
	caps := len(originSlice) / size
	if len(originSlice)%size > 0 {
		caps++
	}

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
func ReverseKeyView(origin map[int]string) (map[string]int, error) {
	revert := make(map[string]int, len(origin))
	for key, value := range origin {
		if _, ok := revert[value]; ok {
			err := errors.New("non-unique value")
			return map[string]int{}, err
		}
		revert[value] = key
	}

	return revert, nil
}

// FilterSlice Фильтрация по захардкоженному списку
func FilterSlice(origin []string) []string {
	filterList := []string{
		"rebus",
		"conus",
		"globus",
	}
	filtered := make([]string, 0, len(origin))

	filter := make(map[string]struct{}, len(origin))
	for _, value := range filterList {
		filter[value] = struct{}{}
	}

	for _, value := range origin {
		if _, ok := filter[value]; !ok {
			filtered = append(filtered, value)
		}
	}

	return filtered
}

func SplitToBulks(origin []course.Course, chunkSize int) [][]course.Course {

	if chunkSize <= 0 {
		return [][]course.Course{}
	}

	caps := len(origin) / chunkSize
	if len(origin)%chunkSize > 0 {
		caps++
	}

	newSlice := make([][]course.Course, 0, caps)

	for i := 0; i < caps; i += chunkSize {
		high := i + chunkSize
		if high > len(origin) {
			high = len(origin)
		}
		newSlice = append(newSlice, origin[i:high])
	}
	return newSlice

}

func SliceToMap(origin []course.Course) (map[uint64]course.Course, error) {

	result := make(map[uint64]course.Course)

	for _, val := range origin {
		if _, ok := result[val.Id()]; ok {
			return nil, errors.New("map has non-unique id")
		}
		result[val.Id()] = val
	}

	return result, nil

}
