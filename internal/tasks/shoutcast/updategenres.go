package shoutcast

import (
	"log"

	"github.com/innovate-technologies/yp-rover/internal/tasks"
	"github.com/innovate-technologies/yp-rover/pkg/shoutcastcom"
	"github.com/innovate-technologies/yp-rover/pkg/store"
)

// UpdateGenres fetches and updates the genre list in the database
func (t *Task) UpdateGenres() ([]tasks.Task, error) {
	api := shoutcastcom.NewClient(t.config.ShoutcastKey)
	genres, err := api.GetAllGenres()
	if err != nil {
		return nil, err
	}

	db, err := store.New(t.config.MySQLURL)
	if err != nil {
		return nil, err
	}

	for _, genre := range genres {
		err = db.AddSHOUTcastGenre(genre)
		if err != nil {
			log.Println(err)
		}
	}

	return nil, nil
}
