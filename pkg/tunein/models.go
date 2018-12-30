package tunein

/*
	General
*/

// Station contains the data of a station listing
type Station struct {
	Name                 string
	Logo                 string
	MediaType            string
	BitRate              string
	GenreID              string
	Reliability          string
	CurrentTrack         string
	CurrentTrackImageURL string
	TuneInURL            string
}
