package model

import (
	"fmt"
	"strings"
	"time"
)

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

func (p *Post) String() string {
	var builder strings.Builder
	builder.Grow(1024)

	builder.WriteString(fmt.Sprintf("%-10s%s\n", "Пост:", p.Title))
	builder.WriteString(fmt.Sprintf("%-10s%s\n", "Создано:", p.CreatedAt.Time().Format("02 Jan 15:04")))
	builder.WriteString(fmt.Sprintf("%-10s%s\n", "Доступ", p.HasAccess))

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
