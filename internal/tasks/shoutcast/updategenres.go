package shoutcast

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/innovate-technologies/yp-rover/internal/tasks"
	"github.com/innovate-technologies/yp-rover/pkg/shoutcastcom"
)

// UpdateGenres fetches and updates the genre list in the database
func (t *Task) UpdateGenres() ([]tasks.Task, error) {
	api := shoutcastcom.NewClient(t.key)
	genres, err := api.GetAllGenres()
	if err != nil {
		return nil, err
	}

	spew.Dump(genres)

	return nil, nil
}
