package shoutcast

import (
	"errors"

	"github.com/innovate-technologies/yp-rover/internal/tasks"
)

// Task is the task handler for SHOUTcast.com
type Task struct {
	key string
}

// New gives a new task hander
func New(key string) Task {
	return Task{
		key: key,
	}
}

// HandleTask handles a given task
func (t *Task) HandleTask(task tasks.Task) ([]tasks.Task, error) {
	switch task.Function {
	case "UpdateGenres":
		return t.UpdateGenres()
	}
	return nil, errors.New("Task not found")
}
