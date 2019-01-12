package shoutcastcom

import "encoding/xml"

/*
	General
*/

// Station contains the data of a station listing
type Station struct {
	ID           string   `xml:"id,attr" json:"shoutcastID" bson:"shoutcastID"`
	Name         string   `xml:"name,attr" json:"name" bson:"name"`
	MediaType    string   `xml:"mt,attr" json:"mediaType" bson:"mediaType"`
	BitRate      string   `xml:"br,attr" json:"bitrate" bson:"bitrate"`
	Genre        string   `xml:"genre,attr" json:"genre" bson:"genre"`
	Genre2       string   `xml:"genre2,attr" json:"genre2" bson:"genre2"`
	Genre3       string   `xml:"genre3,attr" json:"genre3" bson:"genre3"`
	Genre4       string   `xml:"genre4,attr" json:"genre4" bson:"genre4"`
	Genre5       string   `xml:"genre5,attr" json:"genre5" bson:"genre5"`
	LogoURL      string   `xml:"logo,attr" json:"logoURL" bson:"logoURL"`
	CurrentTrack string   `xml:"ct,attr" json:"curentTrack" bson:"curentTrack"`
	ListnerCount int      `xml:"lc,attr" json:"listenerCount" bson:"listenerCount"`
	ListenURLs   []string `json:"listenURLs" bson:"listenURLs"`
}

type stationlist struct {
	XMLName  xml.Name  `xml:"stationlist"`
	Stations []Station `xml:"station"`
}

/*
	Genre list
*/
type genrelist struct {
	XMLName xml.Name        `xml:"genrelist"`
	Genres  []genrelistItem `xml:"genre"`
}

type genrelistItem struct {
	Name  string `xml:"name,attr"`
	Count string `xml:"count,attr"`
}
