package pixiv

type oauthResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type illustsResp struct {
	Illusts []*Illust `json:"illusts"`
	NextUrl string    `json:"next_url"`
}

type illustResp struct {
	Illust *Illust `json:"illust"`
}

type Illust struct {
	Id             int         `json:"id"`
	Title          string      `json:"title"`
	Type           string      `json:"type"`
	ImageUrls      *imageUrls  `json:"image_urls"`
	User           *user       `json:"user"`
	Tags           []tag       `json:"tags"`
	PageCount      int         `json:"page_count"`
	Width          int         `json:"width"`
	Height         int         `json:"height"`
	MetaSinglePage *metaPage   `json:"meta_single_page"`
	MetaPages      []*metaPage `json:"meta_pages"`
	TotalView      int         `json:"total_view"`
	TotalBookmarks int         `json:"total_bookmarks"`
}

type imageUrls struct {
	SquareMedium string `json:"square_medium"`
	Medium       string `json:"medium"`
	Large        string `json:"large"`
	Original     string `json:"original"`
}

type tag struct {
	Name           string `json:"name"`
	TranslatedName string `json:"translated_name"`
}

type metaPage struct {
	ImageUrls *imageUrls `json:"image_urls"`
}

type MemberResp struct {
	User *user `json:"user"`
}

type user struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Account string `json:"account"`
}
