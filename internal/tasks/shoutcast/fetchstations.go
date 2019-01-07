package shoutcast

import (
	"fmt"
	"log"
	"time"

	"github.com/innovate-technologies/yp-rover/internal/tasks"
	"github.com/innovate-technologies/yp-rover/pkg/shoutcastcom"
	"github.com/innovate-technologies/yp-rover/pkg/store"
	"github.com/innovate-technologies/yp-rover/pkg/streamchecker"
)

// FetchForGenre fetches 50 stations for a genre and dispatches a new job if needed
func (t *Task) FetchForGenre(genre string, offset int64) ([]tasks.Task, error) {
	api := shoutcastcom.NewClient(t.config.ShoutcastKey)
	db, err := store.New(t.config.MySQLURL)
	if err != nil {
		return nil, err
	}

	log.Printf("Getting stations for genre %s with offset %d", genre, offset)
	stations, err := api.GetByGenre(genre, "", "", fmt.Sprintf("%d,50", offset))
	if err != nil {
		return nil, err
	}

	for _, station := range stations {
		log.Printf("Checking station %s", station.ID)
		urls, _ := api.GeTuneInURLs(station.ID)
		time.Sleep(300 * time.Millisecond) // for rate limiting
		if urls == nil || len(urls) == 0 { // non existing station
			continue
		}
		valid := []string{}
		for _, url := range urls {
			if streamchecker.CheckValidStream(url) {
				valid = append(valid, url)
			}
		}
		station.ListenURLs = valid
		time.Sleep(200 * time.Millisecond) // try not to fetch too fast
		if len(valid) > 0 {
			log.Printf("Saving station %s", station.ID)
			err := db.AddSHOUTcastStation(station)
			if err != nil {
				log.Println(err)
			}
		}
	}

	if len(stations) != 0 {
		log.Printf("More stations for genre %s to be fetched!", genre)
		time.Sleep(5 * time.Second) // try not to fetch too fast
		return []tasks.Task{
			tasks.Task{
				Unit:     "shoutcastcom",
				Function: "UpdateStations",
				Args: map[string]string{
					"genre":  genre,
					"offset": fmt.Sprintf("%d", offset+50),
				},
			},
		}, nil
	}

	return nil, nil
}
