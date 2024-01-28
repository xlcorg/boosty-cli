package main

import (
	"context"
	"fmt"
	"time"

	"boosty/internal/clients/boosty"
	"boosty/internal/clients/boosty/models"
	"boosty/pkg/logger"
)

func main() {
	logger.Init(true)
	defer logger.Sync()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger.Info(ctx, "Starting service...")
	defer logger.Info(ctx, "Service stopped")

	config := boosty.NewConfig()
	client, err := boosty.NewClientWithConfig("dinablin", config)
	if err != nil {
		logger.Fatal(ctx, err, "failed to initialize boosty client")
		panic(err)
	}

	posts, err := client.GetPosts(ctx, 2)
	if err != nil {
		logger.Error(ctx, err, "client.GetPosts:")
		return
	}

	logger.Info(ctx, "Received posts", "count", len(posts))

	for _, post := range posts {
		fmt.Println(post)

		for _, pd := range post.Details {
			fmt.Println(pd.Type)
			if pd.Type == models.VideoDataType {
				masterUrl, err := pd.GetMasterPlaylistUrl()
				if err != nil {
					logger.Error(ctx, err, "GetMasterPlaylistUrl:")
				}

				fmt.Printf("master playlist url: %s", masterUrl)
			}
		}

	}

	//fmt.Println(posts.String())

	//
	//hlsDL := hlsdl.New(
	//	"https://vd342.mycdn.me/expires/1706463815131/srcIp/46.138.165.224/pr/42/srcAg/UNKNOWN/ms/185.226.53.58/type/5/sig/fePiVYfyWBs/ct/8/urls/185.226.52.17/clientType/18/id/5941128268408/video/",
	//	nil,
	//	"download", "test1.mp4", 20, true)
	//
	//filepath, err := hlsDL.Download()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(filepath)
}

//hlsDL := hlsdl.New(
//" https://vd342.mycdn.me/expires/1706463815131/srcIp/46.138.165.224/pr/42/srcAg/UNKNOWN/ms/185.226.53.58/type/5/sig/fePiVYfyWBs/ct/8/urls/185.226.52.17/clientType/18/id/5941128268408/video/",
////"https://vd342.mycdn.me/?expires=1706463815131&srcIp=46.138.165.224&pr=42&srcAg=UNKNOWN&ms=185.226.53.58&type=5&sig=fePiVYfyWBs&ct=0&urls=185.226.52.17&clientType=18&id=5941128268408",
////"https://vd348.mycdn.me/expires/1706463815132/srcIp/46.138.165.224/pr/42/srcAg/UNKNOWN/ms/45.136.22.71/type/5/sig/MCMLiwBlp5g/ct/8/urls/185.226.52.41/clientType/18/id/5934303349368/video/",
//nil,
//"download", "test1.ts", 20, true)
//
//// /expires/1706463815132/srcIp/46.138.165.224/pr/42/srcAg/UNKNOWN/ms/45.136.22.71/type/5/sig/MCMLiwBlp5g/ct/8/urls/185.226.52.41/clientType/18/id/5934303349368/video/
//// https://vd348.mycdn.me/video.m3u8?cmd=videoPlayerCdn&expires=1706463815132&srcIp=46.138.165.224&pr=42&srcAg=UNKNOWN&ms=45.136.22.71&type=2&sig=DmIsooG_kE8&ct=8&urls=185.226.52.41&clientType=18&id=5934303349368
//// https://vd348.mycdn.me/expires/1706463815132/srcIp/46.138.165.224/pr/42/srcAg/UNKNOWN/ms/45.136.22.71/type/5/sig/MCMLiwBlp5g/ct/8/urls/185.226.52.41/clientType/18/id/5934303349368/video/
//filepath, err := hlsDL.Download()
//if err != nil {
//panic(err)
//}
//
//fmt.Println(filepath)
