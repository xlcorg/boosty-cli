package get

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"runtime"
	"time"

	"boosty/internal/cmd/flags"
	"boosty/internal/storage"
	"boosty/internal/util"

	"boosty/internal/boosty"
	"boosty/internal/boosty/model"
	"github.com/canhlinh/hlsdl"
	"github.com/spf13/cobra"
)

var (
	store      storage.Storage
	postsLimit int
)

const (
	blogNameKey = "blog"
	tokenKey    = "token"
)

func NewCommand(s storage.Storage) *cobra.Command {
	store = s
	var cmdGet = &cobra.Command{
		Use:     "get [video id] {directory}",
		Short:   "Download a video by ID.",
		Args:    cobra.MinimumNArgs(1),
		Run:     runGetCommand,
		Aliases: []string{"download"},
	}

	return cmdGet
}

func initClientFromFlags(cmd *cobra.Command) *boosty.Client {
	blogName := flags.GetValue(blogNameKey, cmd, store)
	util.CheckError(util.VerifyName(blogName))

	token := flags.GetValue(tokenKey, cmd, store)
	config := boosty.NewConfig()
	if token != "" {
		config = config.WithToken(token)
	}

	client, err := boosty.NewClientWithConfig(blogName, config)
	util.CheckError(err)

	return client
}

func parseParam(args []string) (videoId string, directory string) {
	videoId = args[0]
	if len(args) > 1 {
		directory = args[1]
	}
	return
}

func runGetCommand(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := initClientFromFlags(cmd)
	videoId, dir := parseParam(args)

	fmt.Printf("Searching video with ID %s:\n---\n", videoId)

	nextPosts := client.GetPostIterator(boosty.Args{
		Limit: 10,
	})

	video := &model.Video{}
	for {
		posts, err := nextPosts(ctx)
		if errors.Is(err, boosty.ErrPostNotFound) {
			fmt.Printf("Not found video with ID %s\n", videoId)
			return
		}
		util.CheckError(err)

		video, err = posts.FindVideo(videoId)
		if err == nil {
			break
		}
		if errors.Is(err, model.ErrVideoNotFound) {
			continue
		}
		util.CheckError(err)
	}

	fmt.Printf("Video found:\n")
	fmt.Println(video.String())

	downloadUrl, err := getDownloadUrl(client, video)
	util.CheckError(err)

	fmt.Println("---")
	fmt.Printf("Starting download... [%d CPU]\n", runtime.NumCPU())

	hlsDL := hlsdl.New(downloadUrl, nil, dir, video.Title+".mp4", runtime.NumCPU(), true)

	filepath, err := hlsDL.Download()
	util.CheckError(err)
	fmt.Println("Saved => ", filepath)

}

func getDownloadUrl(client *boosty.Client, video *model.Video) (string, error) {
	p, err := client.GetM3u8MasterPlaylist(video.PlaylistUrl)
	if err != nil {
		return "", err
	}

	bestQuality := model.GetMaxQualityVariant(p.Variants)
	playlistUrl, _ := url.Parse(video.PlaylistUrl)

	return "https://" + playlistUrl.Host + bestQuality.URI, nil
}
