package main

import (
	"os"
	"sort"
	"time"

	"github.com/gorilla/feeds"
	"github.com/zaptross/portfoligo/internal/pages"
	"github.com/zaptross/portfoligo/internal/types"
)

func generateRSS() {
	pages := append(pages.GetAllPagesByType(types.TYPE_BLOG), pages.GetAllPagesByType(types.TYPE_PROJECT)...)
	sort.SliceStable(pages, func(i, j int) bool {
		return pages[i].Written.After(pages[j].Written)
	})

	feed := &feeds.Feed{
		Title:       "matthewprice.dev posts",
		Link:        &feeds.Link{Href: "https://matthewprice.dev"},
		Description: "Posts from matthewprice.dev",
		Author:      &feeds.Author{Name: "Matthew Price", Email: "matt@matthewprice.dev"},
		Created:     time.Now(),
	}

	feed.Items = make([]*feeds.Item, len(pages))

	for i, page := range pages {
		feed.Items[i] = &feeds.Item{
			Title:       page.Title,
			Link:        &feeds.Link{Href: "https://matthewprice.dev" + page.GetRelativeURL()},
			Description: page.Description,
			Author:      feed.Author,
			Created:     page.Written,
		}
	}

	rss, err := feed.ToRss()
	if err != nil {
		panic(err)
	}

	os.WriteFile("dist/public/rss.xml", []byte(rss), 0644)
}
