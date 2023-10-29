package pages

import (
	"github.com/samber/lo"
	g "github.com/zaptross/gorgeous"
	c "github.com/zaptross/portfoligo/internal/components"
	"github.com/zaptross/portfoligo/internal/types"
)

var (
	_ = registerPage(types.PageDetails{
		Title:       "Projects",
		Description: "Projects I've worked on",
		Slug:        "projects",
		Type:        TYPE_ROOT,
		Content: func() *g.HTMLElement {
			return g.Div(g.EB{
				Children: []*g.HTMLElement{
					g.H1(g.EB{
						Text: "Matthew Price",
					}),
					c.P(g.EB{
						Text: "These are some of the projects I've worked on:",
					}),
					g.Div(g.EB{
						Style: g.CSSProps{
							"display":        "flex",
							"flex-direction": "column",
						},
						Children: lo.Map(GetAllPagesByType(TYPE_PROJECT), func(page types.PageDetails, _ int) *g.HTMLElement {
							return c.Link(g.Text(page.Title), page.GetRelativeURL())
						}),
					}),
				}},
			)
		},
	})
)
