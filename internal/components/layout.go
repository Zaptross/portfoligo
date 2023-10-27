package components

import (
	"time"

	"github.com/samber/lo"
	g "github.com/zaptross/gorgeous"
	p "github.com/zaptross/portfoligo/internal/provider"
	"github.com/zaptross/portfoligo/internal/types"
)

func Layout(page types.PageDetails) *g.HTMLElement {
	theme := p.ThemeProvider.GetTheme()

	contentClass := "content-main"
	g.Class(&g.CSSClass{
		Selector: "." + contentClass,
		Props: g.CSSProps{
			// Decrease the sidebar width as the screen gets smaller
			"width": "max(50%, min(100%, calc(40vw + 270px)))",
		},
	})
	g.Class(&g.CSSClass{
		Include:  true,
		Selector: "." + contentClass + " > *",
		Props: g.CSSProps{
			// Increase padding as the screen gets smaller
			"padding":          "min(1rem, max(0rem, calc(1rem - 1.65vw)))",
			"background-color": theme.Base02,
			"height":           "100%",
		},
	})

	g.Class(&g.CSSClass{
		Include:  true,
		Selector: "html, body",
		Props: g.CSSProps{
			"background-color": theme.Base03,
			"height":           "100%",
			"width":            "100%",
			"display":          "flex",
			"flex-direction":   "column",
			"align-items":      "center",
		},
	})

	g.Class(&g.CSSClass{
		Include:  true,
		Selector: "body",
		Props: g.CSSProps{
			"margin": "0",
		},
	})

	return g.Body(g.EB{
		Children: g.CE{
			layoutPageHeader(page),
			g.Div(g.EB{
				ClassList: []string{contentClass},
				Children: g.CE{
					page.Content(),
				},
			}),
		},
	})
}

func layoutPageHeader(page types.PageDetails) *g.HTMLElement {
	theme := p.ThemeProvider.GetTheme()

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
			"color":         theme.Base2,
			"margin-bottom": "0rem",
		},
	})

	return g.Div(g.EB{
		ClassList: []string{pageLayoutHeaderClass},
		Children: g.CE{
			g.H1(g.EB{
				ClassList: []string{titleClass},
				Text:      page.Title,
			}),
			layoutPageDateTags(page),
		},
	})
}

func layoutPageDateTags(page types.PageDetails) *g.HTMLElement {
	theme := p.ThemeProvider.GetTheme()

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
			"color":       theme.Base1,
			"margin-top":  "0.25rem",
			"margin-left": "0.25rem",
		},
	})

	elements := g.CE{}
	hasTags := page.Tags != nil && len(page.Tags) > 0
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
			lo.Map(page.Tags, func(tag string, _ int) *g.HTMLElement {
				return g.H3(g.EB{
					ClassList: []string{dateTagsClass},
					Text:      tag,
				})
			})...,
		)
	}

	return g.Div(g.EB{
		ClassList: []string{dateTagsContainerClass},
		Children:  elements,
	})
}
