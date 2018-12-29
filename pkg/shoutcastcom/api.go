package shoutcastcom

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"

	"gopkg.in/resty.v1"
)

// SHOUTcastAPIClient is a client for the SHOUTcast.com (legacy) directory
type SHOUTcastAPIClient struct {
	key string
}

// NewClient gives a SHOUTcastAPIClient instance
func NewClient(key string) *SHOUTcastAPIClient {
	return &SHOUTcastAPIClient{
		key: key,
	}
}

func (s *SHOUTcastAPIClient) doRequest(endpoint string, params ...map[string]string) (*resty.Response, error) {
	r := resty.R()
	r = r.SetQueryParam("k", s.key)
	for _, paramSet := range params {
		r = r.SetQueryParams(paramSet)
	}

	resp, err := r.Get(fmt.Sprintf("https://api.shoutcast.com/legacy/%s", endpoint))

	if resp.StatusCode() != http.StatusOK {
		return resp, fmt.Errorf("HTTP error %d: %s", resp.StatusCode(), string(resp.Body()))
	}

	return resp, err
}

// GeTuneInURLs gives the URLs or a certain SHOUTcast ID
func (s *SHOUTcastAPIClient) GeTuneInURLs(id string) ([]string, error) {
	r := resty.R()
	r = r.SetQueryParam("id", id)

	resp, err := r.Get("https://yp.shoutcast.com/sbin/tunein-station.m3u")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode(), string(resp.Body()))
	}

	body := string(resp.Body())
	lines := strings.Split(body, "\n")

	out := []string{}

	for _, line := range lines {
		if govalidator.IsURL(line) {
			out = append(out, line)
		}
	}

	if len(out) == 0 {
		return nil, errors.New("No URLs found")
	}
	return out, nil
}

// GetAllGenres returns a list with all available genres
func (s *SHOUTcastAPIClient) GetAllGenres() ([]string, error) {
	resp, err := s.doRequest("genrelist")
	if err != nil {
		return nil, err
	}

	list := genrelist{}
	xml.Unmarshal(resp.Body(), &list) // error not handled due weird output of actual API

	out := []string{}
	for _, item := range list.Genres {
		out = append(out, item.Name)
	}

	if len(out) == 0 {
		return nil, errors.New("Invalid XML response")
	}

	return out, nil
}

// GetTop500 gets the top 500 stations
// All paramaters are optional and can be left out if an empty string is used (TODO: improve these option setters)
func (s *SHOUTcastAPIClient) GetTop500(mediaType string, bitRate string, limit string) ([]Station, error) {
	resp, err := s.doRequest("Top500", map[string]string{
		"limit": limit,
		"br":    bitRate,
		"mt":    mediaType,
	})

	if err != nil {
		return nil, err
	}

	list := stationlist{}
	xml.Unmarshal(resp.Body(), &list) // error not handled due weird output of actual API

	return list.Stations, nil
}

// GetByGenre gets the stations by genre
// All paramaters except genre are optional and can be left out if an empty string is used (TODO: improve these option setters)
func (s *SHOUTcastAPIClient) GetByGenre(genre string, mediaType string, bitRate string, limit string) ([]Station, error) {
	resp, err := s.doRequest("genresearch", map[string]string{
		"genre": genre,
		"limit": limit,
		"br":    bitRate,
		"mt":    mediaType,
	})

	if err != nil {
		return nil, err
	}

	list := stationlist{}
	xml.Unmarshal(resp.Body(), &list) // error not handled due weird output of actual API

	return list.Stations, nil
}
