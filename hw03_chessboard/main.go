package main

import "fmt"

func main() {
	var size int
	fmt.Print("Введите размер шахматной доски:")
	_, err := fmt.Scanln(&size)
	if err != nil || size < 1 || size == 0 {
		fmt.Println("Введите число больше нуля!")
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println() // Переход на новою строку иначе все будет в 1 линию
	}
}
