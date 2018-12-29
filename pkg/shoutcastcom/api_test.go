package shoutcastcom

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/jarcoal/httpmock"
)

func TestSHOUTcastAPIClient_GetAllGenres(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// mock to list out the genres
	httpmock.RegisterResponder("GET", "https://api.shoutcast.com/legacy/genrelist?k=test",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusOK, `
			<?xml version="1.0" encoding="UTF-8" standalone="yes" ?>
			<genrelist><genre name="00s" count="28" /><genre name="30s" count="4" /><genre name="40s" count="1" /><genre name="50s" count="16" /><genre name="60s" count="512" /><genre name="70s"count="454" /><genre name="80s" count="1187" /><genre name="90s" count="529" /><genre name="Acid House" count="6" /><genre name="Acid Jazz" count="85" /><genre name="Acoustic Blues" count="3689" /><genre name="Adult" count="14" /><genre name="Adult Album Alternative" count="8" /><genre name="Adult Alternative" count="34" /><genre name="Adult Contemporary" count="146" /><genre name="African" count="446" /><genre name="Afrikaans" count="10" /><genre name="Alt Country" count="16" /><genre name="Alternative" count="141" /><genre name="AlternativeFolk" count="14" /><genre name="Alternative Rap" count="1" /><genre name="Ambient" count="1538" /><genre name="Americana" count="19" /><genre name="Anime" count="12" /><genre name="Anniversary" count="0" /><genre name="Arabic" count="370" /><genre name="Asian" count="422" /><genre name="Avant Garde" count="1" /><genre name="Bachata" count="5" /><genre name="Banda" count="5" /><genre name="Barbershop" count="1" /><genre name="Baroque" count="201" /><genre name="Best Of" count="10" /><genre name="Big Band" count="2" /><genre name="Big Beat" count="0" /><genre name="Birthday" count="0" /><genre name="Black Metal" count="169" /><genre name="BlogTalk" count="5" /><genre name="Bluegrass" count="115" /><genre name="Blues" count="5469" /><genre name="Bollywood" count="11" /><genre name="Bop" count="23" /><genre name="Bossa Nova" count="26" /><genre name="Brazilian" count="361" /><genre name="Breakbeat" count="342" /><genre name="British Invasion" count="0" /><genre name="Britpop" count="288" /><genre name="Bubblegum Pop" count="0" /><genre name="Cajun and Zydeco" count="12" /><genre name="Caribbean" count="16" /><genre name="Celtic" count="55" /><genre name="Celtic Rock" count="34" /><genre name="Chamber" count="50" /><genre name="Chicago Blues" count="1" /><genre name="Chill" count="1556" /><genre name="Chinese" count="0" /><genre name="Choral" count="1" /><genre name="Christian" count="1125" /><genre name="Christian Metal" count="1" /><genre name="Christian Rap" count="2" /><genre name="Christian Rock" count="213" /><genre name="Christmas" count="176" /><genre name="Classic Alternative" count="9" /><genre name="Classic Christian" count="7" /><genre name="Classic Country" count="26" /><genre name="Classic Jazz" count="127" /><genre name="Classic Metal" count="3" /><genre name="Classic R&amp;B" count="16" /><genre name="Classic Rock" count="1314" /><genre name="Classical" count="1379" /><genre name="Classical Period" count="6" /><genre name="College" count="64" /><genre name="Comedy" count="262" /><genre name="Community" count="45" /><genre name="Contemporary Bluegrass" count="0" /><genre name="Contemporary Blues" count="4" /><genre name="Contemporary Country" count="11" /><genre name="Contemporary Folk" count="3" /><genre name="Contemporary Gospel" count="10" /><genre name="Contemporary R&amp;B" count="5" /><genre name="Contemporary Reggae" count="10" /><genre name="Cool Jazz" count="6" /><genre name="Country" count="2092" /><genre name="Country Blues" count="873" /><genre name="Creole" count="57" /><genre name="Cumbia" count="125" /><genre name="Dance" count="2962" /><genre name="Dance Pop" count="82" /><genre name="Dancehall" count="350" /><genre name="Dancepunk" count="14" /><genre name="Death Metal" count="74" /><genre name="Decades" count="89" /><genre name="Delta Blues" count="302" /><genre name="Demo" count="26" /><genre name="Dirty South" count="65" /><genre name="Disco" count="478" /><genre name="Doo Wop" count="1" /><genre name="Downtempo" count="172" /><genre name="Dream Pop" count="4" /><genre name="Drum and Bass" count="223" /><genre name="Dub" count="56" /><genre name="Dubstep" count="637" /><genre name="Early Classical" count="2" /><genre name="East Coast Rap" count="1" /><genre name="Easy Listening" count="1369" /><genre name="Eclectic" count="33" /><genre name="Educational" count="376" /><genre name="Electric Blues" count="886" /><genre name="Electro" count="2048" /><genre name="Electronic" count="2899"/><genre name="Emo" count="37" /><genre name="Environmental" count="0" /><genre name="Ethnic Fusion" count="35" /><genre name="European" count="104" /><genre name="Exotica" count="2"/><genre name="Experimental" count="289" /><genre name="Extreme Metal" count="2" /><genre name="Female" count="0" /><genre name="Filipino" count="8" /><genre name="Flamenco" count="67" /><genre name="Folk" count="769" /><genre name="Folk Rock" count="7" /><genre name="Freestyle" count="165" /><genre name="French" count="6" /><genre name="Funk" count="406" /><genre name="Fusion" count="21" /><genre name="Gangsta Rap" count="380" /><genre name="Garage" count="145" /><genre name="Garage Rock" count="68" /><genre name="German" count="18" /><genrename="Glam" count="41" /><genre name="Gospel" count="994" /><genre name="Goth" count="100" /><genre name="Government" count="3" /><genre name="Greek" count="113" /><genre name="Grindcore" count="0" /><genre name="Grunge" count="91" /><genre name="Hair Metal" count="0" /><genre name="Halloween" count="4" /><genre name="Hanukkah" count="0" /><genre name="Hard Bop" count="0" /><genre name="Hard House" count="4" /><genre name="Hard Rock" count="456" /><genre name="Hardcore" count="318" /><genre name="Hawaiian and Pacific" count="15" /><genre name="Healing" count="5" /><genre name="Heartache" count="0" /><genre name="Heavy Metal" count="393" /><genre name="Hebrew" count="23" /><genre name="Hindi" count="9" /><genre name="Hip Hop" count="2020" /><genre name="Honeymoon" count="0" /><genre name="Honky Tonk" count="37" /><genre name="Hot Country Hits" count="116" /><genre name="House" count="1777" /><genre name="IDM" count="3" /><genre name="Idols" count="2" /><genre name="Impressionist" count="1" /><genre name="Indian" count="18" /><genre name="Indie Pop" count="421" /><genre name="Indie Rock" count="238" /><genre name="Industrial" count="147" /><genre name="Inspirational" count="448" /><genre name="Instrumental" count="49" /><genre name="International" count="885" /><genre name="Islamic" count="20" /><genre name="Jam Bands" count="4" /><genre name="Japanese" count="236" /><genre name="Jazz" count="710" /><genre name="JPOP" count="21" /><genre name="JROCK" count="4" /><genre name="Jungle" count="25" /><genre name="Kids" count="188" /><genre name="Klezmer" count="0" /><genre name="Korean" count="2" /><genre name="KPOP" count="26" /><genre name="Kwanzaa" count="0" /><genre name="Latin" count="1578" /><genre name="Latin Dance" count="194" /><genre name="Latin Jazz" count="21" /><genre name="Latin Pop" count="729" /><genre name="Latin Rap and Hip Hop" count="6" /><genre name="Latin Rock" count="196" /><genre name="LGBT" count="3" /><genre name="Light Rock" count="28" /><genre name="LoFi" count="38" /><genre name="Lounge" count="388" /><genre name="Love and Romance" count="18" /><genre name="Mariachi" count="54" /><genre name="Meditation" count="4" /><genre name="Mediterranean" count="1" /><genre name="Merengue" count="196" /><genre name="Metal" count="615" /><genre name="Metalcore" count="38" /><genre name="Middle Eastern" count="61" /><genre name="Misc" count="539" /><genre name="Mixtapes" count="11" /><genre name="Modern" count="9" /><genre name="Modern Rock" count="10" /><genre name="Motown" count="56" /><genre name="Neo Soul" count="56" /><genre name="New Acoustic" count="1" /><genre name="New Age" count="460" /><genre name="New Wave" count="131" /><genre name="News" count="1478" /><genre name="Noise Pop" count="8" /><genre name="North American" count="1" /><genre name="Old School" count="3" /><genre name="Old Time" count="12" /><genre name="Old Time Radio" count="34" /><genre name="Oldies" count="970" /><genre name="Opera" count="113" /><genre name="Orchestral Pop" count="2" /><genre name="Original Score" count="9" /><genre name="Other Talk" count="343" /><genre name="Party Mix" count="10" /><genre name="Patriotic" count="2" /><genre name="Piano" count="2" /><genre name="Piano Rock" count="2" /><genre name="Political" count="93" /><genre name="Polka" count="47" /><genre name="Pop" count="6528" /><genre name="Pop Reggae" count="80" /><genre name="Post Punk" count="16" /><genre name="Power Metal" count="2" /><genre name="Power Pop" count="38" /><genre name="Praise and Worship" count="497" /><genre name="Prog Rock" count="154" /><genre name="Progressive" count="17" /><genre name="Progressive Metal" count="32" /><genre name="Psychedelic" count="83" /><genre name="Public Radio" count="77" /><genre name="Punk" count="175" /><genre name="Quiet Storm" count="3" /><genre name="R&amp;B and Urban" count="595" /><genre name="Ragga" count="18" /><genre name="Rainy Day Mix" count="0" /><genre name="Ranchera" count="3" /><genre name="Rap" count="766" /><genre name="Rap Metal" count="0" /><genre name="Reality" count="0" /><genre name="Reggae" count="449" /><genre name="Reggae Roots" count="63" /><genre name="Reggaeton" count="434" /><genre name="Regional Mexican" count="20" /><genre name="Rock" count="1777" /><genre name="Rock &amp; Roll" count="876" /><genre name="Rock Steady" count="1" /><genre name="Rockabilly" count="135" /><genre name="Romantic" count="223" /><genre name="Russian" count="12" /><genre name="Salsa" count="415" /><genre name="Samba" count="37" /><genre name="Scanner" count="9" /><genre name="Seasonal and Holiday" count="47" /><genre name="Sermons and Services" count="16" /><genre name="Sexy" count="2" /><genre name="Showtunes" count="1" /><genre name="Shuffle" count="1" /><genre name="Singer and Songwriter" count="4" /><genre name="Ska" count="68" /><genre name="Smooth Jazz" count="241" /><genre name="Soca" count="11" /><genre name="Soft Rock" count="712" /><genre name="Soul" count="268" /><genre name="Soundtracks" count="245" /><genre name="South American" count="5" /><genre name="Southern Gospel" count="7" /><genre name="Space Age Pop" count="5" /><genre name="Spiritual" count="1506" /><genre name="Spoken Word" count="13" /><genre name="Sports" count="298" /><genre name="Surf" count="2" /><genre name="Swing" count="49" /><genre name="Symphony" count="70" /><genre name="Talk" count="1812" /><genre name="Tamil" count="37" /><genre name="Tango" count="44" /><genre name="Techno" count="670" /><genre name="Technology" count="7" /><genre name="Teen Pop" count="1" /><genre name="Tejano" count="10" /><genre name="Themes" count="9" /><genre name="Thrash Metal" count="45" /><genre name="Top 40" count="4575" /><genre name="Traditional Folk" count="15" /><genre name="Traditional Gospel" count="5" /><genre name="Trance" count="527" /><genre name="Travel Mix" count="1" /><genre name="Tribal" count="22" /><genre name="Tribute" count="2" /><genre name="Trip Hop" count="201" /><genre name="Trippy" count="2" /><genre name="Tropicalia" count="16" /><genre name="Turkish" count="20" /><genre name="Turntablism" count="0" /><genre name="Underground Hip Hop" count="9" /><genre name="Urban Contemporary" count="207" /><genre name="Valentine" count="0" /><genre name="Video Game Music" count="36" /><genre name="Vocal Jazz" count="27" /><genre name="Weather" count="34" /><genre name="Wedding" count="0" /><genre name="West Coast Rap" count="1" /><genre name="Western" count="5" /><genre name="Winter" count="1" /><genre name="Work Mix" count="11" /><genre name="World Folk" count="69" /><genre name="World Fusion" count="3" /><genre name="World Pop" count="78" /><genre name="Worldbeat" count="128" /><genre name="Xtreme" count="1" /><genre name="Zouk" count="145" /></genrelist>
			`), nil
		},
	)
	httpmock.RegisterResponder("GET", "https://api.shoutcast.com/legacy/genrelist?k=emptytest",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusOK, `
			<?xml version="1.0" encoding="UTF-8" standalone="yes" ?>
			`), nil
		},
	)
	httpmock.RegisterResponder("GET", "https://api.shoutcast.com/legacy/genrelist?k=500",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusInternalServerError, `
			Java Server Down
			`), nil
		},
	)
	httpmock.RegisterResponder("GET", "https://api.shoutcast.com/legacy/genrelist?k=500",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusInternalServerError, `
			Java Server Down
			`), nil
		},
	)
	httpmock.RegisterResponder("GET", "https://api.shoutcast.com/legacy/genrelist?k=inetdown",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusInternalServerError, ""), errors.New("Internet Down")
		},
	)

	type fields struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			name: "Test HTTP 500 error",
			fields: fields{
				key: "500",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test HTTP request error",
			fields: fields{
				key: "inetdown",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test empty XML response",
			fields: fields{
				key: "emptytest",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test OK response",
			fields: fields{
				key: "test",
			},
			want: []string{
				"00s",
				"30s",
				"40s",
				"50s",
				"60s",
				"70s",
				"80s",
				"90s",
				"Acid House",
				"Acid Jazz",
				"Acoustic Blues",
				"Adult",
				"Adult Album Alternative",
				"Adult Alternative",
				"Adult Contemporary",
				"African",
				"Afrikaans",
				"Alt Country",
				"Alternative",
				"AlternativeFolk",
				"Alternative Rap",
				"Ambient",
				"Americana",
				"Anime",
				"Anniversary",
				"Arabic",
				"Asian",
				"Avant Garde",
				"Bachata",
				"Banda",
				"Barbershop",
				"Baroque",
				"Best Of",
				"Big Band",
				"Big Beat",
				"Birthday",
				"Black Metal",
				"BlogTalk",
				"Bluegrass",
				"Blues",
				"Bollywood",
				"Bop",
				"Bossa Nova",
				"Brazilian",
				"Breakbeat",
				"British Invasion",
				"Britpop",
				"Bubblegum Pop",
				"Cajun and Zydeco",
				"Caribbean",
				"Celtic",
				"Celtic Rock",
				"Chamber",
				"Chicago Blues",
				"Chill",
				"Chinese",
				"Choral",
				"Christian",
				"Christian Metal",
				"Christian Rap",
				"Christian Rock",
				"Christmas",
				"Classic Alternative",
				"Classic Christian",
				"Classic Country",
				"Classic Jazz",
				"Classic Metal",
				"Classic R&B",
				"Classic Rock",
				"Classical",
				"Classical Period",
				"College",
				"Comedy",
				"Community",
				"Contemporary Bluegrass",
				"Contemporary Blues",
				"Contemporary Country",
				"Contemporary Folk",
				"Contemporary Gospel",
				"Contemporary R&B",
				"Contemporary Reggae",
				"Cool Jazz",
				"Country",
				"Country Blues",
				"Creole",
				"Cumbia",
				"Dance",
				"Dance Pop",
				"Dancehall",
				"Dancepunk",
				"Death Metal",
				"Decades",
				"Delta Blues",
				"Demo",
				"Dirty South",
				"Disco",
				"Doo Wop",
				"Downtempo",
				"Dream Pop",
				"Drum and Bass",
				"Dub",
				"Dubstep",
				"Early Classical",
				"East Coast Rap",
				"Easy Listening",
				"Eclectic",
				"Educational",
				"Electric Blues",
				"Electro",
				"Electronic",
				"Emo",
				"Environmental",
				"Ethnic Fusion",
				"European",
				"Exotica",
				"Experimental",
				"Extreme Metal",
				"Female",
				"Filipino",
				"Flamenco",
				"Folk",
				"Folk Rock",
				"Freestyle",
				"French",
				"Funk",
				"Fusion",
				"Gangsta Rap",
				"Garage",
				"Garage Rock",
				"German",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SHOUTcastAPIClient{
				key: tt.fields.key,
			}
			got, err := s.GetAllGenres()
			if (err != nil) != tt.wantErr {
				t.Errorf("SHOUTcastAPIClient.GetAllGenres() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SHOUTcastAPIClient.GetAllGenres() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHOUTcastAPIClient_GetTop500(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.shoutcast.com/legacy/Top500?br=&k=top10&limit=10&mt=",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusOK, `
			<?xml version="1.0" encoding="UTF-8"?>
<stationlist><tunein base="/sbin/tunein-station.pls" base-m3u="/sbin/tunein-station.m3u" base-xspf="/sbin/tunein-station.xspf" /><station name="Hitradio OE3" mt="audio/mpeg" id="1826116" br="192" genre="00s" genre2="80s" genre3="90s" logo="https://somesite.com/radios/200/6/6f2c/6f2ca76b-4d1d-464d-a336-db0913fb5229.jpg" ct="Nathan Trent - Good Vibes" lc="15568" /><station name="Radio 2.0 - Valli di Bergamo" mt="audio/mpeg" id="1738454" br="192" genre="Pop" genre2="Rock" genre3="80s" genre4="70s" genre5="Top 40" logo="http://somesite.com/document/radios/7/7ae5/7ae585ab-4287-44fe-a39d-d8e29eda7ff2.png" ct="Fiorello - Spiagge" lc="6976" /><station name="ANTENA1 | 94,7 FM" mt="audio/aacp" id="1796249" br="64" genre="Pop" genre2="Rock" genre3="Top 40" logo="http://somesite.com/document/radios/9/9e64/9e6458c0-3a81-427f-b8cf-e3ab5721fbf9.png" lc="6393" /><station name="Enjoy Hits!" mt="audio/mpeg" id="1444774" br="192" genre="Latin" lc="5552" /><station name="FastDance.FM | Electronic Music Radio" mt="audio/aacp" id="1736269" br="32" genre="Dance" genre2="Easy Listening" genre3="Electronic" genre4="Techno" genre5="Top 40" logo="http://somesite.com/document/radios/0/0eb4/0eb472e2-c8c9-49d4-aded-bdc6570b6ea7.png" lc="4053" /><station name="NobodyLovesMe.de" mt="audio/aacp" id="1704212" br="32" genre="Easy Listening" genre2="Electro" genre3="Top 40" genre4="Jazz" genre5="Educational" logo="http://somesite.com/document/radios/5/5aca/5acaae4c-5086-430e-83e7-f59048dbff2c.png" lc="3922" /><station name="UpDance.de // Fresh MIX!" mt="audio/aacp" id="1829597" br="32" genre="Top 40" genre2="News" genre3="80s" genre4="90s" genre5="Oldies" logo="http://somesite.com/document/radios/e/e296/e2969202-27af-44e4-b44e-89761239b37c.png" lc="3823" /><station name="Radio 2.0 - Valli di Bergamo HD" mt="audio/mpeg" id="1271585" br="760" genre="Pop" genre2="Rock" genre3="80s" genre4="70s" genre5="Top 40" logo="http://somesite.com/document/radios/d/d9cc/d9cc0fcf-2fa3-4a3b-8aa1-8681d68bc229.png" ct="Fiorello - Spiagge" lc="3770" /><station name="RockRadio.cf" mt="audio/aacp" id="1653990" br="32" genre="Country" genre2="Hard Rock" genre3="Rock" genre4="Rockabilly" genre5="Rock &amp; Roll" logo="http://somesite.com/document/radios/f/f41c/f41c19bf-625f-4a3b-9142-95b8358d86a1.png" lc="3392" /><station name="Radio FM4" mt="audio/mpeg" id="1783457" br="160" genre="Misc" genre2="Electronic" genre3="Rock" genre4="Pop" genre5="Hip Hop" logo="http://somesite.com/document/radios/1/1bb3/1bb3b2af-27dc-4b1f-88f9-c1c6655d355d.jpg" ct="Best of FM4 Acoustic Sessions: Teil 4" lc="3332" /></stationlist>
			`), nil
		},
	)

	type fields struct {
		key string
	}
	type args struct {
		mediaType string
		bitRate   string
		limit     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Station
		wantErr bool
	}{
		{
			name: "Get Top 10",
			fields: fields{
				key: "top10",
			},
			args: args{
				limit: "10",
			},
			want: []Station{
				Station{
					ID:           "1826116",
					Name:         "Hitradio OE3",
					MediaType:    "audio/mpeg",
					BitRate:      "192",
					Genre:        "00s",
					Genre2:       "80s",
					Genre3:       "90s",
					Genre4:       "",
					Genre5:       "",
					LogoURL:      "https://somesite.com/radios/200/6/6f2c/6f2ca76b-4d1d-464d-a336-db0913fb5229.jpg",
					CurrentTrack: "Nathan Trent - Good Vibes",
					ListnerCount: 15568,
				},
				Station{
					ID:           "1738454",
					Name:         "Radio 2.0 - Valli di Bergamo",
					MediaType:    "audio/mpeg",
					BitRate:      "192",
					Genre:        "Pop",
					Genre2:       "Rock",
					Genre3:       "80s",
					Genre4:       "70s",
					Genre5:       "Top 40",
					LogoURL:      "http://somesite.com/document/radios/7/7ae5/7ae585ab-4287-44fe-a39d-d8e29eda7ff2.png",
					CurrentTrack: "Fiorello - Spiagge",
					ListnerCount: 6976,
				},
				Station{
					ID:           "1796249",
					Name:         "ANTENA1 | 94,7 FM",
					MediaType:    "audio/aacp",
					BitRate:      "64",
					Genre:        "Pop",
					Genre2:       "Rock",
					Genre3:       "Top 40",
					Genre4:       "",
					Genre5:       "",
					LogoURL:      "http://somesite.com/document/radios/9/9e64/9e6458c0-3a81-427f-b8cf-e3ab5721fbf9.png",
					CurrentTrack: "",
					ListnerCount: 6393,
				},
				Station{
					ID:           "1444774",
					Name:         "Enjoy Hits!",
					MediaType:    "audio/mpeg",
					BitRate:      "192",
					Genre:        "Latin",
					Genre2:       "",
					Genre3:       "",
					Genre4:       "",
					Genre5:       "",
					LogoURL:      "",
					CurrentTrack: "",
					ListnerCount: 5552,
				},
				Station{
					ID:           "1736269",
					Name:         "FastDance.FM | Electronic Music Radio",
					MediaType:    "audio/aacp",
					BitRate:      "32",
					Genre:        "Dance",
					Genre2:       "Easy Listening",
					Genre3:       "Electronic",
					Genre4:       "Techno",
					Genre5:       "Top 40",
					LogoURL:      "http://somesite.com/document/radios/0/0eb4/0eb472e2-c8c9-49d4-aded-bdc6570b6ea7.png",
					CurrentTrack: "",
					ListnerCount: 4053,
				},
				Station{
					ID:           "1704212",
					Name:         "NobodyLovesMe.de",
					MediaType:    "audio/aacp",
					BitRate:      "32",
					Genre:        "Easy Listening",
					Genre2:       "Electro",
					Genre3:       "Top 40",
					Genre4:       "Jazz",
					Genre5:       "Educational",
					LogoURL:      "http://somesite.com/document/radios/5/5aca/5acaae4c-5086-430e-83e7-f59048dbff2c.png",
					CurrentTrack: "",
					ListnerCount: 3922,
				},
				Station{
					ID:           "1829597",
					Name:         "UpDance.de // Fresh MIX!",
					MediaType:    "audio/aacp",
					BitRate:      "32",
					Genre:        "Top 40",
					Genre2:       "News",
					Genre3:       "80s",
					Genre4:       "90s",
					Genre5:       "Oldies",
					LogoURL:      "http://somesite.com/document/radios/e/e296/e2969202-27af-44e4-b44e-89761239b37c.png",
					CurrentTrack: "",
					ListnerCount: 3823,
				},
				Station{
					ID:           "1271585",
					Name:         "Radio 2.0 - Valli di Bergamo HD",
					MediaType:    "audio/mpeg",
					BitRate:      "760",
					Genre:        "Pop",
					Genre2:       "Rock",
					Genre3:       "80s",
					Genre4:       "70s",
					Genre5:       "Top 40",
					LogoURL:      "http://somesite.com/document/radios/d/d9cc/d9cc0fcf-2fa3-4a3b-8aa1-8681d68bc229.png",
					CurrentTrack: "Fiorello - Spiagge",
					ListnerCount: 3770,
				},
				Station{
					ID:           "1653990",
					Name:         "RockRadio.cf",
					MediaType:    "audio/aacp",
					BitRate:      "32",
					Genre:        "Country",
					Genre2:       "Hard Rock",
					Genre3:       "Rock",
					Genre4:       "Rockabilly",
					Genre5:       "Rock & Roll",
					LogoURL:      "http://somesite.com/document/radios/f/f41c/f41c19bf-625f-4a3b-9142-95b8358d86a1.png",
					CurrentTrack: "",
					ListnerCount: 3392,
				},
				Station{
					ID:           "1783457",
					Name:         "Radio FM4",
					MediaType:    "audio/mpeg",
					BitRate:      "160",
					Genre:        "Misc",
					Genre2:       "Electronic",
					Genre3:       "Rock",
					Genre4:       "Pop",
					Genre5:       "Hip Hop",
					LogoURL:      "http://somesite.com/document/radios/1/1bb3/1bb3b2af-27dc-4b1f-88f9-c1c6655d355d.jpg",
					CurrentTrack: "Best of FM4 Acoustic Sessions: Teil 4",
					ListnerCount: 3332,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SHOUTcastAPIClient{
				key: tt.fields.key,
			}
			got, err := s.GetTop500(tt.args.mediaType, tt.args.bitRate, tt.args.limit)
			if (err != nil) != tt.wantErr {
				spew.Dump(err)
				t.Errorf("SHOUTcastAPIClient.GetTop500() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SHOUTcastAPIClient.GetTop500() = %v, want %v", got, tt.want)
				spew.Dump(got)
			}
		})
	}
}

func TestSHOUTcastAPIClient_GetByGenre(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.shoutcast.com/legacy/genresearch?br=&genre=test&k=test&limit=&mt=",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusOK, `
			<?xml version="1.0" encoding="UTF-8"?>
<stationlist><tunein base="/sbin/tunein-station.pls" base-m3u="/sbin/tunein-station.m3u" base-xspf="/sbin/tunein-station.xspf" /><station name="Hitradio OE3" mt="audio/mpeg" id="1826116" br="192" genre="00s" genre2="80s" genre3="90s" logo="https://somesite.com/radios/200/6/6f2c/6f2ca76b-4d1d-464d-a336-db0913fb5229.jpg" ct="Nathan Trent - Good Vibes" lc="15568" /><station name="Radio 2.0 - Valli di Bergamo" mt="audio/mpeg" id="1738454" br="192" genre="Pop" genre2="Rock" genre3="80s" genre4="70s" genre5="Top 40" logo="http://somesite.com/document/radios/7/7ae5/7ae585ab-4287-44fe-a39d-d8e29eda7ff2.png" ct="Fiorello - Spiagge" lc="6976" /><station name="ANTENA1 | 94,7 FM" mt="audio/aacp" id="1796249" br="64" genre="Pop" genre2="Rock" genre3="Top 40" logo="http://somesite.com/document/radios/9/9e64/9e6458c0-3a81-427f-b8cf-e3ab5721fbf9.png" lc="6393" /><station name="Enjoy Hits!" mt="audio/mpeg" id="1444774" br="192" genre="Latin" lc="5552" /><station name="FastDance.FM | Electronic Music Radio" mt="audio/aacp" id="1736269" br="32" genre="Dance" genre2="Easy Listening" genre3="Electronic" genre4="Techno" genre5="Top 40" logo="http://somesite.com/document/radios/0/0eb4/0eb472e2-c8c9-49d4-aded-bdc6570b6ea7.png" lc="4053" /><station name="NobodyLovesMe.de" mt="audio/aacp" id="1704212" br="32" genre="Easy Listening" genre2="Electro" genre3="Top 40" genre4="Jazz" genre5="Educational" logo="http://somesite.com/document/radios/5/5aca/5acaae4c-5086-430e-83e7-f59048dbff2c.png" lc="3922" /><station name="UpDance.de // Fresh MIX!" mt="audio/aacp" id="1829597" br="32" genre="Top 40" genre2="News" genre3="80s" genre4="90s" genre5="Oldies" logo="http://somesite.com/document/radios/e/e296/e2969202-27af-44e4-b44e-89761239b37c.png" lc="3823" /><station name="Radio 2.0 - Valli di Bergamo HD" mt="audio/mpeg" id="1271585" br="760" genre="Pop" genre2="Rock" genre3="80s" genre4="70s" genre5="Top 40" logo="http://somesite.com/document/radios/d/d9cc/d9cc0fcf-2fa3-4a3b-8aa1-8681d68bc229.png" ct="Fiorello - Spiagge" lc="3770" /><station name="RockRadio.cf" mt="audio/aacp" id="1653990" br="32" genre="Country" genre2="Hard Rock" genre3="Rock" genre4="Rockabilly" genre5="Rock &amp; Roll" logo="http://somesite.com/document/radios/f/f41c/f41c19bf-625f-4a3b-9142-95b8358d86a1.png" lc="3392" /><station name="Radio FM4" mt="audio/mpeg" id="1783457" br="160" genre="Misc" genre2="Electronic" genre3="Rock" genre4="Pop" genre5="Hip Hop" logo="http://somesite.com/document/radios/1/1bb3/1bb3b2af-27dc-4b1f-88f9-c1c6655d355d.jpg" ct="Best of FM4 Acoustic Sessions: Teil 4" lc="3332" /></stationlist>
			`), nil
		},
	)

	type fields struct {
		key string
	}
	type args struct {
		genre     string
		mediaType string
		bitRate   string
		limit     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Station
		wantErr bool
	}{
		{
			name: "Get 10 entries",
			fields: fields{
				key: "test",
			},
			args: args{
				genre: "test",
			},
			want: []Station{
				Station{
					ID:           "1826116",
					Name:         "Hitradio OE3",
					MediaType:    "audio/mpeg",
					BitRate:      "192",
					Genre:        "00s",
					Genre2:       "80s",
					Genre3:       "90s",
					Genre4:       "",
					Genre5:       "",
					LogoURL:      "https://somesite.com/radios/200/6/6f2c/6f2ca76b-4d1d-464d-a336-db0913fb5229.jpg",
					CurrentTrack: "Nathan Trent - Good Vibes",
					ListnerCount: 15568,
				},
				Station{
					ID:           "1738454",
					Name:         "Radio 2.0 - Valli di Bergamo",
					MediaType:    "audio/mpeg",
					BitRate:      "192",
					Genre:        "Pop",
					Genre2:       "Rock",
					Genre3:       "80s",
					Genre4:       "70s",
					Genre5:       "Top 40",
					LogoURL:      "http://somesite.com/document/radios/7/7ae5/7ae585ab-4287-44fe-a39d-d8e29eda7ff2.png",
					CurrentTrack: "Fiorello - Spiagge",
					ListnerCount: 6976,
				},
				Station{
					ID:           "1796249",
					Name:         "ANTENA1 | 94,7 FM",
					MediaType:    "audio/aacp",
					BitRate:      "64",
					Genre:        "Pop",
					Genre2:       "Rock",
					Genre3:       "Top 40",
					Genre4:       "",
					Genre5:       "",
					LogoURL:      "http://somesite.com/document/radios/9/9e64/9e6458c0-3a81-427f-b8cf-e3ab5721fbf9.png",
					CurrentTrack: "",
					ListnerCount: 6393,
				},
				Station{
					ID:           "1444774",
					Name:         "Enjoy Hits!",
					MediaType:    "audio/mpeg",
					BitRate:      "192",
					Genre:        "Latin",
					Genre2:       "",
					Genre3:       "",
					Genre4:       "",
					Genre5:       "",
					LogoURL:      "",
					CurrentTrack: "",
					ListnerCount: 5552,
				},
				Station{
					ID:           "1736269",
					Name:         "FastDance.FM | Electronic Music Radio",
					MediaType:    "audio/aacp",
					BitRate:      "32",
					Genre:        "Dance",
					Genre2:       "Easy Listening",
					Genre3:       "Electronic",
					Genre4:       "Techno",
					Genre5:       "Top 40",
					LogoURL:      "http://somesite.com/document/radios/0/0eb4/0eb472e2-c8c9-49d4-aded-bdc6570b6ea7.png",
					CurrentTrack: "",
					ListnerCount: 4053,
				},
				Station{
					ID:           "1704212",
					Name:         "NobodyLovesMe.de",
					MediaType:    "audio/aacp",
					BitRate:      "32",
					Genre:        "Easy Listening",
					Genre2:       "Electro",
					Genre3:       "Top 40",
					Genre4:       "Jazz",
					Genre5:       "Educational",
					LogoURL:      "http://somesite.com/document/radios/5/5aca/5acaae4c-5086-430e-83e7-f59048dbff2c.png",
					CurrentTrack: "",
					ListnerCount: 3922,
				},
				Station{
					ID:           "1829597",
					Name:         "UpDance.de // Fresh MIX!",
					MediaType:    "audio/aacp",
					BitRate:      "32",
					Genre:        "Top 40",
					Genre2:       "News",
					Genre3:       "80s",
					Genre4:       "90s",
					Genre5:       "Oldies",
					LogoURL:      "http://somesite.com/document/radios/e/e296/e2969202-27af-44e4-b44e-89761239b37c.png",
					CurrentTrack: "",
					ListnerCount: 3823,
				},
				Station{
					ID:           "1271585",
					Name:         "Radio 2.0 - Valli di Bergamo HD",
					MediaType:    "audio/mpeg",
					BitRate:      "760",
					Genre:        "Pop",
					Genre2:       "Rock",
					Genre3:       "80s",
					Genre4:       "70s",
					Genre5:       "Top 40",
					LogoURL:      "http://somesite.com/document/radios/d/d9cc/d9cc0fcf-2fa3-4a3b-8aa1-8681d68bc229.png",
					CurrentTrack: "Fiorello - Spiagge",
					ListnerCount: 3770,
				},
				Station{
					ID:           "1653990",
					Name:         "RockRadio.cf",
					MediaType:    "audio/aacp",
					BitRate:      "32",
					Genre:        "Country",
					Genre2:       "Hard Rock",
					Genre3:       "Rock",
					Genre4:       "Rockabilly",
					Genre5:       "Rock & Roll",
					LogoURL:      "http://somesite.com/document/radios/f/f41c/f41c19bf-625f-4a3b-9142-95b8358d86a1.png",
					CurrentTrack: "",
					ListnerCount: 3392,
				},
				Station{
					ID:           "1783457",
					Name:         "Radio FM4",
					MediaType:    "audio/mpeg",
					BitRate:      "160",
					Genre:        "Misc",
					Genre2:       "Electronic",
					Genre3:       "Rock",
					Genre4:       "Pop",
					Genre5:       "Hip Hop",
					LogoURL:      "http://somesite.com/document/radios/1/1bb3/1bb3b2af-27dc-4b1f-88f9-c1c6655d355d.jpg",
					CurrentTrack: "Best of FM4 Acoustic Sessions: Teil 4",
					ListnerCount: 3332,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SHOUTcastAPIClient{
				key: tt.fields.key,
			}
			got, err := s.GetByGenre(tt.args.genre, tt.args.mediaType, tt.args.bitRate, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("SHOUTcastAPIClient.GetByGenre() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SHOUTcastAPIClient.GetByGenre() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHOUTcastAPIClient_GeTuneInURLs(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://yp.shoutcast.com/sbin/tunein-station.m3u?id=1",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusOK, `
#EXTM3U
#EXTINF:-1,Hitradio OE3
http://185.85.28.140:80
#EXTINF:-1,Hitradio OE3
http://185.85.29.140:80
#EXTINF:-1,Hitradio OE3
http://185.85.29.140:8000
#EXTINF:-1,Hitradio OE3
http://185.85.28.140:8000
#EXTINF:-1,Hitradio OE3
http://185.85.29.166:8000
#EXTINF:-1,Hitradio OE3
http://185.85.28.166:8000
			`), nil
		},
	)
	httpmock.RegisterResponder("GET", "https://yp.shoutcast.com/sbin/tunein-station.m3u?id=2",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusOK, `
#EXTM3U
#EXTINF:-1,Hitradio OE3
http://185.85.28.140:80
		`), nil
		},
	)
	httpmock.RegisterResponder("GET", "https://yp.shoutcast.com/sbin/tunein-station.m3u?id=3",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusOK, `
#EXTM3U
#EXTINF:-1,Hitradio OE3
	`), nil
		},
	)

	type fields struct {
		key string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test load balanced station",
			args: args{
				id: "1",
			},
			wantErr: false,
			want: []string{
				"http://185.85.28.140:80",
				"http://185.85.29.140:80",
				"http://185.85.29.140:8000",
				"http://185.85.28.140:8000",
				"http://185.85.29.166:8000",
				"http://185.85.28.166:8000",
			},
		},
		{
			name: "Test single endpoint station",
			args: args{
				id: "2",
			},
			wantErr: false,
			want: []string{
				"http://185.85.28.140:80",
			},
		},
		{
			name: "Test down station",
			args: args{
				id: "3",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SHOUTcastAPIClient{
				key: tt.fields.key,
			}
			got, err := s.GeTuneInURLs(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SHOUTcastAPIClient.GeTuneInURLs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SHOUTcastAPIClient.GeTuneInURLs() = %v, want %v", got, tt.want)
			}
		})
	}
}
