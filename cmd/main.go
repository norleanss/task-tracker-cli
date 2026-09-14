package main

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker-cli/internal/service"
)

func main() {
	if len(os.Args) < 2 {
		println("задайте параметры")
		return
	}
	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Укажите описание задачи")
			return
		}
		err := service.Add(os.Args[2])
		if err != nil {
			fmt.Println("Ошибка ", err)
		} else {
			fmt.Println("Успешное добавление")
		}
	case "list":
		status := ""
		if len(os.Args) > 2 {
			status = os.Args[2]
		}
		tasks, err := service.List(status)
		if err != nil {
			fmt.Println("Ошибка ", err)
			return
		} else {
			fmt.Println(tasks)
		}
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Укажите id и новое описание задачи")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Некорректный id: ", err)
			return
		}
		err = service.Update(id, os.Args[3])
		if err != nil {
			fmt.Println("Ошибка ", err)
		} else {
			fmt.Println("Успешное обновление")
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Укажите id задачи")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Неккоректный id")
			return
		}
		err = service.Delete(id)
		if err != nil {
			fmt.Println("Ошибка ", err)
		} else {
			fmt.Println("Успешное удаление")
		}
	case "mark":
		if len(os.Args) < 4 {
			fmt.Println("Укажите id и статус задачи")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Неккоретный id")
			return
		}
		err = service.Mark(id, os.Args[3])
		if err != nil {
			fmt.Println("Ошибка ", err)
		} else {
			fmt.Println("Статус успешно изменен")
		}
	}
}
