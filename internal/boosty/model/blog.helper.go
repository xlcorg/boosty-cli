package model

import (
	"fmt"
	"strings"
)

func (b *Blog) String() string {
	var builder strings.Builder
	builder.Grow(1024)

	builder.WriteString(fmt.Sprintf("%-15s%s (%s)\n", "Блог:", b.Title, b.URL))
	builder.WriteString(fmt.Sprintf("%-15s%d\n", "Подписчиков:", b.Stats.Subscribers))
	builder.WriteString(fmt.Sprintf("%-15s%d\n", "Постов:", b.Stats.Posts))
	builder.WriteString(fmt.Sprintf("%-15s%s", "Подписка:", b.IsSubscribed))

	return builder.String()
}
