package tunein

/*
	General
*/

// Station contains the data of a station listing
type Station struct {
	Name                 string   `json:"name" bson:"name"`
	Logo                 string   `json:"logo" bson:"logo"`
	MediaType            string   `json:"mediaType" bson:"mediaType"`
	BitRate              string   `json:"bitrate" bson:"bitrate"`
	GenreID              string   `json:"GenreID" bson:"GenreID"`
	Reliability          string   `json:"reliability" bson:"reliability"`
	CurrentTrack         string   `json:"currentTrack" bson:"currentTrack"`
	CurrentTrackImageURL string   `json:"currentTrackImageURL" bson:"currentTrackImageURL"`
	TuneInURLs           []string `json:"tuneInURL" bson:"tuneInURL"`
}

// Genre represents the TuneIn representatation of a genre
type Genre struct {
	Name    string `json:"name" bson:"name"`
	GuideID string `json:"guideID" bson:"guideID"`
}
