package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"myproject/task"
	"myproject/worker"
)

func main() {
	tasks := []task.Task[string]{
		{ID: 1, Name: "Отправка email", Status: "pending", Data: "hello@example.com"},
		{ID: 2, Name: "Обработка платежа", Status: "pending", Data: "12345"},
		{ID: 3, Name: "Генерация отчета", Status: "pending", Data: "sales.xlsx"},
		{ID: 4, Name: "sdjjjj ", Status: "pending", Data: "sales.xlsx"},
		{ID: 5, Name: "qqqqq qqq q", Status: "pending", Data: "sales.xlsx"},
		{ID: 6, Name: "aaaa asdsd", Status: "pending", Data: "sales.xlsx"},
		{ID: 7, Name: "ppppop ppop", Status: "pending", Data: "sales.xlsx"},
	}

	storage := task.TaskStorage[string]{Tasks: tasks}
	storage.SaveToFile("tasks.json")

	storage.LoadFromFile("task.json")
	fmt.Println(storage.Tasks)

	taskChan := make(chan task.Task[string], len(tasks))
	for _, t := range storage.Tasks {
		taskChan <- t
	}
	close(taskChan)

	wp := worker.WorkerPool[string]{Tasks: taskChan}

	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nОстановка воркеров...")
		cancel()
	}()

	numWorkers := 3
	wp.Start(numWorkers, ctx)

	wp.WG.Wait()
	fmt.Println("Все задачи обработаны. Завершение.")
}
