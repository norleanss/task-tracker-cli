package service

import (
	"fmt"
	"task-tracker-cli/internal/storage"
	"task-tracker-cli/model"
	"time"
)

func Add(desctription string) error {
	load, err := storage.Load()
	if err != nil {
		return err
	}
	task := model.Task{
		ID:          len(load) + 1,
		Description: desctription,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	load = append(load, task)
	if err = storage.Save(load); err != nil {
		return err
	}
	return nil
}

func Update(id int, description string) error {
	load, err := storage.Load()
	if err != nil {
		return err
	}
	index, err := getTask(load, id)
	if err != nil {
		return err
	}
	load[index].Description = description
	load[index].UpdatedAt = time.Now()
	if err = storage.Save(load); err != nil {
		return err
	}
	return nil
}

func Delete(id int) error {
	load, err := storage.Load()
	if err != nil {
		return err
	}
	index, err := getTask(load, id)
	if err != nil {
		return err
	}
	load = append(load[:index], load[index+1:]...)
	if err = storage.Save(load); err != nil {
		return err
	}
	return nil
}

func Mark(id int, status string) error {
	load, err := storage.Load()
	if err != nil {
		return err
	}
	index, err := getTask(load, id)
	if err != nil {
		return err
	}
	load[index].Status = status
	load[index].UpdatedAt = time.Now()
	if err = storage.Save(load); err != nil {
		return err
	}
	return nil

}

func List(status string) ([]model.Task, error) {
	load, err := storage.Load()
	if err != nil {
		return []model.Task{}, err
	}
	tasks := []model.Task{}
	if status == "" {
		return load, nil
	}
	for _, value := range load {
		if value.Status == status {
			tasks = append(tasks, value)
		}
	}
	return tasks, nil
}

func getTask(tasks []model.Task, id int) (int, error) {
	for i, task := range tasks {
		if task.ID == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("задача не найдена")
}
