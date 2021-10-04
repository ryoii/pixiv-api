package pixiv

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"pixiv_api/util"
	"strings"
	"time"
)

func (c *Client) get(url string, value url.Values) []byte {
	resp, err := c.http.Do(c.getRequest(url, value.Encode()))
	if err != nil {
		panic("network error. " + err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode == 400 {
		c.RefreshToken()
		return c.get(url, value)
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	return bytes
}

func (c *Client) getStr(url string, value url.Values) string {
	return util.Byte2Str(c.get(url, value))
}

func (c *Client) post(url string, value url.Values) []byte {
	resp, err := c.http.Do(c.postRequest(url, strings.NewReader(value.Encode())))
	if err != nil {
		panic("network error. " + err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode == 400 {
		c.RefreshToken()
		return c.get(url, value)
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	return bytes
}

func (c *Client) postStr(url string, value url.Values) string {
	return util.Byte2Str(c.post(url, value))
}

func (c *Client) getRequest(rawUrl, param string) *http.Request {
	parse, _ := url.Parse(rawUrl)
	parse.RawQuery = param
	return c.defaultRequest(http.MethodGet, parse.String(), nil)
}

func (c *Client) postRequest(url string, body io.Reader) *http.Request {
	return c.defaultRequest(http.MethodPost, url, body)
}

func (c *Client) defaultRequest(method, url string, body io.Reader) *http.Request {
	request, _ := http.NewRequest(method, url, body)
	c.fillHeader(request)

	return request
}

func (c *Client) fillHeader(request *http.Request) {
	header := request.Header

	clientTime := time.Now().Format(timeFormatter)
	hash := fmt.Sprintf("%x", md5.Sum(util.Str2Byte(clientTime+hashSecret)))

	header.Add("Authorization", "Bearer "+c.Cxt.Token)
	header.Add("User-Agent", userAgent)
	header.Add("Accept-Language", acceptLanguage)
	header.Add("App-OS", appOS)
	header.Add("App-OS-Version", appVersion)
	header.Add("X-Client-Time", clientTime)
	header.Add("X-Client-Hash", hash)
	header.Add("Content-Type", contentType)
}
