package components

import (
	"fmt"
	"strings"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	"github.com/zaptross/portfoligo/internal/theme"
	"github.com/zaptross/portfoligo/internal/types"
)

func Preview(page types.PageDetails) *g.HTMLElement {
	t := theme.UseTheme()
	previewClass := "preview-card"
	g.Class(&g.CSSClass{
		Selector: "." + previewClass,
		Props: g.CSSProps{
			"background-color": t.Colors.Background.Primary,
			"box-shadow":       "1px 8px 0.5rem " + t.Colors.Background.Secondary,
			"transition":       "transform 0.3s ease-in-out",
		},
	})
	mediaAspect := "(max-aspect-ratio: 7/9)"
	g.Media(mediaAspect, "."+previewClass, g.CSSProps{
		"flex-direction": "column",
	})

	cardImageClass := "preview-card-image"
	g.Class(&g.CSSClass{
		Selector: "." + cardImageClass,
		Props: g.CSSProps{
			"aspect-ratio":  "3/2",
			"padding-right": "clamp(0.25rem, 1vw, 1rem)",
			"object-fit":    "cover",
		},
	})
	g.Media(mediaAspect, "."+cardImageClass, g.CSSProps{
		"height":        "clamp(100px, 15vw, 300px)",
		"margin-bottom": "0.5rem",
		"padding-right": "0",
	})

	g.Class(&g.CSSClass{
		Selector: "." + previewClass + ":hover",
		Props: g.CSSProps{
			"transform": "scale(1.05)",
		},
	})

	tagsClass := "preview-card-tags"
	g.Class(&g.CSSClass{
		Selector: "." + tagsClass,
		Props: g.CSSProps{
			"color":     t.Colors.Text.Secondary,
			"font-size": "0.8rem",
			"margin":    "0",
		},
	})

	titleClass := "preview-card-title"
	g.Class(&g.CSSClass{
		Selector: "." + titleClass,
		Props: g.CSSProps{
			"margin-top":    "0",
			"margin-bottom": "0.3rem",
		},
	})

	mediaColumnClassTarget := "." + previewClass + " > div"
	g.Media(mediaAspect, mediaColumnClassTarget, g.CSSProps{
		"padding": "0 0.5rem",
	})

	link := a.LinkNav(
		a.Row(g.CE{
			a.ImageRel(fmt.Sprintf("/preview/%s.png", page.Slug), []string{cardImageClass}),
			a.Col(g.CE{
				g.H2(g.EB{
					Text:      page.Title,
					ClassList: []string{titleClass},
				}),
				g.P(g.EB{
					Text:      fmt.Sprintf("%s - %s", page.Written.Format("2006-01-02"), strings.Join(page.GetTags(), ", ")),
					ClassList: []string{tagsClass},
				}),
				a.P(g.EB{
					Text:      page.Description,
					ClassList: []string{},
				}),
			}, []string{ch.Margin("auto 0"), ch.PadR("clamp(0.25rem, 1vw, 0.5rem)")}),
		}, []string{previewClass}),
		page.GetRelativeURL(),
	)

	if link.Props == nil {
		link.Props = g.Props{}
	}

	link.Props["data-searchable"] = strings.Join(append([]string{
		page.Written.Format("2006-01-02"),
		page.Title,
		page.Description},
		page.GetTags()...), " ")

	return link
}
