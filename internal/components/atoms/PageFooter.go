package atoms

import (
	"fmt"
	"time"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	"github.com/zaptross/portfoligo/internal/theme"
)

func PageFooter() *g.HTMLElement {
	t := theme.UseTheme()
	layoutPageFooterClass := "page-layout-footer"
	g.Class(&g.CSSClass{
		Selector: "." + layoutPageFooterClass,
		Props: g.CSSProps{
			"padding":        "0 2rem",
			"display":        "flex",
			"flex-direction": "row",
			"flex-wrap":      "wrap",
			"color":          t.Colors.Text.Secondary,
		},
	})
	g.Class(&g.CSSClass{
		Selector: "." + layoutPageFooterClass + " > a",
		Props: g.CSSProps{
			"margin-left": "0.25rem",
		},
	})

	ch.MediaPhone(layoutPageFooterClass, g.CSSProps{
		"flex-direction": "column",
	})

	return Col(
		g.CE{
			BackToTop(),
			Row(
				g.CE{
					P(g.EB{
						ClassList: []string{layoutPageFooterClass},
						Children: g.CE{
							g.Span(g.EB{
								Children: g.CE{
									g.Text("Powered by "),
									Link(g.Text("my code"), "https://github.com/zaptross/portfoligo"),
								},
							}),
							g.Span(g.EB{
								Children: g.CE{
									g.Text("Written in "),
									Link(g.Text("Gorgeous"), "https://gorgeous.zaptross.com"),
								},
							}),
							g.Span(g.EB{
								Text: fmt.Sprintf(" © %s Matthew Price", time.Now().Format("2006")),
							}),
						},
					}),
				}, nil),
		}, []string{ch.MarginT("1.5rem")})
}
