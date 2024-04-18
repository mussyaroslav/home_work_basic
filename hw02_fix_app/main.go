package main

import (
	"fmt"
	"github.com/fixme_my_friend/hw02_fix_app/printer"
	"github.com/fixme_my_friend/hw02_fix_app/reader"
	"github.com/fixme_my_friend/hw02_fix_app/types"
)

func main() {
	var path string = "data.json"

	fmt.Printf("Введите путь к файлу данных: ")
	fmt.Scanln(&path)

	if len(path) == 0 {
		path = "data.json"
	}

	var err error
	var staff []types.Employee

	staff, err = reader.ReadJSON(path, -1)

	if err != nil {
		fmt.Println("Ошибка при чтении данных из файла:", err)
		return
	}

	printer.PrintStaff(staff)
}
