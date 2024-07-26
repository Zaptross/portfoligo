package organisms

import (
	"fmt"
	"strings"

	g "github.com/zaptross/gorgeous"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	"github.com/zaptross/portfoligo/internal/theme"
	"github.com/zaptross/portfoligo/internal/types"
)

func Layout(page types.PageDetails, allPages []types.PageDetails) *g.HTMLElement {
	t := theme.UseTheme()

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
			"padding":          fmt.Sprintf("max(min(1rem, max(0rem, calc(1rem - 1.65vw))), %s)", t.Spacing(2)),
			"background-color": t.Colors.Background.Primary,
			"border":           fmt.Sprintf("0.5rem solid %s", t.Colors.Background.Primary),
			"border-radius":    "0.5rem",
			"height":           "	100%",
		},
	})

	g.Class(&g.CSSClass{
		Include:  true,
		Selector: "html, body",
		Props: g.CSSProps{
			"background-color": t.Colors.Background.Secondary,
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

	content := page.Content(page)

	if page.Series != "" {
		content.Children = append(content.Children, SeriesNav(page, allPages))
	}

	return g.Body(g.EB{
		ClassList: []string{a.ROOT_STYLE_CLASS},
		Children: g.CE{
			absoluteLinks(),
			a.ArticleHeader(page),
			g.Div(g.EB{
				ClassList: []string{contentClass},
				Children: g.CE{
					content,
				},
			}),
			a.PageFooter(),
		},
	})
}

func absoluteLinks() *g.HTMLElement {
	t := theme.UseTheme()
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
		"width":       "1.15rem",
		"color":       t.Colors.Text.Secondary,
		"margin":      "0 0.25rem 0 0.25rem",
		"padding-top": "0.25rem",
		"transition":  "color 0.25s ease-in-out",
	}
	nameColorClass := "name-color"
	g.Class(&g.CSSClass{
		Selector: "." + nameColorClass,
		Props: g.CSSProps{
			"color":  t.Colors.Text.Secondary,
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

	// These colors are brand specific
	addHoverColor(".fa-solid.fa-rss", "#F99000")
	addHoverColor(".fa-brands.fa-github", "#6e5494")
	addHoverColor(".fa-brands.fa-linkedin", "#0072B1")

	return g.Div(g.EB{
		ClassList: []string{linksDivClass},
		Children: g.CE{
			absoluteDiv(
				[]string{"top", "left"},
				g.CE{
					a.LinkNav(
						a.Row(g.CE{
							a.FAS("home", faCSS),
							g.H3(g.EB{
								ClassList: []string{nameColorClass},
								Text:      "Matthew Price",
							}),
						}, nil),
						"/",
					),
					a.Row(g.CE{
						a.LinkNav(g.Text("Projects"), "/projects"),
						a.LinkNav(g.Text("Blog"), "/blog"),
						a.LinkNav(g.Text("Search"), "/search"),
						a.LinkNav(g.Text("About"), "/about"),
					}, nil),
				},
			),
			absoluteDiv(
				[]string{"top", "right"},
				g.CE{
					a.Row(g.CE{
						a.ThemeSelector(),
						a.LinkIcon(a.FAS("rss", faCSS), "/public/rss.xml"),
						a.LinkIcon(a.FAB("github", faCSS), "https://github.com/zaptross"),
						a.LinkIcon(a.FAB("linkedin", faCSS), "https://linkedin.com/in/mpdd"),
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
