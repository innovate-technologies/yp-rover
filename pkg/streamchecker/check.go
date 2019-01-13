package streamchecker

import (
	"log"
	"strings"

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
	log.Println(content)

	if strings.Contains(content, "audio/mpeg") || strings.Contains(content, "audio/aacp") || strings.Contains(content, "audio/aac") || strings.Contains(content, "audio/ogg") || strings.Contains(content, "application/ogg") {
		return true
	}

	return false
}

// CheckValidPlaylist checks if URL serves a playlist file
func CheckValidPlaylist(url string) bool {
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(30)) // because radionomy
	r := resty.R()
	r.Header.Set("User-Agent", "VLC/3.0.4 LibVLC/3.0.4")
	resp, err := r.Head(url)
	if err != nil {
		return false
	}

	content := resp.Header().Get("content-type")
	log.Println(content)

	if strings.Contains(content, "audio/x-scpls") || strings.Contains(content, "audio/x-mpegurl") {
		log.Println("OK")
		return true
	}

	return false
}
