package crawler

import (
	"favorite-jobs/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type crawler struct {
	client *http.Client
	Header map[string]string
	UserAgent []string
	ProxySwitch bool
}

type Option func(*crawler)

func NewCrawler(opts ...Option) *crawler {
	c := crawler{
		client: &http.Client{},
	}

	for _, opt := range opts {
		opt(&c)
	}
	return &c
}

func UserAgent(uas ...string) func(*crawler) {
	return func(c *crawler) {
		for _, ua := range uas {
			c.UserAgent = append(c.UserAgent, ua)
		}
	}
}

func WithHttpProxy() func(*crawler) {
	return func(c *crawler) {
		c.ProxySwitch = true
	}
}

func WithHeader(key, val string) func(*crawler) {
	return func(c *crawler) {
		c.Header[key] = val
	}
}


func (c *crawler) http (method string, link string, values url.Values) []byte {
	method = strings.ToUpper(method)

	req, err := http.NewRequest(method, link, strings.NewReader(values.Encode()))
	utils.CheckErr(err, "创建 http.NewRequest 出错")
	c.setHeader(req)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")

	c.proxy()

	resp, err := c.client.Do(req)
	utils.CheckErr(err, "发送 client.Do 出错")
	if resp.StatusCode != 200 {
		fmt.Println("返回码不是200:", link)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	utils.CheckErr(err, "解析 resp.Body 出错")
	return []byte(body)
}

func (c *crawler) setHeader(r *http.Request) {
	for key, val := range c.Header {
		r.Header.Set(key, val)
	}
}

func (c *crawler) proxy() {
	proxyPool, err := utils.ReadProxyList("./config/proxys.txt")
	utils.CheckErr(err, "生成proxy列表失败")

	n := utils.RandNum(len(proxyPool))
	proxyUrl, _ := url.Parse(proxyPool[n])
	fmt.Printf("代理IP: %s \n", proxyPool[n])
	tr := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	c.client.Transport = tr
}


func (c *crawler) Get(url string) []byte {
	return c.http("GET", url, nil)
}

