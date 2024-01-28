package models

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/grafov/m3u8"
	"github.com/stretchr/testify/assert"
)

//go:embed testdata/v1blogresponse.json
var v1blogresponse []byte

//go:embed testdata/v1getpostsreponse.json
var getPostResponseData []byte

func TestParse(t *testing.T) {

	t.Run("Parse V1GetBlogResponse", func(t *testing.T) {
		var blog V1GetBlogResponse
		err := json.Unmarshal(v1blogresponse, &blog)
		assert.NoError(t, err)

		fmt.Println(blog)
	})

	t.Run("Parse V1GetPostsResponse", func(t *testing.T) {
		var posts V1GetPostsResponse
		err := json.Unmarshal(getPostResponseData, &posts)
		assert.NoError(t, err)

		for _, post := range posts.Data {
			//fmt.Println(post.String())
			videos := post.GetVideos()

			for i, v := range videos {
				fmt.Printf("=> %d\n", i+1)
				fmt.Println(v.String())
			}
			fmt.Println("---")
		}
	})
}

//go:embed testdata/playlist.m3u8
var playlistData []byte

func TestParsePlaylist(t *testing.T) {
	p, listType, err := parseM3u8Playlist(playlistData)
	assert.NoError(t, err)
	assert.Equal(t, m3u8.MASTER, listType)
	masterpl := p.(*m3u8.MasterPlaylist)

	bestQuality := getMaxQualityVariant(masterpl.Variants)
	fmt.Println(bestQuality)

	fmt.Println(bestQuality.URI)
}

func parseM3u8Playlist(data []byte) (m3u8.Playlist, m3u8.ListType, error) {
	p, t, err := m3u8.Decode(*bytes.NewBuffer(data), false)
	if err != nil {
		return nil, 0, err
	}
	return p, t, err
}

func getMaxQualityVariant(v []*m3u8.Variant) *m3u8.Variant {
	var res *m3u8.Variant
	maxLen := 0
	for _, x := range v {
		if len(x.Resolution) > maxLen {
			res = x
			maxLen = len(x.Resolution)
		}
	}
	return res
}

func TestParseUrl(t *testing.T) {
	rawUrl := "https://vd477.mycdn.me/video.m3u8?cmd=videoPlayerCdn&expires=1706543942819&srcIp=46.138.165.224&pr=42&srcAg=UNKNOWN&ms=185.226.53.93&type=2&sig=fEhhnrPf_uQ&ct=8&urls=45.136.22.25&clientType=18&id=5916247984760"
	uri, err := url.Parse(rawUrl)

	assert.NoError(t, err)
	assert.Equal(t, "https://vd477.mycdn.me", uri.Host)

}
