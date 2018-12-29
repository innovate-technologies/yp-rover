package streamchecker

import (
	resty "gopkg.in/resty.v1"
)

// CheckValidStream will call a stream URL and check if a radio stream is present
func CheckValidStream(url string) bool {
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(30)) // because radionomy
	r := resty.R()
	r.Header.Set("User-Agent", "VLC/3.0.4 LibVLC/3.0.4")
	resp, err := r.Head(url)
	if err != nil {
		return false
	}

	content := resp.Header().Get("content-type")

	if content == "audio/mpeg" || content == "audio/aacp" || content == "audio/aac" || content == "audio/ogg" || content == "application/ogg" {
		return true
	}

	return false
}
