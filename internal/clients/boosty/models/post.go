package models

import (
	"fmt"
	"strings"
	"time"
)

const (
	VideoDataType PostDataType = "ok_video"
	TextDataType  PostDataType = "text"
)

type V1GetPostsResponse struct {
	Extra struct {
		IsLast bool   `json:"isLast"`
		Offset string `json:"offset"`
	} `json:"extra"`
	Data []Post `json:"data"`
}

type Tag struct {
	Title string `json:"title"`
	Id    int    `json:"id"`
}

type Reactions struct {
	Laught  int `json:"laught"`
	Fire    int `json:"fire"`
	Heart   int `json:"heart"`
	Wonder  int `json:"wonder"`
	Sad     int `json:"sad"`
	Dislike int `json:"dislike"`
	Angry   int `json:"angry"`
	Like    int `json:"like"`
}

type PostStats struct {
	Comments  int       `json:"comments"`
	Reactions Reactions `json:"reactions"`
	Likes     int       `json:"likes"`
}

type PlayerURL struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type PostDataType string

type PostDetails struct {
	Id       string       `json:"id,omitempty"`
	Title    string       `json:"title,omitempty"`
	Type     PostDataType `json:"type"`
	Content  string       `json:"content,omitempty"`
	Duration int          `json:"duration,omitempty"`
	Width    int          `json:"width,omitempty"`
	Height   int          `json:"height,omitempty"`

	ShowViewsCounter bool `json:"showViewsCounter,omitempty"`
	ViewsCounter     int  `json:"viewsCounter,omitempty"`

	Modificator    string      `json:"modificator,omitempty"`
	Url            string      `json:"url,omitempty"`
	PreviewId      string      `json:"previewId,omitempty"`
	DefaultPreview string      `json:"defaultPreview,omitempty"`
	TimeCode       int         `json:"timeCode,omitempty"`
	PlayerUrls     []PlayerURL `json:"playerUrls,omitempty"`
	Preview        string      `json:"preview,omitempty"`
	UploadStatus   string      `json:"uploadStatus,omitempty"`
	Vid            string      `json:"vid,omitempty"`
	Complete       bool        `json:"complete,omitempty"`
	FailoverHost   string      `json:"failoverHost,omitempty"`
}

type Video struct {
	Id          string
	Title       string
	Duration    time.Duration
	Width       int
	Height      int
	PlaylistUrl string
}

type Posts []*Post

type Post struct {
	Id               string        `json:"id"`
	IntId            int           `json:"int_id"`
	Count            PostStats     `json:"count"`
	Tags             []Tag         `json:"tags"`
	Teaser           []interface{} `json:"teaser"`
	IsCommentsDenied bool          `json:"isCommentsDenied"`
	AdvertiserInfo   interface{}   `json:"advertiserInfo"`
	IsRecord         bool          `json:"isRecord"`
	IsDeleted        bool          `json:"isDeleted"`
	User             User          `json:"user"`
	IsWaitingVideo   bool          `json:"isWaitingVideo"`
	Donations        int           `json:"donations"`
	SignedQuery      string        `json:"signedQuery"`
	Title            string        `json:"title"`
	Details          []PostDetails `json:"data"`
	Price            int           `json:"price"`
	ShowViewsCounter bool          `json:"showViewsCounter"`
	HasAccess        bool          `json:"hasAccess"`
	IsLiked          bool          `json:"isLiked"`
	IsPublished      bool          `json:"isPublished"`
	PublishAt        int64         `json:"publishTime"`
	CreatedAt        int64         `json:"createdAt"`
	UpdatedAt        int64         `json:"updatedAt"`
}

func (p *Post) PublishDate() time.Time {
	return time.Unix(p.PublishAt, 0)
}

func (p *Post) CreatedDate() time.Time {
	return time.Unix(p.CreatedAt, 0)
}

func (p *Post) UpdatedDate() time.Time {
	return time.Unix(p.UpdatedAt, 0)
}

func (p *PostDetails) GetMasterPlaylistUrl() (string, error) {
	for _, x := range p.PlayerUrls {
		if x.Type == "hls" {
			return x.Url, nil
		}
	}
	return "", fmt.Errorf("not found hls url")
}

func (p *Post) GetVideos() []*Video {
	var res []*Video

	for i := 0; i < len(p.Details); i++ {
		if p.Details[i].Type == VideoDataType {
			v := &p.Details[i]
			url, err := v.GetMasterPlaylistUrl()
			if err != nil {
				continue
			}
			res = append(res, &Video{
				Id:          v.Id,
				Title:       v.Title,
				Duration:    time.Duration(v.Duration) * time.Second,
				Width:       v.Width,
				Height:      v.Height,
				PlaylistUrl: url,
			})
		}
	}

	return res
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

func (p *Post) String() string {
	var builder strings.Builder
	builder.Grow(1024)

	builder.WriteString(fmt.Sprintf("%-10s%s\n", "Пост:", p.Title))
	//builder.WriteString(fmt.Sprintf("%-15s%s\n", "Id:", p.Id))
	builder.WriteString(fmt.Sprintf("%-10s%s\n", "Создано:", p.CreatedDate().Format("02 Jan 15:04")))

	for i := 0; i < len(p.Details); i++ {
		if p.Details[i].Type == VideoDataType {
			v := &p.Details[i]
			builder.WriteString(fmt.Sprintf("=> %-15s%s\n", "Видео:", v.Title))
			builder.WriteString(fmt.Sprintf("   %-15s%s\n", "Id:", v.Id))
			builder.WriteString(fmt.Sprintf("   %-15s%vx%v\n", "Разрешение:", v.Width, v.Height))
			if v.ShowViewsCounter {
				builder.WriteString(fmt.Sprintf("   %-15s%d\n", "Просмотров:", v.ViewsCounter))
			}
			builder.WriteString(fmt.Sprintf("   %-15s%s\n", "Длительность:", time.Duration(v.Duration)*time.Second))
		}
	}

	return builder.String()
}

func (p *PostDetails) String() string {
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

func (p Posts) String() string {
	var builder strings.Builder
	builder.Grow(1024)

	for i := 0; i < len(p); i++ {
		for _, data := range p[i].Details {
			if data.Type == VideoDataType {
				builder.WriteString(fmt.Sprintf("%-15s%s\n", "Название:", data.Title))
				builder.WriteString(fmt.Sprintf("%-15s%s\n", "Длительность:", time.Duration(data.Duration)*time.Second))
				builder.WriteString(fmt.Sprintf("%-15s%vx%v\n", "Разрешение:", data.Width, data.Height))
				if data.ShowViewsCounter {
					builder.WriteString(fmt.Sprintf("%-15s%d\n", "Просмотров:", data.ViewsCounter))
				}

				builder.WriteString("---")
			}
		}
	}

	return builder.String()
}
