package components

import (
	g "github.com/zaptross/gorgeous"
	p "github.com/zaptross/portfoligo/internal/provider"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
)

func P(eb g.EB) *g.HTMLElement {
	theme := p.ThemeProvider.GetTheme()

	className := "basic-p"
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"color": theme.Base2,
		},
	})

	eb.ClassList = append(eb.ClassList, className)

	return g.P(eb)
}

func AuthorQuote(text string, author string, color string, classList []string) *g.HTMLElement {
	className := "author-quote" + ch.ClassNameSanitiser(color)
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"font-style":  "italic",
			"border-left": "solid thick " + color,
		},
	})

	return g.Blockquote(g.EB{
		ClassList: append(classList, className),
		Children: g.CE{
			P(g.EB{Text: text, ClassList: []string{ch.MarginL("1rem")}}),
			g.Em(g.EB{Text:"— " + author, ClassList: []string{ch.MarginL("1rem"),ch.FontColor(color)}}),
		},
	})
}
