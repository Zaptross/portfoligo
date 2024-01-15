package components

import (
	"fmt"
	"strings"
	"time"

	"github.com/samber/lo"
	g "github.com/zaptross/gorgeous"
	p "github.com/zaptross/portfoligo/internal/provider"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
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
			"border":           fmt.Sprintf("0.5rem solid %s", theme.Base02),
			"border-radius":    "0.5rem",
			"height":           "	100%",
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
			absoluteLinks(),
			layoutPageHeader(page),
			g.Div(g.EB{
				ClassList: []string{contentClass},
				Children: g.CE{
					page.Content(page),
				},
			}),
			layoutPageFooter(),
		},
	})
}

func layoutPageFooter() *g.HTMLElement {
	layoutPageFooterClass := "page-layout-footer"
	g.Class(&g.CSSClass{
		Selector: "." + layoutPageFooterClass,
		Props: g.CSSProps{
			"margin-bottom":  "0.5rem",
			"padding":        "0 2rem",
			"display":        "flex",
			"flex-direction": "row",
			"flex-wrap":      "wrap",
		},
	})
	g.Class(&g.CSSClass{
		Selector: "." + layoutPageFooterClass + " > a",
		Props: g.CSSProps{
			"margin-left": "0.25rem",
		},
	})
	return Col(
		g.CE{
			BackToTop(),
			Row(
				g.CE{
					P(g.EB{
						ClassList: []string{layoutPageFooterClass},
						Children: g.CE{
							g.Text("Powered by "),
							Link(g.Text("my code"), "https://github.com/zaptross/portfoligo"),
							g.Text(", written in "),
							Link(g.Text("Gorgeous"), "https://gorgeous.zaptross.com"),
							g.Text(
								fmt.Sprintf(" © %s Matthew Price", time.Now().Format("2006")),
							),
						},
					}),
				}, nil),
		}, []string{ch.MarginT("1.5rem")})
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
			"margin-bottom": "0.5rem",
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
			"transition":  "color 0.25s ease-in-out",
		},
	})
	g.Class(&g.CSSClass{
		Include:  true,
		Selector: "." + dateTagsClass + ":hover",
		Props: g.CSSProps{
			"color": theme.Violet,
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
				return LinkNav(g.H3(g.EB{
					ClassList: []string{dateTagsClass},
					Text:      tag,
				}), "/search/?q="+tag)
			})...,
		)
	}

	return g.Div(g.EB{
		ClassList: []string{dateTagsContainerClass},
		Children:  elements,
	})
}

func absoluteLinks() *g.HTMLElement {
	linksDivClass := "absolute-links"
	g.Class(&g.CSSClass{
		Selector: "." + linksDivClass,
		Props: g.CSSProps{
			"padding-top": "0.25rem",
			"height":      "max(0px, calc(10vh - 10vw))",
			"min-height":  "max(0px, calc(10vh - 10vw))",
		},
	})

	faCSS := g.CSSProps{
		"color":       p.ThemeProvider.GetTheme().Base1,
		"margin":      "0 0.25rem 0 0.25rem",
		"padding-top": "0.25rem",
		"transition":  "color 0.25s ease-in-out",
	}
	nameColorClass := "name-color"
	g.Class(&g.CSSClass{
		Selector: "." + nameColorClass,
		Props: g.CSSProps{
			"color":  p.ThemeProvider.GetTheme().Base1,
			"margin": "0 0.25rem",
		},
	})

	g.Class(&g.CSSClass{
		Include:  true,
		Selector: ".top-left > div > .linknav",
		Props: g.CSSProps{
			"margin-right": "0.5rem",
		},
	})

	addHoverColor(".fa-solid.fa-rss", "#F99000")
	addHoverColor(".fa-brands.fa-github", "#6e5494")
	addHoverColor(".fa-brands.fa-linkedin", "#0072B1")

	return g.Div(g.EB{
		ClassList: []string{linksDivClass},
		Children: g.CE{
			absoluteDiv(
				[]string{"top", "left"},
				g.CE{
					LinkNav(
						Row(g.CE{
							FAS("home", faCSS),
							g.H3(g.EB{
								ClassList: []string{nameColorClass},
								Text:      "Matthew Price",
							}),
						}, nil),
						"/",
					),
					Row(g.CE{
						LinkNav(g.Text("Projects"), "/projects"),
						LinkNav(g.Text("Blog"), "/blog"),
						LinkNav(g.Text("Search"), "/search"),
						LinkNav(g.Text("About"), "/about"),
					}, nil),
				},
			),
			absoluteDiv(
				[]string{"top", "right"},
				g.CE{
					Row(g.CE{
						LinkIcon(FAS("rss", faCSS), "/public/rss.xml"),
						LinkIcon(FAB("github", faCSS), "https://github.com/zaptross"),
						LinkIcon(FAB("linkedin", faCSS), "https://linkedin.com/in/mpdd"),
					}, nil),
				},
			),
		},
	})
}

func absoluteDiv(zeroing []string, children g.CE) *g.HTMLElement {
	zeroingClassName := strings.Join(zeroing, "-")
	cssProps := g.CSSProps{
		"position": "absolute",
		"margin":   "0.5rem",
	}

	for _, prop := range zeroing {
		cssProps[prop] = "0"
	}

	g.Class(&g.CSSClass{
		Selector: "." + zeroingClassName,
		Props:    cssProps,
	})

	return g.Div(g.EB{
		ClassList: []string{zeroingClassName},
		Children:  children,
	})
}

func addHoverColor(className string, color string) {
	g.Class(&g.CSSClass{
		Selector: className + ":hover",
		Include:  true,
		Props: g.CSSProps{
			"color": color + " !important",
		},
	})
}
