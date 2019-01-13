package tunein

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	resty "gopkg.in/resty.v1"
)

const waitTime = 4 * time.Minute

// API is an API wrapper for opml.radiotime.com
// opml.radiotime.com is a deprecated endpoint, we do not guarantee this will stay up
type API struct {
}

// NewClient gives a new API instance
func NewClient() *API {
	return &API{}
}

func (a *API) doRequest(endpoint string, params ...map[string]string) (*resty.Response, error) {
	endpoint = strings.Replace(endpoint, "http://opml.radiotime.com/", "", -1) // add ability to pass the direct URL given in the OPML
	//http://opml.radiotime.com
	r := resty.R()
	r = r.SetHeaders(map[string]string{
		"Connection":      "keep-alive",
		"Cache-Control":   "max-age=0",
		"User-Agent":      "Mozilla/5.0 (X11; Linux i686) AppleWebKit/534.24 (KHTML, like Gecko) Chrome/11.0.696.71 Safari/534.24",
		"Accept":          "application/xml,application/xhtml+xml,text/html;q=0.9,text/plain;q=0.8,image/png,*/*;q=0.5",
		"Accept-Encoding": "gzip",
		"Accept-Language": "en-US,en;q=0.8",
		"Accept-Charset":  "ISO-8859-1,utf-8;q=0.7,*;q=0.3",
	})

	for _, paramSet := range params {
		r = r.SetQueryParams(paramSet)
	}

	resp, err := r.Get(fmt.Sprintf("https://opml.radiotime.com/%s", endpoint))

	if resp.StatusCode() == 403 { // we have been rate limited
		log.Println("Hit rate limit... sleeping")
		time.Sleep(waitTime)
		return a.doRequest(endpoint)
	}

	if resp.StatusCode() != http.StatusOK {
		return resp, fmt.Errorf("HTTP error %d: %s", resp.StatusCode(), string(resp.Body()))
	}

	return resp, err
}

// GetGenreGuides gets a list of all guide IDs for genres
func (a *API) GetGenreGuides() (map[string]string, error) {
	resp, err := a.doRequest("Browse.ashx", map[string]string{
		"c": "music",
	})
	if err != nil {
		return nil, err
	}

	o, err := parseOPML(resp.Body())
	if err != nil {
		return nil, err
	}

	genres := map[string]string{}

	for _, entry := range o.Body.Outlines {
		if entry.Type == "link" {
			genres[entry.Text] = entry.GuideID
		}
	}

	return genres, nil
}

// BrowseStations will give back radio stations in a given guide
func (a *API) BrowseStations(guide string, offset int64) ([]Station, int64, error) {
	resp, err := a.doRequest("Browse.ashx", map[string]string{
		"id":     guide,
		"offset": fmt.Sprintf("%d", offset),
		"filter": "s", // stations only
	})

	if err != nil {
		return nil, 0, err
	}

	o, err := parseOPML(resp.Body())
	if err != nil {
		return nil, 0, err
	}

	stations := []Station{}
	var newOffset int64

	for _, entry := range o.Body.Outlines {
		if entry.Type != "audio" {
			if entry.Key == "nextStations" {
				u, err := url.Parse(entry.URL)
				if err != nil {
					continue
				}
				offsetStrinf := u.Query().Get("offset")
				off, err := strconv.ParseInt(offsetStrinf, 10, 64)
				if err != nil {
					continue
				}
				newOffset = off
			}
			continue
		}
		time.Sleep(time.Second) // preventing a rate limit here
		tuneResp, err := a.doRequest(entry.URL)
		if err != nil {
			continue
		}

		stations = append(stations, Station{
			Name:                 entry.Text,
			MediaType:            entry.Formats,
			Logo:                 entry.Image,
			CurrentTrack:         entry.Playing,
			CurrentTrackImageURL: entry.PlayingImage,
			Reliability:          entry.Reliability,
			GenreID:              entry.GenreID,
			BitRate:              entry.Bitrate,
			TuneInURL:            string(tuneResp.Body()),
		})
	}

	return stations, newOffset, nil
}
