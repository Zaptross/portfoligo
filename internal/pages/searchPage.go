package pages

import (
	"sort"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	c "github.com/zaptross/portfoligo/internal/components"
	m "github.com/zaptross/portfoligo/internal/components/molecules"
	"github.com/zaptross/portfoligo/internal/types"
)

var (
	_ = registerPage(types.PageDetails{
		Title:       "Search",
		Description: "Searching projects and blog posts.",
		Slug:        "search",
		Type:        types.TYPE_ROOT,
		Content: func(_ types.PageDetails) *g.HTMLElement {
			searchableParent := g.CreateRef("searchable-elements")
			searchableElements := searchableParent.Get(g.Div(g.EB{
				Children:  getSearchableElements(),
				ClassList: []string{ch.PadB("0.25rem"), ch.FlexCol(), ch.JustifyContent(ch.Content.SpaceBetween)},
			}))
			search, searchScript := m.Search(searchableParent, "searchable", []string{ch.FlexGrow(1)})

			return g.Div(g.EB{
				ClassList: []string{ch.FlexCol(), ch.JustifyContent(ch.Content.SpaceBetween)},
				Children: g.CE{
					search,
					searchableElements,
				},
				Script: searchScript,
			})
		},
	})
)

func getSearchableElements() g.CE {
	previews := []*g.HTMLElement{}
	pages := GetAllPages()

	sort.Slice(pages, func(i, j int) bool {
		return pages[i].Written.After(pages[j].Written)
	})

	for _, page := range pages {
		if page.Type == types.TYPE_ROOT {
			continue
		}

		pv := c.Preview(page)
		pv.ClassList = append(pv.ClassList, ch.MarginB("0.5rem"))
		previews = append(previews, pv)
	}

	return previews
}
