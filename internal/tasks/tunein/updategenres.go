package tuneintasks

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/innovate-technologies/yp-rover/internal/tasks"
	"github.com/innovate-technologies/yp-rover/pkg/tunein"
)

// UpdateGenres fetches and updates the genre list in the database
func (t *Task) UpdateGenres() ([]tasks.Task, error) {
	api := tunein.NewClient()
	genres, err := api.GetGenreGuides()
	if err != nil {
		return nil, err
	}

	spew.Dump(genres)

	return nil, nil
}
