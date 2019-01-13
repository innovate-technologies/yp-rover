package tuneintasks

import (
	"fmt"
	"log"
	"time"

	"github.com/innovate-technologies/yp-rover/internal/tasks"
	"github.com/innovate-technologies/yp-rover/pkg/playlists"
	"github.com/innovate-technologies/yp-rover/pkg/store"
	"github.com/innovate-technologies/yp-rover/pkg/streamchecker"
	"github.com/innovate-technologies/yp-rover/pkg/tunein"
)

// FetchForGenre fetches tations for a genre and dispatches a new job if needed
func (t *Task) FetchForGenre(genre string, offset int64) ([]tasks.Task, error) {
	log.Printf("Getting TuneIn stations for genre %s and offset %d\n", genre, offset)
	api := tunein.NewClient()
	stations, newOffset, err := api.BrowseStations(genre, offset)
	log.Println("Getting TuneIn station list")
	if err != nil {
		log.Printf("Got error: %s\n", err)
		return nil, err
	}

	db, err := store.New(t.config)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	//spew.Dump(stations)
	for _, station := range stations {
		log.Printf("Checking %s\n", station.Name)
		valid := []string{}
		for _, url := range station.TuneInURLs {
			if streamchecker.CheckValidStream(url) {
				valid = append(valid, url)
			} else if streamchecker.CheckValidPlaylist(url) {
				entries, err := playlists.GetEntryURLs(url)
				if err == nil {
					valid = append(valid, entries...)
				}
			}
		}

		if len(valid) == 0 {
			continue
		}
		station.TuneInURLs = valid

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
