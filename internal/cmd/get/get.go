package get

import (
	"boosty/internal/cmd/flags"
	"boosty/internal/storage"
	"boosty/internal/util"
	"context"
	"fmt"
	"net/url"
	"runtime"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := initClientFromFlags(cmd)
	videoId, dir := parseParam(args)

	fmt.Printf("Searching video with ID %s:\n---\n", videoId)

	// TODO: refactor download video
	// video := client.SearchVideo(videoId)
	// hlsDL := hlsdl.New(video, downloadUrl)

	posts, err := client.GetPosts(ctx, 10)
	util.CheckError(err)

	for _, post := range posts {
		videos := post.GetVideos()
		for _, video := range videos {
			if video.Id == videoId {
				fmt.Printf("Video found:\n")
				fmt.Println(video.String())
				fmt.Println("---")
				fmt.Println("Searching for the best quality...")

				p, err := client.GetM3u8MasterPlaylist(video.PlaylistUrl)
				util.CheckError(err)
				bestQuality := model.GetMaxQualityVariant(p.Variants)
				playlistUrl, _ := url.Parse(video.PlaylistUrl)
				downloadUrl := "https://" + playlistUrl.Host + bestQuality.URI
				fmt.Printf("Best Quality URL: %s\n", downloadUrl)
				fmt.Println("---")
				fmt.Printf("Starting download... [%d CPU]\n", runtime.NumCPU())

				hlsDL := hlsdl.New(downloadUrl, nil, dir, videoId+".mp4", runtime.NumCPU(), true)

				filepath, err := hlsDL.Download()
				if err != nil {
					util.CheckError(err)
				}

				fmt.Println("Saved => ", filepath)
				return
			}
		}
	}
	fmt.Println("Error: Video not found")
}
