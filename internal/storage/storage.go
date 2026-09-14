package storage

import (
	"encoding/json"
	"io"
	"os"
	"task-tracker-cli/model"
)

const fileName = "data.json"

func Load() ([]model.Task, error) {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Task{}, err
		}
		return []model.Task{}, err
	}
	defer file.Close()

	var tasks []model.Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		if err == io.EOF {
			return []model.Task{}, nil
		}
		return nil, err
	}
	return tasks, nil
}

func Save(tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}
