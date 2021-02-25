package pixiv

import "time"

type Api interface {
	//
	Login()
	RefreshToken()
	// Illustration
	Illust(pid int) *Illust
	Related(pid, offset int) []*Illust
	// User
	Member(id int) *user
	MemberIllusts(id, offset int) []*Illust
	// Site
	Rank(mode string, offset int, date *time.Time) []*Illust
	SearchByTitle(title string) []*Illust
	SearchByTags(tag string) []*Illust
	SearchByTagsStrict(tag string) []*Illust
}

type Context struct {
	Token      string
	RefreshKey string
	Cookie     string
	Proxy      string
}

func NewContext(refreshKey string) Context {
	return Context{
		RefreshKey: refreshKey,
	}
}
