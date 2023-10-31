package types

import (
	"fmt"
	"time"

	g "github.com/zaptross/gorgeous"
)

type PageDetails struct {
	Title       string
	Description string
	Slug        string
	Type        string
	Tags        []string
	Written     time.Time
	Content     func(PageDetails) *g.HTMLElement
}

func (p PageDetails) GetRelativeURL() string {
	if p.Slug == "home" {
		return "/"
	}

	if p.Type == "root" {
		return fmt.Sprintf("/%s", p.Slug)
	}

	return fmt.Sprintf("/%s/%s", p.Type, p.Slug)
}
