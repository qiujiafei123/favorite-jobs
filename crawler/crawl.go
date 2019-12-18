package crawler

import (
	"bytes"
	"favorite-jobs/utils"
	"github.com/PuerkitoBio/goquery"
)

type crawlCallback func(int, *goquery.Selection)

func Visit(url string, selector string, callback crawlCallback) {
	c := NewCrawler()
	html := c.Get(url)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	utils.CheckErr(err, "goquery.NewDocumentFromReader 出错")
	doc.Find(selector).Each(callback)
}
