package main

import (
	"fmt"

	"github.com/mussyaroslav/home_work_basic/hw02_fix_app/printer"
	"github.com/mussyaroslav/home_work_basic/hw02_fix_app/reader"
	"github.com/mussyaroslav/home_work_basic/hw02_fix_app/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Введите путь к файлу данных: ")
	fmt.Scanln(&path)

	if len(path) == 0 {
		path = "data.json"
	}

	var err error
	var staff []types.Employee

	staff, err = reader.ReadJSON(path)
	if err != nil {
		fmt.Println("Ошибка при чтении данных из файла:", err)
		return
	}

	printer.PrintStaff(staff)
}
