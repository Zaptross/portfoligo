package components

import (
	_ "embed"
	"fmt"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	p "github.com/zaptross/portfoligo/internal/provider"
)

//go:embed search.js
var searchScript string

func Search(collectionParent *g.Reference, dataProperty string, classList []string) (*g.HTMLElement, g.JavaScript) {
	t := p.ThemeProvider.GetTheme()
	searchRef := g.CreateRef("search-input")

	searchInputClass := "search-input"
	g.Class(&g.CSSClass{
		Selector: "." + searchInputClass,
		Props: g.CSSProps{
			"margin-left":      "1rem",
			"padding-left":     "0.5rem",
			"background-color": t.Base02,
			"border":           fmt.Sprintf("solid thin %s", t.Base03),
			"border-radius":    "0.5rem",
			"color":            t.Base1,
			"font-size":        "16px",
			"flex-grow":        "1",
		},
	})
	g.Class(&g.CSSClass{
		Selector: "input[type=text]." + searchInputClass + ":focus",
		Include:  true,
		Props: g.CSSProps{
			"border-color":  t.Cyan,
			"outline-color": t.Cyan,
		},
	})

	return g.Div(g.EB{
		ClassList: append([]string{ch.FlexRow(), ch.JustifyContent(ch.Content.Center)}, classList...),
		Children: g.CE{
			P(g.EB{Text: "Search:", ClassList: []string{ch.Margin("auto 0")}}),
			searchRef.Get(g.Input(g.EB{
				ClassList: []string{searchInputClass},
				Props: g.Props{
					"placeholder": "Start typing to search...",
					"type":        "text",
				},
			})),
		},
	}), setupSearch(collectionParent, dataProperty, searchRef)
}

func setupSearch(collectionParent *g.Reference, dataProperty string, searchInput *g.Reference) g.JavaScript {
	g.Service("search", g.JavaScript(searchScript))

	// import the hidden class
	ch.Hidden()

	return g.JavaScript(fmt.Sprintf(`
		configureSearch(%s,"%s",%s)
	`, collectionParent.Javascript(), dataProperty, searchInput.Javascript()))
}
