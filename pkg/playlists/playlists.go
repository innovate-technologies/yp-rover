package playlists

import (
	"errors"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
	resty "gopkg.in/resty.v1"
)

var plsFile = regexp.MustCompile(`File.*=(.*)$`)

func GetEntryURLs(url string) ([]string, error) {
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(30)) // because radionomy
	r := resty.R()
	r.Header.Set("User-Agent", "VLC/3.0.4 LibVLC/3.0.4")
	resp, err := r.Get(url)
	if err != nil {
		return nil, err
	}

	content := resp.Header().Get("content-type")

	if strings.Contains(content, "audio/x-mpegurl") {
		lines := strings.Split(string(resp.Body()), "\n")

		out := []string{}

		for _, line := range lines {
			if govalidator.IsURL(line) {
				out = append(out, line)
			}
		}

		return out, nil
	}

	if strings.Contains(content, "audio/x-scpls") {
		lines := strings.Split(string(resp.Body()), "\n")
		out := []string{}

		for _, line := range lines {
			matched := plsFile.FindAllStringSubmatch(line, -1)
			if matched == nil || len(matched) == 0 {
				continue
			}
			if len(matched[0]) < 2 {
				continue // no correct submatch
			}

			if govalidator.IsURL(matched[0][1]) {
				out = append(out, matched[0][1])
			}
		}

		return out, nil
	}

	return nil, errors.New("No playlist found")
}
