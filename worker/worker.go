package worker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"myproject/task"
)

type WorkerPool[T any] struct {
	Tasks   chan task.Task[T]
	WG      sync.WaitGroup
	Mutex   sync.Mutex
	Running int
}

func (wp *WorkerPool[T]) Start(numWorkers int, ctx context.Context) {
	for i := 0; i < numWorkers; i++ {
		wp.WG.Add(1)
		go wp.worker(i, ctx)
	}
}

func (wp *WorkerPool[T]) worker(id int, ctx context.Context) {
	defer wp.WG.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[Worker %d] Остановлен\n", id)
			return
		case t, ok := <-wp.Tasks:
			if !ok {
				return
			}

			wp.Mutex.Lock()
			wp.Running++
			wp.Mutex.Unlock()

			fmt.Printf("[Worker %d] Обрабатывает задачу: %s\n", id, t.Name)
			time.Sleep(time.Duration(2+id) * time.Second)

			wp.Mutex.Lock()
			wp.Running--
			wp.Mutex.Unlock()

			fmt.Printf("[Worker %d] Задача выполнена: %s\n", id, t.Name)
		}
	}
}
