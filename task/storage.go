package task

import (
	"encoding/json"
	"os"
)

type TaskStorage[T any] struct {
	Tasks []Task[T] `json:"tasks"`
}

func (s *TaskStorage[T]) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(s, "", " ")

	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (s *TaskStorage[T]) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, s)
}
