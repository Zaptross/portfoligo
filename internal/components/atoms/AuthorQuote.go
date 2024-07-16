package atoms

import (
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	"github.com/zaptross/portfoligo/internal/theme"
)

func AuthorQuote(text string, author string, color string, classList []string) *g.HTMLElement {
	t := theme.UseTheme()
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
			P(g.EB{Text: text, ClassList: []string{ch.MarginL(t.Spacing(5))}}),
			g.Em(g.EB{Text: "— " + author, ClassList: []string{t.Spacing(5), ch.FontColor(color)}}),
		},
	})
}
