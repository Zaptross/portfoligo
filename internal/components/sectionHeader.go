package components

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	p "github.com/zaptross/portfoligo/internal/provider"
)

type SectionHeaderProps struct {
	TagKey    string
	Text      string
	ClassList []string
	Header    func(g.EB) *g.HTMLElement
	Color     string
}

func SectionHeader(pr SectionHeaderProps) *g.HTMLElement {
	t := p.ThemeProvider.GetTheme()
	borderClass := "section-header-border-background"
	g.Class(&g.CSSClass{
		Selector: "." + borderClass,
		Props: g.CSSProps{
			"padding-left":  "1rem",
			"margin-right":  "1rem",
			"margin-top":    "0.5rem",
			"background":    t.Base03,
			"border":        "0.5rem solid " + t.Base03,
			"border-radius": "1rem",
			"transition":    "box-shadow 0.25s ease-in-out",
			"pointer":       "cursor",
		},
	})

	g.Class(&g.CSSClass{
		Selector: "a[name] > div > *:hover",
		Include:  true,
		Props: g.CSSProps{
			"text-decoration": "underline !important",
			"cursor":          "copy !important",
		},
	})

	return Hashlink(g.CE{
		Row(g.CE{
			pr.Header(g.EB{
				Children: g.CE{
					FAS("hashtag", g.CSSProps{"padding-right": "0.5rem"}),
					g.Text(pr.Text),
				},
				OnClick: onClick(pr.TagKey),
			}),
		}, []string{ch.FontColor(pr.Color), borderClass}),
	}, pr.TagKey, pr.ClassList)
}

func onClick(hash string) g.JavaScript {
	return g.JavaScript(fmt.Sprintf(`(() => {
		window.location.hash = "#%s";
		%s
	})()`, hash, clipboardCopy(`window.location.href`)))
}
