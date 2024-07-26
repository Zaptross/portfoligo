package atoms

import (
	"fmt"
	"math"
	"time"

	"github.com/samber/lo"
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	"github.com/zaptross/portfoligo/internal/theme"
	"github.com/zaptross/portfoligo/internal/types"
)

func ArticleHeader(page types.PageDetails) *g.HTMLElement {
	t := theme.UseTheme()

	pageLayoutHeaderClass := "page-layout-header"
	g.Class(&g.CSSClass{
		Selector: "." + pageLayoutHeaderClass,
		Props: g.CSSProps{
			"display":        "flex",
			"flex-direction": "column",
			"align-items":    "center",
			"width":          "100%",
		},
	})

	titleClass := "page-title"
	g.Class(&g.CSSClass{
		Selector: "." + titleClass,
		Props: g.CSSProps{
			"color":         t.Colors.Text.Primary,
			"margin-bottom": "0.5rem",
		},
	})
	ch.MediaPhone(titleClass, g.CSSProps{
		"max-width": "95vw",
		"font-size": fmt.Sprintf("clamp(6vw, %dvw, 8vw)", int(math.Ceil(95.0/float64(len(page.Title))))),
	})

	return g.Div(g.EB{
		ClassList: []string{pageLayoutHeaderClass},
		Children: g.CE{
			g.H1(g.EB{
				ClassList: []string{titleClass},
				Text:      page.Title,
			}),
			ArticleDateTags(page),
		},
	})
}

func ArticleDateTags(page types.PageDetails) *g.HTMLElement {
	t := theme.UseTheme()

	dateTagsContainerClass := "date-tags-container"
	g.Class(&g.CSSClass{
		Selector: "." + dateTagsContainerClass,
		Props: g.CSSProps{
			"display":         "flex",
			"flex-direction":  "row",
			"justify-content": "center",
			"width":           "100%",
		},
	})

	dateTagsClass := "date-tags"
	g.Class(&g.CSSClass{
		Selector: "." + dateTagsClass,
		Props: g.CSSProps{
			"color":       t.Colors.Text.Secondary,
			"margin-top":  "0.25rem",
			"margin-left": "0.25rem",
			"transition":  "color 0.25s ease-in-out",
		},
	})
	g.Class(&g.CSSClass{
		Include:  true,
		Selector: "." + dateTagsClass + ":hover",
		Props: g.CSSProps{
			"color": t.Colors.Text.Hover,
		},
	})

	elements := g.CE{}
	hasTags := page.GetTags() != nil && len(page.GetTags()) > 0
	hasDate := page.Written != time.Time{}

	if !hasDate && !hasTags {
		return g.Empty()
	}

	if hasDate {
		elements = append(elements,
			g.H3(g.EB{
				ClassList: []string{dateTagsClass},
				Text:      page.Written.Format("2006-01-02"),
			}),
		)
	}

	if hasDate && hasTags {
		elements = append(elements,
			g.H3(g.EB{
				ClassList: []string{dateTagsClass},
				Text:      "—",
			}),
		)
	}

	if hasTags {
		elements = append(elements,
			lo.Map(page.GetTags(), func(tag string, _ int) *g.HTMLElement {
				return LinkNav(g.H3(g.EB{
					ClassList: []string{dateTagsClass},
					Text:      tag,
				}), "/search/?q="+tag)
			})...,
		)
	}

	return Row(
		elements,
		[]string{dateTagsContainerClass, ch.FlexWrap()},
	)
}
