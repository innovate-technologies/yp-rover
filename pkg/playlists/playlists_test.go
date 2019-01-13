package playlists

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetEntryURLs(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://scserv.com/.pls",
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(http.StatusOK, `
[playlist]
numberofentries=1
File1=https://opencast.radioca.st/streams/128kbps
Title1=OPENcast on DJ
Length1=-1
version=2
			`)
			res.Header.Set("Content-Type", "audio/x-scpls")
			return res, nil
		})
	httpmock.RegisterResponder("GET", "http://scserv.com/.m3u",
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(http.StatusOK, `
#EXTM3U

https://opencast.radioca.st/streams/128kbps
			`)
			res.Header.Set("Content-Type", "audio/x-mpegurl")
			return res, nil
		})
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test pls",
			args: args{
				url: "http://scserv.com/.pls",
			},
			want: []string{
				"https://opencast.radioca.st/streams/128kbps",
			},
		},
		{
			name: "Test m3u",
			args: args{
				url: "http://scserv.com/.pls",
			},
			want: []string{
				"https://opencast.radioca.st/streams/128kbps",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEntryURLs(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEntryURLs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEntryURLs() = %v, want %v", got, tt.want)
			}
		})
	}
}
