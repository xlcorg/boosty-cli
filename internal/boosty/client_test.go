package boosty

import (
	"context"
	"fmt"
	"testing"

	"github.com/canhlinh/hlsdl"
	"github.com/stretchr/testify/assert"
)

const (
	DefaultBlogName = "dinablin"
)

func TestClient(t *testing.T) {
	client, err := NewClient(DefaultBlogName)
	assert.NoError(t, err)

	t.Run("Invalid token", func(t *testing.T) {
		conf := NewConfig().WithDebugEnable().WithToken("42")
		client, err := NewClientWithConfig(DefaultBlogName, conf)
		assert.NoError(t, err)

		_, err = client.GetBlog(context.Background())
		assert.Error(t, err)
	})

	t.Run("GetBlog", func(t *testing.T) {
		blog, err := client.GetBlog(context.Background())
		assert.NoError(t, err)

		assert.Equal(t, "dinablin", blog.URL)
	})

	t.Run("GetPosts", func(t *testing.T) {
		posts, err := client.GetPosts(context.Background(), 5)
		assert.NoError(t, err)

		assert.Equal(t, 5, len(posts))

	})
}

func TestDownloadVideo(t *testing.T) {
	hlsDL := hlsdl.New(
		" https://vd342.mycdn.me/expires/1706463815131/srcIp/46.138.165.224/pr/42/srcAg/UNKNOWN/ms/185.226.53.58/type/5/sig/fePiVYfyWBs/ct/8/urls/185.226.52.17/clientType/18/id/5941128268408/video/",
		//"https://vd342.mycdn.me/?expires=1706463815131&srcIp=46.138.165.224&pr=42&srcAg=UNKNOWN&ms=185.226.53.58&type=5&sig=fePiVYfyWBs&ct=0&urls=185.226.52.17&clientType=18&id=5941128268408",
		//"https://vd348.mycdn.me/expires/1706463815132/srcIp/46.138.165.224/pr/42/srcAg/UNKNOWN/ms/45.136.22.71/type/5/sig/MCMLiwBlp5g/ct/8/urls/185.226.52.41/clientType/18/id/5934303349368/video/",
		nil,
		"download", "test1.ts", 20, true)

	// /expires/1706463815132/srcIp/46.138.165.224/pr/42/srcAg/UNKNOWN/ms/45.136.22.71/type/5/sig/MCMLiwBlp5g/ct/8/urls/185.226.52.41/clientType/18/id/5934303349368/video/
	// https://vd348.mycdn.me/video.m3u8?cmd=videoPlayerCdn&expires=1706463815132&srcIp=46.138.165.224&pr=42&srcAg=UNKNOWN&ms=45.136.22.71&type=2&sig=DmIsooG_kE8&ct=8&urls=185.226.52.41&clientType=18&id=5934303349368
	// https://vd348.mycdn.me/expires/1706463815132/srcIp/46.138.165.224/pr/42/srcAg/UNKNOWN/ms/45.136.22.71/type/5/sig/MCMLiwBlp5g/ct/8/urls/185.226.52.41/clientType/18/id/5934303349368/video/
	filepath, err := hlsDL.Download()
	if err != nil {
		panic(err)
	}

	fmt.Println(filepath)
}
