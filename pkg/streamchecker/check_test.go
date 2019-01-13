package streamchecker

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestCheckValidStream(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("HEAD", "http://listen.radionomy.com:80/Maartje",
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(http.StatusPermanentRedirect, "")
			res.Header.Set("Location", "http://streaming.radionomy.com/Maartje")
			return res, nil
		})
	httpmock.RegisterResponder("HEAD", "http://streaming.radionomy.com/Maartje",
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(http.StatusOK, "")
			res.Header.Set("Content-Type", "audio/mpeg")
			return res, nil
		})

	httpmock.RegisterResponder("HEAD", "http://icecast.com/stream",
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(http.StatusOK, "")
			res.Header.Set("Content-Type", "audio/ogg")
			return res, nil
		})

	httpmock.RegisterResponder("HEAD", "http://scserv.com/;",
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(http.StatusOK, "")
			res.Header.Set("Content-Type", "audio/ogg")
			return res, nil
		})

	httpmock.RegisterResponder("HEAD", "http://pdfstream.com/;",
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(http.StatusOK, "")
			res.Header.Set("Content-Type", "application/pdf")
			return res, nil
		})

	httpmock.RegisterResponder("HEAD", "http://icecastnohead.com/",
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(http.StatusBadRequest, "")
			return res, nil
		})
	httpmock.RegisterResponder("GET", "http://icecastnohead.com/",
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(http.StatusOK, "")
			res.Header.Set("Content-Type", "audio/ogg")
			return res, nil
		})

	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Radionomy with redirect",
			args: args{
				url: "http://listen.radionomy.com:80/Maartje",
			},
			want: true,
		},
		{
			name: "Icecast",
			args: args{
				url: "http://icecast.com/stream",
			},
			want: true,
		},
		{
			name: "Icecast whithout HEAD support",
			args: args{
				url: "http://icecastnohead.com/",
			},
			want: true,
		},
		{
			name: "SHOUTcast",
			args: args{
				url: "http://scserv.com/;",
			},
			want: true,
		},
		{
			name: "PDF FM",
			args: args{
				url: "http://pdfstream.com/;",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckValidStream(tt.args.url); got != tt.want {
				t.Errorf("CheckValidStream() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckValidPlaylist(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://listen.radionomy.com:80/Maartje",
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(http.StatusPermanentRedirect, "")
			res.Header.Set("Location", "http://streaming.radionomy.com/Maartje.pls")
			return res, nil
		})
	httpmock.RegisterResponder("GET", "http://streaming.radionomy.com/Maartje.pls",
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(http.StatusOK, "")
			res.Header.Set("Content-Type", "audio/x-scpls")
			return res, nil
		})

	httpmock.RegisterResponder("GET", "http://icecast.com/stream.m3u",
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(http.StatusOK, "")
			res.Header.Set("Content-Type", "audio/x-mpegurl")
			return res, nil
		})

	httpmock.RegisterResponder("GET", "http://scserv.com/.pls",
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(http.StatusOK, "")
			res.Header.Set("Content-Type", "audio/x-scpls")
			return res, nil
		})

	httpmock.RegisterResponder("GET", "http://pdfstream.com/;",
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(http.StatusOK, "")
			res.Header.Set("Content-Type", "application/pdf")
			return res, nil
		})

	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Radionomy with redirect",
			args: args{
				url: "http://listen.radionomy.com:80/Maartje",
			},
			want: true,
		},
		{
			name: "m3u",
			args: args{
				url: "http://icecast.com/stream.m3u",
			},
			want: true,
		},
		{
			name: "pls",
			args: args{
				url: "http://scserv.com/.pls",
			},
			want: true,
		},
		{
			name: "PDF FM",
			args: args{
				url: "http://pdfstream.com/;",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckValidPlaylist(tt.args.url); got != tt.want {
				t.Errorf("CheckValidPlaylist() = %v, want %v", got, tt.want)
			}
		})
	}
}
