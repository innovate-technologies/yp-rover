package tunein

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/jarcoal/httpmock"
)

func TestAPI_BrowseStations(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// mock to list out the genres
	httpmock.RegisterResponder("GET", "https://opml.radiotime.com/Browse.ashx?filter=s&id=c57943&offset=0",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusOK, `
<?xml version="1.0" encoding="UTF-8"?>
<opml version="1">
	<head>
	<title>Top 40 &amp; Pop Music</title>
	<status>200</status>
	
	</head>
	<body>
<outline type="audio" text="ClubFm (Zele)" URL="http://opml.radiotime.com/Tune.ashx?id=s294876&amp;filter=s" bitrate="128" reliability="93" guide_id="s294876" subtext="De Oost-Vlaamse Regioradio. Voor Meer Info. vind je op www.radio" genre_id="g61" formats="mp3" item="station" image="http://cdn-profiles.tunein.com/s294876/images/logoq.png?t=152888" now_playing_id="s294876" preset_id="s294876"/>
<outline type="audio" text="Radio Beverland (Antwerp)" URL="http://opml.radiotime.com/Tune.ashx?id=s136890&amp;filter=s" bitrate="32" reliability="88" guide_id="s136890" subtext="English" genre_id="g61" formats="mp3" item="station" image="http://cdn-radiotime-logos.tunein.com/s136890q.png" now_playing_id="s136890" preset_id="s136890"/>
<outline type="audio" text="Radio FG Vlaanderen (Antwerp)" URL="http://opml.radiotime.com/Tune.ashx?id=s121313&amp;filter=s" bitrate="192" reliability="95" guide_id="s121313" subtext="F*cking Good Music Playlist" genre_id="g61" formats="mp3" show_id="p1178790" item="station" image="http://cdn-radiotime-logos.tunein.com/s121313q.png" current_track="F*cking Good Music Playlist" now_playing_id="s121313" preset_id="s121313"/>
<outline type="link" text="More Stations" URL="http://opml.radiotime.com/Browse.ashx?offset=26&amp;id=c57943&amp;filter=s" key="nextStations"/>
</body>
</opml>
			`), nil
		},
	)

	httpmock.RegisterResponder("GET", "https://opml.radiotime.com/Tune.ashx?id=s294876&filter=s",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusOK, "https://demo.radioca.st/test"), nil
		},
	)
	httpmock.RegisterResponder("GET", "https://opml.radiotime.com/Tune.ashx?id=s136890&filter=s",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusOK, "https://demo.radioca.st/test"), nil
		},
	)
	httpmock.RegisterResponder("GET", "https://opml.radiotime.com/Tune.ashx?id=s121313&filter=s",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusOK, "https://demo.radioca.st/test"), nil
		},
	)

	type args struct {
		guide  string
		offset int64
	}
	tests := []struct {
		name       string
		a          *API
		args       args
		want       []Station
		wantOffset int64
		wantErr    bool
	}{
		{
			name: "Test fetch stations",
			args: args{
				guide:  "c57943",
				offset: 0,
			},
			wantOffset: 26,
			want: []Station{
				Station{
					Name:                 "ClubFm (Zele)",
					Logo:                 "http://cdn-profiles.tunein.com/s294876/images/logoq.png?t=152888",
					MediaType:            "mp3",
					BitRate:              "128",
					GenreID:              "g61",
					Reliability:          "93",
					CurrentTrack:         "",
					CurrentTrackImageURL: "",
					TuneInURL:            "https://demo.radioca.st/test",
				},
				Station{
					Name:                 "Radio Beverland (Antwerp)",
					Logo:                 "http://cdn-radiotime-logos.tunein.com/s136890q.png",
					MediaType:            "mp3",
					BitRate:              "32",
					GenreID:              "g61",
					Reliability:          "88",
					CurrentTrack:         "",
					CurrentTrackImageURL: "",
					TuneInURL:            "https://demo.radioca.st/test",
				},
				Station{
					Name:                 "Radio FG Vlaanderen (Antwerp)",
					Logo:                 "http://cdn-radiotime-logos.tunein.com/s121313q.png",
					MediaType:            "mp3",
					BitRate:              "192",
					GenreID:              "g61",
					Reliability:          "95",
					CurrentTrack:         "",
					CurrentTrackImageURL: "",
					TuneInURL:            "https://demo.radioca.st/test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &API{}
			got, gotOffset, err := a.BrowseStations(tt.args.guide, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.BrowseStations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOffset != tt.wantOffset {
				t.Errorf("API.BrowseStations() offset = %v, want %v", gotOffset, tt.wantOffset)
			}
			if !reflect.DeepEqual(got, tt.want) {
				spew.Dump(got)
				t.Errorf("API.BrowseStations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_GetGenreGuides(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// mock to list out the genres
	httpmock.RegisterResponder("GET", "https://opml.radiotime.com/Browse.ashx?c=music",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusOK, `
			<?xml version="1.0" encoding="UTF-8"?>
			<opml version="1">
				<head>
				<title>Music</title>
				<status>200</status>
				
				</head>
				<body>
			<outline type="link" text="00&apos;s" URL="http://opml.radiotime.com/Browse.ashx?id=g2754" guide_id="g2754"/>
			<outline type="link" text="50&apos;s" URL="http://opml.radiotime.com/Browse.ashx?id=g390" guide_id="g390"/>
			<outline type="link" text="60&apos;s" URL="http://opml.radiotime.com/Browse.ashx?id=g407" guide_id="g407"/>
			<outline type="link" text="70&apos;s" URL="http://opml.radiotime.com/Browse.ashx?id=c100000783" guide_id="c100000783"/>
			<outline type="link" text="80&apos;s" URL="http://opml.radiotime.com/Browse.ashx?id=c100000781" guide_id="c100000781"/>
			<outline type="link" text="90&apos;s" URL="http://opml.radiotime.com/Browse.ashx?id=c100000946" guide_id="c100000946"/>
			<outline type="link" text="Adult Hits" URL="http://opml.radiotime.com/Browse.ashx?id=c57935" guide_id="c57935"/>
			<outline type="link" text="Blues Music" URL="http://opml.radiotime.com/Browse.ashx?id=c100001870" guide_id="c100001870"/>
			<outline type="link" text="Children&apos;s Music" URL="http://opml.radiotime.com/Browse.ashx?id=c530749" guide_id="c530749"/>
			<outline type="link" text="Classic Hits" URL="http://opml.radiotime.com/Browse.ashx?id=g2755" guide_id="g2755"/>
			<outline type="link" text="Classic Rock Music" URL="http://opml.radiotime.com/Browse.ashx?id=g54" guide_id="g54"/>
			<outline type="link" text="Classical Music" URL="http://opml.radiotime.com/Browse.ashx?id=c57939" guide_id="c57939"/>
			<outline type="link" text="College Radio" URL="http://opml.radiotime.com/Browse.ashx?id=g6" guide_id="g6"/>
			<outline type="link" text="Country Music" URL="http://opml.radiotime.com/Browse.ashx?id=c57940" guide_id="c57940"/>
			<outline type="link" text="Dance &amp; Electronic" URL="http://opml.radiotime.com/Browse.ashx?id=c57941" guide_id="c57941"/>
			<outline type="link" text="Easy Listening" URL="http://opml.radiotime.com/Browse.ashx?id=c10635888" guide_id="c10635888"/>
			<outline type="link" text="Folk Music" URL="http://opml.radiotime.com/Browse.ashx?id=g79" guide_id="g79"/>
			<outline type="link" text="Funk" URL="http://opml.radiotime.com/Browse.ashx?id=c100002460" guide_id="c100002460"/>
			<outline type="link" text="Hip Hop Music" URL="http://opml.radiotime.com/Browse.ashx?id=c57942" guide_id="c57942"/>
			<outline type="link" text="Holiday Music" URL="http://opml.radiotime.com/Browse.ashx?id=c100001750" guide_id="c100001750"/>
			<outline type="link" text="Indie Music" URL="http://opml.radiotime.com/Browse.ashx?id=c100000952" guide_id="c100000952"/>
			<outline type="link" text="Jazz Music" URL="http://opml.radiotime.com/Browse.ashx?id=c57944" guide_id="c57944"/>
			<outline type="link" text="Latin Music" URL="http://opml.radiotime.com/Browse.ashx?id=c100002533" guide_id="c100002533"/>
			<outline type="link" text="Music Podcasts" URL="http://opml.radiotime.com/Browse.ashx?id=c100000086" guide_id="c100000086"/>
			<outline type="link" text="R&amp;B Music" URL="http://opml.radiotime.com/Browse.ashx?id=g4152" guide_id="g4152"/>
			<outline type="link" text="Reggae Music" URL="http://opml.radiotime.com/Browse.ashx?id=g85" guide_id="g85"/>
			<outline type="link" text="Religious Music" URL="http://opml.radiotime.com/Browse.ashx?id=c57950" guide_id="c57950"/>
			<outline type="link" text="Rock Music" URL="http://opml.radiotime.com/Browse.ashx?id=c57951" guide_id="c57951"/>
			<outline type="link" text="Soul" URL="http://opml.radiotime.com/Browse.ashx?id=c1367173" guide_id="c1367173"/>
			<outline type="link" text="Top 40 &amp; Pop Music" URL="http://opml.radiotime.com/Browse.ashx?id=c57943" guide_id="c57943"/>
			<outline type="link" text="World Music" URL="http://opml.radiotime.com/Browse.ashx?id=g22" guide_id="g22"/>
				</body>
			</opml>
			`), nil
		},
	)
	tests := []struct {
		name    string
		a       *API
		want    map[string]string
		wantErr bool
	}{
		{
			name:    "Test genres read",
			wantErr: false,
			want: map[string]string{
				"50's":               "g390",
				"Classic Rock Music": "g54",
				"Music Podcasts":     "c100000086",
				"Folk Music":         "g79",
				"Religious Music":    "c57950",
				"Top 40 & Pop Music": "c57943",
				"90's":               "c100000946",
				"Jazz Music":         "c57944",
				"Blues Music":        "c100001870",
				"Country Music":      "c57940",
				"Indie Music":        "c100000952",
				"Rock Music":         "c57951",
				"Classic Hits":       "g2755",
				"Easy Listening":     "c10635888",
				"Funk":               "c100002460",
				"R&B Music":          "g4152",
				"Reggae Music":       "g85",
				"Soul":               "c1367173",
				"Dance & Electronic": "c57941",
				"Holiday Music":      "c100001750",
				"60's":               "g407",
				"70's":               "c100000783",
				"Adult Hits":         "c57935",
				"Hip Hop Music":      "c57942",
				"Latin Music":        "c100002533",
				"World Music":        "g22",
				"00's":               "g2754",
				"80's":               "c100000781",
				"Children's Music":   "c530749",
				"Classical Music":    "c57939",
				"College Radio":      "g6",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &API{}
			got, err := a.GetGenreGuides()
			if (err != nil) != tt.wantErr {
				t.Errorf("API.GetGenreGuides() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("API.GetGenreGuides() = %v, want %v", got, tt.want)
			}
		})
	}
}
