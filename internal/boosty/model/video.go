package model

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrVideoNotFound = errors.New("video not found")
)

type Video struct {
	Id          string
	Title       string
	Duration    time.Duration
	Width       int
	Height      int
	PlaylistUrl string
}

func (v *Video) String() string {
	var builder strings.Builder
	builder.Grow(1024)

	builder.WriteString(fmt.Sprintf("%-15s%s\n", "Видео:", v.Title))
	builder.WriteString(fmt.Sprintf("%-15s%s\n", "Id:", v.Id))
	builder.WriteString(fmt.Sprintf("%-15s%vx%v\n", "Разрешение:", v.Width, v.Height))
	builder.WriteString(fmt.Sprintf("%-15s%s\n", "Длительность:", v.Duration))
	builder.WriteString(fmt.Sprintf("%-15s%s", "URL:", v.PlaylistUrl))

	return builder.String()
}

func (p Posts) FindVideo(videoId string) (*Video, error) {
	for _, post := range p {
		videos := post.GetVideos()
		for _, v := range videos {
			if v.Id == videoId {
				return v, nil
			}
		}
	}
	return nil, ErrVideoNotFound
}
