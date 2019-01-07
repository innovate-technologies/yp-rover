package shoutcast

import (
	"fmt"
	"time"

	"github.com/innovate-technologies/yp-rover/internal/tasks"
	"github.com/innovate-technologies/yp-rover/pkg/shoutcastcom"
	"github.com/innovate-technologies/yp-rover/pkg/streamchecker"
)

// FetchForGenre fetches 50 stations for a genre and dispatches a new job if needed
func (t *Task) FetchForGenre(genre string, offset int64) ([]tasks.Task, error) {
	api := shoutcastcom.NewClient(t.config.ShoutcastKey)
	stations, err := api.GetByGenre(genre, "", "", fmt.Sprintf("%d,50", offset))
	if err != nil {
		return nil, err
	}

	//spew.Dump(stations)
	for _, station := range stations {
		urls, _ := api.GeTuneInURLs(station.ID)
		if urls == nil || len(urls) == 0 { // non existing station
			continue
		}
		valid := []string{}
		for _, url := range urls {
			if streamchecker.CheckValidStream(url) {
				valid = append(valid, url)
			}
		}
		time.Sleep(200 * time.Millisecond) // try not to fetch too fast
		// TODO: add me to a database
	}

	if len(stations) != 0 {
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
