package tuneintasks

import (
	"log"

	"github.com/innovate-technologies/yp-rover/internal/tasks"
	"github.com/innovate-technologies/yp-rover/pkg/store"
	"github.com/innovate-technologies/yp-rover/pkg/tunein"
)

// UpdateGenres fetches and updates the genre list in the database
func (t *Task) UpdateGenres() ([]tasks.Task, error) {
	api := tunein.NewClient()
	genres, err := api.GetGenreGuides()
	if err != nil {
		return nil, err
	}

	db, err := store.New(t.config)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	nt := []tasks.Task{}
	for name, gid := range genres {
		log.Printf("Adding TuneIn genre %s\n", gid)
		err = db.AddTuneInGenre(tunein.Genre{
			Name:    name,
			GuideID: gid,
		})
		if err != nil {
			log.Println(err)
		}

		log.Printf("Queue TuneIn station fetch for %s\n", gid)
		nt = append(nt, tasks.Task{
			Unit:     "tunein",
			Function: "UpdateStations",
			Args: map[string]string{
				"genre":  gid,
				"offset": "0",
			},
		})
	}

	return nt, nil
}
