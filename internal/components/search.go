package components

import (
	_ "embed"
	"fmt"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	"github.com/zaptross/portfoligo/internal/theme"
)

//go:embed search.js
var searchScript string

func Search(collectionParent *g.Reference, dataProperty string, classList []string) (*g.HTMLElement, g.JavaScript) {
	t := theme.UseTheme()
	searchRef := g.CreateRef("search-input")

	searchInputClass := "search-input"
	g.Class(&g.CSSClass{
		Selector: "." + searchInputClass,
		Props: g.CSSProps{
			"margin-left":      "1rem",
			"padding-left":     "0.5rem",
			"background-color": t.Colors.Background.Primary,
			"border":           fmt.Sprintf("solid thin %s", t.Colors.Background.Secondary),
			"border-radius":    "0.5rem",
			"color":            t.Colors.Text.Secondary,
			"font-size":        "16px",
			"flex-grow":        "1",
		},
	})
	g.Class(&g.CSSClass{
		Selector: "input[type=text]." + searchInputClass + ":focus",
		Include:  true,
		Props: g.CSSProps{
			"border-color":  t.Colors.Background.Hover,
			"outline-color": t.Colors.Background.Hover,
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
