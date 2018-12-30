package tunein

/*
	This file is based on github.com/gilliek/go-opml
	However this is modified to follow the format used in the RadioTime OPML
*/

import "encoding/xml"

type opml struct {
	XMLName xml.Name `xml:"opml"`
	Version string   `xml:"version,attr"`
	Head    head     `xml:"head"`
	Body    body     `xml:"body"`
}

// head holds some meta information about the document.
type head struct {
	Title           string `xml:"title"`
	DateCreated     string `xml:"dateCreated,omitempty"`
	DateModified    string `xml:"dateModified,omitempty"`
	OwnerName       string `xml:"ownerName,omitempty"`
	OwnerEmail      string `xml:"ownerEmail,omitempty"`
	OwnerID         string `xml:"ownerId,omitempty"`
	Docs            string `xml:"docs,omitempty"`
	ExpansionState  string `xml:"expansionState,omitempty"`
	VertScrollState string `xml:"vertScrollState,omitempty"`
	WindowTop       string `xml:"windowTop,omitempty"`
	WindowBottom    string `xml:"windowBottom,omitempty"`
	WindowLeft      string `xml:"windowLeft,omitempty"`
	WindowRight     string `xml:"windowRight,omitempty"`
}

// body is the parent structure of all outlines.
type body struct {
	Outlines []outline `xml:"outline"`
}

// outline holds all information about an outline.
type outline struct {
	Outlines     []outline `xml:"outline"`
	Text         string    `xml:"text,attr"`
	URL          string    `xml:"URL,attr,omitempty"`
	Type         string    `xml:"type,attr,omitempty"`
	Bitrate      string    `xml:"bitrate,attr,omitempty"`
	Reliability  string    `xml:"reliability,attr,omitempty"`
	GuideID      string    `xml:"guide_id,attr,omitempty"`
	SubText      string    `xml:"subtext,attr,omitempty"`
	GenreID      string    `xml:"genre_id,attr,omitempty"`
	Formats      string    `xml:"formats,attr,omitempty"`
	Playing      string    `xml:"playing,attr,omitempty"`
	PlayingImage string    `xml:"playing_image,attr,omitempty"`
	Image        string    `xml:"image,attr,omitempty"`
	Key          string    `xml:"key,attr,omitempty"`
}

func parseOPML(b []byte) (*opml, error) {
	var root opml
	err := xml.Unmarshal(b, &root)
	if err != nil {
		return nil, err
	}

	return &root, nil
}
