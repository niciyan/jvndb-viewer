package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://support.gmocloud.com/info/atom_tech.xml")
	fmt.Printf("author: %v\n", feed.Author.Name)
	fmt.Printf("email: %v\n", feed.Author.Email)
	fmt.Printf("title: %v\n", feed.Title)
	fmt.Printf("description: %v\n", feed.Description)

	for idx, item := range feed.Items {
		fmt.Printf("%v: %v\n", idx, item.Title)
		fmt.Printf("   %v\n", item.Link)
	}
}
