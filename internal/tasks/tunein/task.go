package tuneintasks

import (
	"errors"
	"strconv"

	"github.com/innovate-technologies/yp-rover/internal/config"
	"github.com/innovate-technologies/yp-rover/internal/tasks"
)

// Task is the task handler for TuneIn
type Task struct {
	config config.Config
}

// New gives a new task hander
func New(config config.Config) Task {
	return Task{
		config: config,
	}
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
