package pages

import (
	"time"

	g "github.com/zaptross/gorgeous"
	c "github.com/zaptross/portfoligo/internal/components"
	"github.com/zaptross/portfoligo/internal/types"
)

var (
	_ = registerPage(types.PageDetails{
		Title:       "Home",
		Description: "Home page for my portfolio",
		Slug:        "home",
		Type:        TYPE_ROOT,
		Written:     time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
		Tags: []string{
			TAG_GAME,
			TAG_LANG_CSHARP,
			TAG_AVAILABLE,
		},
		Content: g.Div(g.EB{
			Children: []*g.HTMLElement{
				g.H1(g.EB{
					Text: "Welcome to my portfolio!",
				}),
				c.LinkRel("NoRhythm", "/project/no-rhythm"),
			}},
		),
	})
)
