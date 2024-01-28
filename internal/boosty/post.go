package boosty

type V1GetPostsResponse struct {
	Extra struct {
		IsLast bool   `json:"isLast"`
		Offset string `json:"offset"`
	} `json:"extra"`
	Data []Post `json:"data"`
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
	Details          []PostDetail  `json:"data"`
	Price            int           `json:"price"`
	ShowViewsCounter bool          `json:"showViewsCounter"`
	HasAccess        bool          `json:"hasAccess"`
	IsLiked          bool          `json:"isLiked"`
	IsPublished      bool          `json:"isPublished"`
	PublishAt        int64         `json:"publishTime"`
	CreatedAt        int64         `json:"createdAt"`
	UpdatedAt        int64         `json:"updatedAt"`
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
