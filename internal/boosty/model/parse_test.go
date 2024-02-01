package model

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
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
	p, listType, err := m3u8.Decode(*bytes.NewBuffer(playlistData), false)
	assert.NoError(t, err)
	assert.Equal(t, m3u8.MASTER, listType)
	masterpl := p.(*m3u8.MasterPlaylist)

	bestQuality := GetMaxQualityVariant(masterpl.Variants)
	fmt.Println(bestQuality)

	fmt.Println(bestQuality.URI)
}
