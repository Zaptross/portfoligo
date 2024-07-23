package atoms

import (
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	"github.com/zaptross/portfoligo/internal/theme"
)

func AuthorQuoteText(text string, author string, color string, classList []string) *g.HTMLElement {
	return AuthorQuote(g.Text(text), g.Text(author), color, classList)
}

func AuthorQuote(text *g.HTMLElement, author *g.HTMLElement, color string, classList []string) *g.HTMLElement {
	t := theme.UseTheme()
	className := "author-quote" + ch.ClassNameSanitiser(color)
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"font-style":  "italic",
			"border-left": "solid thick " + color,
			"margin":      t.Spacing(5),
		},
	})

	return g.Blockquote(g.EB{
		ClassList: append(classList, className),
		Children: g.CE{
			P(g.EB{Children: g.CE{text}, ClassList: []string{ch.MarginL(t.Spacing(5))}}),
			g.Em(g.EB{Children: g.CE{g.Text("— "), author}, ClassList: []string{ch.MarginL(t.Spacing(4)), ch.FontColor(color)}}),
		},
	})
}
