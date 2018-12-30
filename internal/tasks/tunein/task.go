package tuneintasks

import (
	"errors"
	"strconv"

	"github.com/innovate-technologies/yp-rover/internal/tasks"
)

// Task is the task handler for TuneIn
type Task struct {
}

// New gives a new task hander
func New() Task {
	return Task{}
}

// HandleTask handles a given task
func (t *Task) HandleTask(task tasks.Task) ([]tasks.Task, error) {
	switch task.Function {
	case "UpdateGenres":
		return t.UpdateGenres()
	case "UpdateStations":
		offset, _ := strconv.ParseInt(task.Args["offset"], 10, 64)
		return t.FetchForGenre(task.Args["genre"], offset)
	}
	return nil, errors.New("Task not found")
}
