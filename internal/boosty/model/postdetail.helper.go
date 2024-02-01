package model

//
//func (p *PostDetail) GetVideos() []*Video {
//	var res []*Video
//	for i := 0; i < len(p.Details); i++ {
//		if p.Details[i].Type == VideoDataType {
//			v := &p.Details[i]
//			url, err := v.GetMasterPlaylistUrl()
//			if err != nil {
//				continue
//			}
//			res = append(res, &Video{
//				Id:          v.Id,
//				Title:       v.Title,
//				Duration:    time.Duration(v.Duration) * time.Second,
//				Width:       v.Width,
//				Height:      v.Height,
//				PlaylistUrl: url,
//			})
//		}
//	}
//	return res
//}
