package models

type User struct {
	Id        int    `json:"id"`
	BlogUrl   string `json:"blogUrl"`
	AvatarUrl string `json:"avatarUrl"`
	Name      string `json:"name"`
	Flags     struct {
		ShowPostDonations bool `json:"showPostDonations"`
	} `json:"flags"`
	HasAvatar bool `json:"hasAvatar"`
}
