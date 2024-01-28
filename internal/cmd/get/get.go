package get

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"
	"runtime"
	"time"

	"boosty/internal/clients/boosty"
	"github.com/canhlinh/hlsdl"
	"github.com/go-resty/resty/v2"
	"github.com/grafov/m3u8"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var cmdGet = &cobra.Command{
		Use:   "get [command]",
		Short: "Display one or many resources.",
		Args:  cobra.MinimumNArgs(2),
		Run:   runGetCommand,
	}

	cmdGet.AddCommand(newCmdGetInfo(), newCmdGetPosts())

	return cmdGet
}

func runGetCommand(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	blogName, err := verifyName(args[0])
	checkError(err)

	videoId := args[1]
	dir := ""
	if len(args) > 2 {
		dir = args[2]
	}

	client, err := boosty.NewClient(blogName)
	checkError(err)

	fmt.Printf("Searching video with ID %s from %s:\n---\n", videoId, blogName)

	posts, err := client.GetPosts(ctx, 10)
	checkError(err)

	for _, post := range posts {
		videos := post.GetVideos()
		for _, video := range videos {
			if video.Id == videoId {
				fmt.Printf("Video found:\n")
				fmt.Println(video.String())
				fmt.Println("---")
				fmt.Println("Searching for the best quality...")

				p, err := getm3u8MasterPlaylist(video.PlaylistUrl)
				checkError(err)
				bestQuality := getMaxQualityVariant(p.Variants)
				playlistUrl, _ := url.Parse(video.PlaylistUrl)
				downloadUrl := "https://" + playlistUrl.Host + bestQuality.URI
				fmt.Printf("Best Quality URL: %s\n", downloadUrl)
				fmt.Println("---")
				fmt.Println("Starting download...")

				hlsDL := hlsdl.New(downloadUrl, nil, dir, videoId+".mp4", runtime.NumCPU(), true)

				filepath, err := hlsDL.Download()
				if err != nil {
					checkError(err)
				}

				fmt.Println("Saved => ", filepath)
				return
			}
		}
	}
	fmt.Println("Error: Video not found")
}

func getm3u8MasterPlaylist(url string) (*m3u8.MasterPlaylist, error) {
	client := resty.New()
	client.SetRetryCount(5).SetRetryWaitTime(time.Second)
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, errors.New(resp.Status())
	}
	p, t, err := m3u8.Decode(*bytes.NewBuffer(resp.Body()), false)
	if err != nil {
		return nil, err
	}
	if t != m3u8.MASTER {
		return nil, errors.New("not master playlist")
	}
	return p.(*m3u8.MasterPlaylist), err
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
