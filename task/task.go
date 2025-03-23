package task

import "encoding/json"

type Task[T any] struct {
	ID     int    `json:id`
	Name   string `json:name`
	Status string `json:status` // "pending", "running", "done"
	Data   T      `json:data`
}

func (t *Task[T]) ToJson() string {
	data, _ := json.MarshalIndent(t, "", " ")
	return string(data)
}
