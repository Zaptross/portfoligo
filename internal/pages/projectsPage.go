package pages

import (
	"sort"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	m "github.com/zaptross/portfoligo/internal/components/molecules"
	"github.com/zaptross/portfoligo/internal/types"
)

var (
	_ = registerPage(types.PageDetails{
		Title:       "Projects",
		Description: "Projects I've worked on",
		Slug:        types.TYPE_PROJECT,
		Type:        types.TYPE_ROOT,
		Content: func(_ types.PageDetails) *g.HTMLElement {
			return g.Div(g.EB{
				ClassList: []string{ch.FlexCol(), ch.JustifyContent(ch.Content.SpaceBetween)},
				Children: append(g.CE{
					a.P(g.EB{
						Text:      "Here are some of the projects I've worked on:",
						ClassList: []string{ch.MarginL("0.5rem")},
					}),
				}, getProjectPreviews()...),
			})
		},
	})
)

func getProjectPreviews() g.CE {
	previews := []*g.HTMLElement{}
	pages := GetAllPagesByType(types.TYPE_PROJECT)

	sort.Slice(pages, func(i, j int) bool {
		return pages[i].Written.After(pages[j].Written)
	})

	for _, page := range pages {
		pv := m.Preview(page)
		pv.ClassList = append(pv.ClassList, ch.MarginB("0.5rem"))
		previews = append(previews, pv)
	}

	return previews
}
