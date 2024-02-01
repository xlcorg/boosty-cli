package model

import (
	"github.com/grafov/m3u8"
)

func GetMaxQualityVariant(v []*m3u8.Variant) *m3u8.Variant {
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
