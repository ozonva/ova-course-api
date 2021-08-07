package main

import (
	"fmt"

	"ova-course-api/internal/utils"
)

func main() {
	fmt.Println("Welcome to ova-template-api")

	fmt.Println("Задание 1 - Разделение слайса на батчи")
	{
		origin := []string{"a", "b", "c", "d", "e"}
		converted := utils.SeparationSlice(origin, 2)
		fmt.Printf("Входной срез: %v\n", origin)
		fmt.Printf("Выходной срез: %v\n", converted)
	}

	fmt.Println("Задание 2 - Обратный ключ ")
	{
		origin := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
		revert := utils.ReverseKeyView(origin)
		fmt.Printf("Входной Map: %v\n", origin)
		fmt.Printf("Выходной Map: %v\n", revert)
	}

	fmt.Println("Задание 3 - Фильтрация по захардкоженному списку")
	{
		origin := []string{"map", "rebus", "car", "globus"}
		filtered := utils.FilterSlice(origin)
		fmt.Printf("Входной срез: %v\n", origin)
		fmt.Printf("Выходной срез: %v\n", filtered)
	}

}