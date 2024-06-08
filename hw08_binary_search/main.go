package main

import (
	"fmt"
)

func main() {
	var low, high, target int
	fmt.Print("Введите диапазон поиска через пробел: ")
	_, err := fmt.Scan(&low, &high)
	if err != nil {
		fmt.Println(err)
		return
	}
	if low > high {
		fmt.Println("Неверный диапазон. Первое число не может быть больше второго!")
		return
	}
	fmt.Print("Введите загаданное число: ")
	_, err = fmt.Scan(&target)
	if err != nil {
		fmt.Println(err)
		return
	}
	arr := make([]int, high-low+1)
	for i := low; i <= high; i++ {
		arr[i-low] = i
	}
	index := BinarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Индекс числа %d: %d\n", target, index)
	} else {
		fmt.Printf("Число %d не найдено в диапазоне.\n", target)
	}
}

func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
