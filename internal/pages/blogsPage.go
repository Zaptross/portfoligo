package pages

import (
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	c "github.com/zaptross/portfoligo/internal/components"
	"github.com/zaptross/portfoligo/internal/types"
	"sort"
)

var (
	_ = registerPage(types.PageDetails{
		Title:       "Blog",
		Description: "Blog posts I've written, about programming and other topics.",
		Slug:        TYPE_BLOG,
		Type:        TYPE_ROOT,
		Content: func(_ types.PageDetails) *g.HTMLElement {
			return g.Div(g.EB{
				ClassList: []string{ch.FlexCol(), ch.JustifyContent(ch.Content.SpaceBetween)},
				Children: append(g.CE{
					c.P(g.EB{
						Text:      "Here is a list of blog posts I've written:",
						ClassList: []string{ch.MarginL("0.5rem")},
					}),
				}, getBlogPreviews()...),
			})
		},
	})
)

func getBlogPreviews() g.CE {
	previews := []*g.HTMLElement{}
	pages := GetAllPagesByType(TYPE_BLOG)

	sort.Slice(pages, func(i, j int) bool {
		return pages[i].Written.After(pages[j].Written)
	})

	for _, page := range pages {
		pv := c.Preview(page)
		pv.ClassList = append(pv.ClassList, ch.MarginB("0.5rem"))
		previews = append(previews, pv)
	}

	return previews
}
