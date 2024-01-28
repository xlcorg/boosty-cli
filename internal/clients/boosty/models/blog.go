package models

import (
	"fmt"
	"strings"
)

type Blog struct {
	Title        string
	URL          string
	Stats        BlogCount
	AccessRights BlogAccessRights
	IsSubscribed bool
}

type BlogCount struct {
	Posts       int `json:"posts"`
	Subscribers int `json:"subscribers"`
}

type BlogAccessRights struct {
	CanSetPayout      bool `json:"canSetPayout"`
	CanCreate         bool `json:"canCreate"`
	CanDeleteComments bool `json:"canDeleteComments"`
	CanView           bool `json:"canView"`
	CanEdit           bool `json:"canEdit"`
	CanCreateComments bool `json:"canCreateComments"`
}

type BlogDescription struct {
	Content     string `json:"content"`
	Modificator string `json:"modificator,omitempty"`
	Type        string `json:"type"`
	Url         string `json:"url,omitempty"`
	Explicit    bool   `json:"explicit,omitempty"`
}

type BlogFlags struct {
	HasTargets             bool `json:"hasTargets"`
	ShowPostDonations      bool `json:"showPostDonations"`
	AllowIndex             bool `json:"allowIndex"`
	HasSubscriptionLevels  bool `json:"hasSubscriptionLevels"`
	IsRssFeedEnabled       bool `json:"isRssFeedEnabled"`
	AllowGoogleIndex       bool `json:"allowGoogleIndex"`
	AcceptDonationMessages bool `json:"acceptDonationMessages"`
}

type BlogOwner struct {
	AvatarUrl string `json:"avatarUrl"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	HasAvatar bool   `json:"hasAvatar"`
}

type BlogSocialLink struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type BlogSocialLinks = []BlogSocialLink

type V1GetBlogResponse struct {
	BlogUrl                string            `json:"blogUrl"`
	Count                  BlogCount         `json:"count"`
	IsBlackListed          bool              `json:"isBlackListed"`
	PublicWebSocketChannel string            `json:"publicWebSocketChannel"`
	HasAdultContent        bool              `json:"hasAdultContent"`
	IsOwner                bool              `json:"isOwner"`
	IsTotalBaned           bool              `json:"isTotalBaned"`
	CoverUrl               string            `json:"coverUrl"`
	AccessRights           BlogAccessRights  `json:"accessRights"`
	Description            []BlogDescription `json:"description"`
	IsSubscribed           bool              `json:"isSubscribed"`
	Flags                  BlogFlags         `json:"flags"`
	SignedQuery            string            `json:"signedQuery"`
	AllowedPromoTypes      []string          `json:"allowedPromoTypes"`
	Title                  string            `json:"title"`
	IsBlackListedByUser    bool              `json:"isBlackListedByUser"`
	IsReadOnly             bool              `json:"isReadOnly"`
	Subscription           interface{}       `json:"subscription"`
	SubscriptionKind       string            `json:"subscriptionKind"`
	Owner                  BlogOwner         `json:"owner"`
	SocialLinks            BlogSocialLinks   `json:"socialLinks"`
}

func (b *Blog) String() string {
	var builder strings.Builder
	builder.Grow(1024)

	builder.WriteString(fmt.Sprintf("%-15s%s (%s)\n", "Блог:", b.Title, b.URL))
	builder.WriteString(fmt.Sprintf("%-15s%d\n", "Подписчиков:", b.Stats.Subscribers))
	builder.WriteString(fmt.Sprintf("%-15s%d\n", "Постов:", b.Stats.Posts))
	builder.WriteString(fmt.Sprintf("%-15s%s", "Подписка:", formatBoolean(b.IsSubscribed)))

	return builder.String()
}

func formatBoolean(value bool) string {
	if value {
		return "ДА"
	}
	return "НЕТ"
}
