package components

import (
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	"github.com/zaptross/portfoligo/internal/theme"
)

func P(eb g.EB) *g.HTMLElement {
	t := theme.UseTheme()

	className := "basic-p"
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"color": t.Colors.Text.Primary,
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
			g.Em(g.EB{Text: "— " + author, ClassList: []string{ch.MarginL("1rem"), ch.FontColor(color)}}),
		},
	})
}
