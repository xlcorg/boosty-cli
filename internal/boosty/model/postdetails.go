package model

import (
	"fmt"
	"strings"
	"time"
)

type PostDataType string

const (
	VideoDataType PostDataType = "ok_video"
	TextDataType  PostDataType = "text"
)

type PostDetail struct {
	Id               string       `json:"id,omitempty"`
	Title            string       `json:"title,omitempty"`
	Type             PostDataType `json:"type"`
	Content          string       `json:"content,omitempty"`
	Duration         int          `json:"duration,omitempty"`
	Width            int          `json:"width,omitempty"`
	Height           int          `json:"height,omitempty"`
	ShowViewsCounter bool         `json:"showViewsCounter,omitempty"`
	ViewsCounter     int          `json:"viewsCounter,omitempty"`
	Modificator      string       `json:"modificator,omitempty"`
	Url              string       `json:"url,omitempty"`
	PreviewId        string       `json:"previewId,omitempty"`
	DefaultPreview   string       `json:"defaultPreview,omitempty"`
	TimeCode         int          `json:"timeCode,omitempty"`
	PlayerUrls       []PlayerURL  `json:"playerUrls,omitempty"`
	Preview          string       `json:"preview,omitempty"`
	UploadStatus     string       `json:"uploadStatus,omitempty"`
	Vid              string       `json:"vid,omitempty"`
	Complete         bool         `json:"complete,omitempty"`
	FailoverHost     string       `json:"failoverHost,omitempty"`
}

func (p *PostDetail) GetMasterPlaylistUrl() (string, error) {
	for _, x := range p.PlayerUrls {
		if x.Type == "hls" {
			return x.Url, nil
		}
	}
	return "", fmt.Errorf("not found hls url")
}

func (p *PostDetail) String() string {
	var builder strings.Builder
	builder.Grow(1024)

	builder.WriteString(fmt.Sprintf("%-15s%s\n", "Название:", p.Title))

	switch p.Type {
	case TextDataType:

	case VideoDataType:
		builder.WriteString(fmt.Sprintf("%-15s%s\n", "Длительность:", time.Duration(p.Duration)*time.Second))
		builder.WriteString(fmt.Sprintf("%-15s%vx%v\n", "Разрешение:", p.Width, p.Height))
		if p.ShowViewsCounter {
			builder.WriteString(fmt.Sprintf("%-15s%d\n", "Просмотров:", p.ViewsCounter))
		}
	}

	return builder.String()
}
