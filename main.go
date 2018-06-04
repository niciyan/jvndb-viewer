package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://jvndb.jvn.jp/ja/rss/jvndb_new.rdf")
	fmt.Printf("title: %v\n", feed.Title)
	fmt.Printf("description: %v\n", feed.Description)

	// func for filterling feed-items cvss-v3
	f := func(item *gofeed.Item) bool {
		cvssSets := item.Extensions["sec"]["cvss"]
		if len(cvssSets) > 0 {
			for _, cvss := range cvssSets {
				if cvss.Attrs["version"] == "3.0" && strings.Compare(cvss.Attrs["score"], "7.0") >= 0 {
					return true
				} else {
					continue
				}
			}
		}
		return false
	}

	for _, item := range filterRss(feed.Items, f) {
		fmt.Printf("[%v][%v]: %v\n", format(item.PublishedParsed), format(item.UpdatedParsed), item.Title)
	}

}

func format(t *time.Time) string {
	if t == nil {
		return "---------"
	}

	return t.Format("2006-01-02")
}

func filterRss(items []*gofeed.Item, f func(*gofeed.Item) bool) []*gofeed.Item {
	var feeds []*gofeed.Item
	for _, item := range items {
		if f(item) {
			feeds = append(feeds, item)
		}
	}
	return feeds
}
