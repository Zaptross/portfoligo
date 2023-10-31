package pages

import (
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
		Content: func(_ types.PageDetails) *g.HTMLElement {
			return g.Div(g.EB{
				Children: []*g.HTMLElement{
					g.H1(g.EB{
						Text: "Welcome to my portfolio!",
					}),
					c.Link(g.Text("NoRhythm"), "/project/no-rhythm"),
				}},
			)
		},
	})
)