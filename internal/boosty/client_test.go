package boosty

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"boosty/internal/boosty/model"
	"github.com/canhlinh/hlsdl"
	"github.com/stretchr/testify/assert"
)

const (
	blogName = "dinablin"
	token    = "d39f342855a32e93e0b66025c1d9f90e4c5557a8b9d83b600a16e9bc03025ddd"
)

func TestClient(t *testing.T) {
	client, err := NewClientWithConfig(blogName, NewConfig().WithToken(token))
	assert.NoError(t, err)

	t.Run("Invalid token", func(t *testing.T) {
		conf := NewConfig().WithDebugEnable().WithToken("42")
		client, err := NewClientWithConfig(blogName, conf)
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
		posts, err := client.GetPosts(context.Background(), Args{
			Limit: 5,
		})
		assert.NoError(t, err)
		assert.Equal(t, 5, len(posts))
	})

	t.Run("Enumerable posts", func(t *testing.T) {
		ctx := context.Background()
		nextPost := client.GetPostIterator(Args{
			Limit:  1,
			Offset: "",
		})

		lastId := ""
		for i := 0; i < 10; i++ {
			p, err := nextPost(ctx)
			assert.NoError(t, err)
			assert.Equal(t, 1, len(p))
			assert.NotEqual(t, lastId, p[0].Id)
		}
	})

	t.Run("Find video", func(t *testing.T) {
		ctx := context.Background()
		videoId := "5106a9ec-104e-41df-9fda-afc8e5769983"
		nextPost := client.GetPostIterator(Args{
			Limit:  10,
			Offset: "",
		})

		for {
			posts, err := nextPost(ctx)
			assert.NoError(t, err)

			video, err := posts.FindVideo(videoId)
			if errors.Is(err, model.ErrVideoNotFound) {
				continue
			}
			assert.NoError(t, err)

			fmt.Printf("%+v\n", video)
			break
		}

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
