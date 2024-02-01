package get

import (
	"context"
	"fmt"
	"net/url"
	"runtime"
	"time"

	"boosty/internal/boosty/model"
	"boosty/pkg/util"

	"boosty/internal/boosty"
	"github.com/canhlinh/hlsdl"
	"github.com/spf13/cobra"
)

var blogName string

func NewCommand() *cobra.Command {
	var cmdGet = &cobra.Command{
		Use:   "get [command]",
		Short: "Display one or many resources.",
		Args:  cobra.ExactArgs(1),
		Run:   runGetCommand,
	}

	//cmdGetInfo := newCmdGetInfo()

	//cmdGet.AddCommand(cmdGetInfo, newCmdGetPosts())

	//cmdGetInfo.PersistentFlags().StringVarP(&blogName, "author", "a", "", "author blog")
	//cmdGet.PersistentFlags().StringVarP(&blogName, "author", "a", "", "author blog")
	//_ = cmdGetInfo.MarkPersistentFlagRequired("author")
	//_ = cmdGet.MarkPersistentFlagRequired("author")

	return cmdGet
}

func runGetCommand(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	blogName, _ := cmd.Flags().GetString("author")
	util.CheckError(util.VerifyName(blogName))

	videoId := args[1]
	dir := ""
	if len(args) > 2 {
		dir = args[2]
	}

	client, err := boosty.NewClient(blogName)
	util.CheckError(err)

	fmt.Printf("Searching video with ID %s from %s:\n---\n", videoId, blogName)

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
				fmt.Println("Starting download...")

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
