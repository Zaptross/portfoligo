package organisms

import (
	"fmt"
	"strings"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	m "github.com/zaptross/portfoligo/internal/components/molecules"
	"github.com/zaptross/portfoligo/internal/theme"
	"github.com/zaptross/portfoligo/internal/types"
)

func PreviewExplained(context string, page types.PageDetails) *g.HTMLElement {
	t := theme.UseTheme()

	return PreviewWithContext(
		a.P(g.EB{
			ClassList: []string{
				ch.FontColorI(t.Colors.Text.Secondary),
				ch.Margin(fmt.Sprintf("%s %s", t.Spacing(2), t.Spacing(5))),
			},
			Text: context,
		}),
		page,
	)
}
func PreviewWithContext(context *g.HTMLElement, page types.PageDetails) *g.HTMLElement {
	t := theme.UseTheme()

	pwcClass := "preview-with-context"
	g.Class(&g.CSSClass{
		Selector: "." + pwcClass,
		Props: g.CSSProps{
			"border":           strings.Join([]string{t.Borders.Thin, t.Borders.Style.Default, t.Colors.Background.Secondary}, " "),
			"border-radius":    t.Radii.Default,
			"background-color": t.Colors.Background.Secondary,
			"padding":          t.Spacing(2),
		},
	})

	g.Class(&g.CSSClass{
		Selector: "." + pwcClass + " > blockquote",
		Props: g.CSSProps{
			"font-style": "normal",
		},
	})

	ch.MediaPhone(
		pwcClass,
		g.CSSProps{
			"padding": t.Spacing(0),
		},
	)

	return g.Div(g.EB{
		ClassList: []string{pwcClass},
		Children: []*g.HTMLElement{
			a.BlockQuote(
				m.Preview(page),
				t.Colors.Pallette.Green,
				nil,
			),
			context,
		},
	})
}
