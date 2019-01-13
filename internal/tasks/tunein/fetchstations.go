package tuneintasks

import (
	"fmt"
	"log"
	"time"

	"github.com/innovate-technologies/yp-rover/internal/tasks"
	"github.com/innovate-technologies/yp-rover/pkg/store"
	"github.com/innovate-technologies/yp-rover/pkg/streamchecker"
	"github.com/innovate-technologies/yp-rover/pkg/tunein"
)

// FetchForGenre fetches tations for a genre and dispatches a new job if needed
func (t *Task) FetchForGenre(genre string, offset int64) ([]tasks.Task, error) {
	api := tunein.NewClient()
	stations, newOffset, err := api.BrowseStations(genre, offset)
	if err != nil {
		return nil, err
	}

	db, err := store.New(t.config)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	//spew.Dump(stations)
	for _, station := range stations {
		if !streamchecker.CheckValidStream(station.TuneInURL) {
			continue
		}
		time.Sleep(200 * time.Millisecond) // try not to fetch too fast
		log.Printf("Saving station %s", station.Name)
		err := db.AddTuneInStation(station)
		if err != nil {
			log.Println(err)
		}
	}

	if len(stations) != 0 {
		time.Sleep(5 * time.Second) // try not to fetch too fast
		return []tasks.Task{
			tasks.Task{
				Unit:     "tunein",
				Function: "UpdateStations",
				Args: map[string]string{
					"genre":  genre,
					"offset": fmt.Sprintf("%d", newOffset),
				},
			},
		}, nil
	}

	return nil, nil
}
