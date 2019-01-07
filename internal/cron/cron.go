package cron

import (
	"time"

	"github.com/innovate-technologies/yp-rover/internal/tasks"
)

// UpdateGenres gives a task back to update the genre list every 24 hours
func UpdateGenres() chan tasks.Task {
	out := make(chan tasks.Task)
	task := tasks.Task{
		Unit:     "shoutcastcom",
		Function: "UpdateGenres",
	}

	timer := time.Tick(24 * time.Hour)
	//go sendTaskOnTick(task, timer, out)

	go func() {
		//out <- task
	}()
	return out
}

func sendTaskOnTick(task tasks.Task, tick <-chan time.Time, to chan tasks.Task) {
	for {
		<-tick
		to <- task
	}
}
