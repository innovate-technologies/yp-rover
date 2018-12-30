package tuneintasks

import (
	"fmt"
	"time"

	"github.com/innovate-technologies/yp-rover/internal/tasks"
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

	//spew.Dump(stations)
	for _, station := range stations {
		if !streamchecker.CheckValidStream(station.TuneInURL) {
			continue
		}
		time.Sleep(200 * time.Millisecond) // try not to fetch too fast
		// TODO: add me to a database
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
