package streamchecker

import (
	"context"
	"strings"
	"time"

	resty "gopkg.in/resty.v1"
)

// CheckValidStream will call a stream URL and check if a radio stream is present
func CheckValidStream(streamurl string) bool {
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(30)) // because radionomy
	r := resty.R()
	r.Header.Set("User-Agent", "VLC/3.0.4 LibVLC/3.0.4")
	resp, err := r.Head(streamurl)

	var content string
	if err == nil { // a version of SHOUTcast returns invalid HTTP on HEAD
		content = resp.Header().Get("content-type")
	}

	if resp.StatusCode() == 400 || content == "" {
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(2 * time.Second)
			cancel()
		}()
		r := resty.R()
		r.Header.Set("User-Agent", "VLC/3.0.4 LibVLC/3.0.4")
		r.SetContext(ctx)
		r = r.SetOutput("/dev/null")
		resp, _ := r.Get(streamurl)
		content = resp.Header().Get("content-type")
		resp.RawBody().Close()
	}

	if strings.Contains(content, "audio/mpeg") || strings.Contains(content, "audio/aacp") || strings.Contains(content, "audio/aac") || strings.Contains(content, "audio/ogg") || strings.Contains(content, "application/ogg") {
		return true
	}

	return false
}

// CheckValidPlaylist checks if URL serves a playlist file
func CheckValidPlaylist(url string) bool {
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(30)) // because radionomy
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()
	r := resty.R()
	r.Header.Set("User-Agent", "VLC/3.0.4 LibVLC/3.0.4")
	r.SetContext(ctx) // if it starts sending audio this is a good thing to have
	resp, err := r.Get(url)
	if err != nil {
		return false
	}

	content := resp.Header().Get("content-type")

	if strings.Contains(content, "audio/x-scpls") || strings.Contains(content, "audio/x-mpegurl") {
		return true
	}

	return false
}
