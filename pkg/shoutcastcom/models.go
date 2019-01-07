package shoutcastcom

import "encoding/xml"

/*
	General
*/

// Station contains the data of a station listing
type Station struct {
	ID           string `xml:"id,attr"`
	Name         string `xml:"name,attr"`
	MediaType    string `xml:"mt,attr"`
	BitRate      string `xml:"br,attr"`
	Genre        string `xml:"genre,attr"`
	Genre2       string `xml:"genre2,attr"`
	Genre3       string `xml:"genre3,attr"`
	Genre4       string `xml:"genre4,attr"`
	Genre5       string `xml:"genre5,attr"`
	LogoURL      string `xml:"logo,attr"`
	CurrentTrack string `xml:"ct,attr"`
	ListnerCount int    `xml:"lc,attr"`
	ListenURLs   []string
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
