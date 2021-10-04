package pixiv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"pixiv_api/network"
	"strconv"
	"time"
)

// Client is an interface for pixiv api.
type Client struct {
	Cxt  Context
	http http.Client
}

func (c *Client) init() {
	c.http = network.NewClient()
}

func (c *Client) Login() {
	c.init()
	if c.Cxt.Token == "" {
		c.RefreshToken()
	}
}

func (c *Client) RefreshToken() {
	fmt.Println("refresh token: " + c.Cxt.Token)
	form := url.Values{
		"client_id":      {clientId},
		"client_secret":  {clientSecret},
		"grant_type":     {"refresh_token"},
		"refresh_token":  {c.Cxt.RefreshKey},
		"include_policy": {"true"},
	}

	bytes := c.post(oauthUrl, form)

	oauth := &oauthResp{}
	err := json.Unmarshal(bytes, oauth)
	if err != nil || oauth.AccessToken == "" || oauth.RefreshToken == "" {
		panic(err)
	}

	c.Cxt.Token = oauth.AccessToken
	c.Cxt.RefreshKey = oauth.RefreshToken
	fmt.Println("refresh token success: " + c.Cxt.Token)
}

func (c *Client) Illust(pid int) *Illust {
	bytes := c.get(illustDetailUrl, url.Values{
		"illust_id": {strconv.Itoa(pid)},
	})

	resp := &illustResp{}
	parseJson(bytes, resp)

	return resp.Illust
}

func (c *Client) Related(pid, offset int) []*Illust {
	values := url.Values{"illust_id": {strconv.Itoa(pid)}}
	if offset != 0 {
		values.Set("offset", strconv.Itoa(offset))
	}
	bytes := c.get(illustRelatedUrl, values)

	resp := &illustsResp{}
	parseJson(bytes, resp)

	return resp.Illusts
}

func (c *Client) Member(id int) *user {
	bytes := c.get(userDetailUrl, url.Values{
		"user_id": {strconv.Itoa(id)}},
	)

	resp := &MemberResp{}
	parseJson(bytes, resp)

	return resp.User
}

func (c *Client) MemberIllusts(id, offset int) []*Illust {
	values := url.Values{"user_id": {strconv.Itoa(id)}}
	if offset != 0 {
		values.Set("offset", strconv.Itoa(offset))
	}
	bytes := c.get(userIllustUrl, values)

	resp := &illustsResp{}
	parseJson(bytes, resp)

	return resp.Illusts
}

func (c *Client) Rank(mode string, offset int, date *time.Time) []*Illust {
	values := url.Values{"mode": {mode}}
	if offset != 0 {
		values.Set("offset", strconv.Itoa(offset))
	}
	if date != nil {
		values.Set("date", date.Format("2006-01-02"))
	}
	bytes := c.get(rankUrl, values)

	resp := &illustsResp{}
	parseJson(bytes, resp)

	return resp.Illusts
}

func (c *Client) SearchByTitle(title string) []*Illust {
	values := url.Values{
		"search_target": {"title_and_caption"},
		"word":          {title},
	}
	bytes := c.get(searchUrl, values)

	resp := &illustsResp{}
	parseJson(bytes, resp)

	return resp.Illusts
}

func (c *Client) SearchByTags(tag string) []*Illust {
	values := url.Values{
		"search_target": {"partial_match_for_tags"},
		"word":          {tag},
	}
	bytes := c.get(searchUrl, values)

	resp := &illustsResp{}
	parseJson(bytes, resp)

	return resp.Illusts
}

func (c *Client) SearchByTagsStrict(tag string) []*Illust {
	values := url.Values{
		"search_target": {"exact_match_for_tags"},
		"word":          {tag},
	}
	bytes := c.get(searchUrl, values)

	resp := &illustsResp{}
	parseJson(bytes, resp)

	return resp.Illusts
}

func parseJson(bytes []byte, resp interface{}) {
	if err := json.Unmarshal(bytes, resp); err != nil {
		panic("Serializer error." + err.Error())
	}
}
