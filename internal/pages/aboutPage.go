package pages

import (
	g "github.com/zaptross/gorgeous"
	c "github.com/zaptross/portfoligo/internal/components"
	"github.com/zaptross/portfoligo/internal/types"
)

var (
	_ = registerPage(types.PageDetails{
		Title:       "About",
		Description: "About Me",
		Slug:        "about",
		Type:        TYPE_ROOT,
		Content: func(_ types.PageDetails) *g.HTMLElement {
			return g.Div(g.EB{
				Children: []*g.HTMLElement{
					g.H1(g.EB{
						Text: "Matthew Price",
					}),
					c.P(g.EB{
						Text: "I am a full stack polyglot developer.",
					}),
					c.Link(g.Text("NoRhythm"), "/project/no-rhythm"),
				}},
			)
		},
	})
)