package shoutcast

import (
	"log"

	"github.com/innovate-technologies/yp-rover/internal/tasks"
	"github.com/innovate-technologies/yp-rover/pkg/shoutcastcom"
	"github.com/innovate-technologies/yp-rover/pkg/store"
)

// UpdateGenres fetches and updates the genre list in the database
func (t *Task) UpdateGenres() ([]tasks.Task, error) {
	log.Println("Updating SHOUTcast genres")
	api := shoutcastcom.NewClient(t.config.ShoutcastKey)
	genres, err := api.GetAllGenres()
	if err != nil {
		return nil, err
	}

	db, err := store.New(t.config)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	nt := []tasks.Task{}
	for _, genre := range genres {
		log.Printf("Adding SHOUTcast genre %s\n", genre)
		err = db.AddSHOUTcastGenre(genre)
		if err != nil {
			log.Println(err)
		}

		log.Printf("Queue SHOUTcast station fetch for %s\n", genre)
		nt = append(nt, tasks.Task{
			Unit:     "shoutcastcom",
			Function: "UpdateStations",
			Args: map[string]string{
				"genre":  genre,
				"offset": "0",
			},
		})
	}

	return nt, nil
}
