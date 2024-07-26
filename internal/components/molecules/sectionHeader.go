package molecules

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	"github.com/zaptross/portfoligo/internal/theme"
)

type SectionHeaderProps struct {
	TagKey    string
	Text      string
	ClassList []string
	Header    func(g.EB) *g.HTMLElement
	Color     string
}

func SectionHeader(pr SectionHeaderProps) *g.HTMLElement {
	t := theme.UseTheme()
	borderClass := "section-header-border-background"
	g.Class(&g.CSSClass{
		Selector: "." + borderClass,
		Props: g.CSSProps{
			"padding-left":  "1rem",
			"margin-right":  "1rem",
			"margin-top":    "0.5rem",
			"background":    t.Colors.Background.Secondary,
			"border":        "0.5rem solid " + t.Colors.Background.Secondary,
			"border-radius": "1rem",
			"transition":    "box-shadow 0.25s ease-in-out",
			"cursor":        "copy",
		},
	})

	g.Class(&g.CSSClass{
		Selector: "a[name] > div > *:hover",
		Include:  true,
		Props: g.CSSProps{
			"text-decoration": "underline !important",
		},
	})

	if pr.Color == "" {
		pr.Color = t.Colors.Text.Heading
	}

	class := "section-header"
	ch.MediaPhone(class, g.CSSProps{
		"margin-top":    t.Spacing(1),
		"margin-bottom": t.Spacing(1),
	})

	return a.Hashlink(g.CE{
		a.Row(g.CE{
			pr.Header(g.EB{
				ClassList: []string{class},
				Children: g.CE{
					a.FAS("hashtag", g.CSSProps{"padding-right": "0.5rem"}),
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
