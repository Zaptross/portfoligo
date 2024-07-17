package organisms

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/samber/lo"
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	c "github.com/zaptross/portfoligo/internal/components"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	"github.com/zaptross/portfoligo/internal/theme"
)

type CarouselProps struct {
	Children  g.CE
	ClassList []string
}

//go:embed carousel.js
var carouselService string

func Carousel(props CarouselProps) *g.HTMLElement {
	t := theme.UseTheme()

	carouselClass := "carousel"
	g.Class(&g.CSSClass{
		Selector: "." + carouselClass,
		Props: g.CSSProps{
			"display":        "flex",
			"flex-direction": "row",
			"align-items":    "center",
			"overflow":       "hidden",
		},
	})

	carouselInnerClass := "carousel-inner"
	g.Class(&g.CSSClass{
		Selector: "." + carouselInnerClass,
		Props: g.CSSProps{
			"display":    "flex",
			"position":   "relative",
			"transition": "transform 0.5s ease-in-out",
		},
	})

	carouselItemClass := "carousel-item"
	g.Class(&g.CSSClass{
		Selector: "." + carouselItemClass,
		Props: g.CSSProps{
			"width": "100%",
		},
	})

	carouselButtonsClass := "carousel-buttons"
	g.Class(&g.CSSClass{
		Selector: "." + carouselButtonsClass,
		Props: g.CSSProps{
			"display":         "flex",
			"flex-direction":  "row",
			"justify-content": "space-between",
		},
	})

	carouselButtonClass := "carousel-button"
	g.Class(&g.CSSClass{
		Selector: "." + carouselButtonClass,
		Props: g.CSSProps{
			"top":              "50%",
			"width":            "5%",
			"background-color": "#0000",
			"border-color":     "#0000",
			"color":            t.Colors.Text.Secondary,
			"text-align":       "center",
			"line-height":      "50px",
			"font-size":        "30px",
			"transition":       "color 0.25s ease-in-out",
			"cursor":           "pointer",
		},
	})

	g.Class(&g.CSSClass{
		Selector: "." + carouselButtonClass + ":hover",
		Include:  true,
		Props: g.CSSProps{
			"color": t.Colors.Text.Primary,
		},
	})

	g.Class(&g.CSSClass{
		Selector: "." + carouselButtonClass + " > .fa-solid",
		Props: g.CSSProps{
			"transform": "translateX(-6px)",
		},
	})

	lr := []string{"left", "right"}
	buttons := g.CE{}
	buttonRefs := []*g.Reference{}
	lo.ForEach(lr, func(direction string, _ int) {
		dir := ""
		if direction == "right" {
			dir = "-"
		}
		btnClass := direction + "-btn"
		g.Class(&g.CSSClass{
			Selector: "." + btnClass,
			Props: g.CSSProps{
				"position":  "relative",
				"z-index":   "10",
				"transform": fmt.Sprintf("translate(%s100%%,-50%%)", dir),
			},
		})

		ref := g.CreateRef("button-" + direction)
		buttonRefs = append(buttonRefs, ref)
		buttons = append(buttons, ref.Get(g.Button(g.EB{
			ClassList: []string{carouselButtonClass, btnClass},
			Children: g.CE{
				a.FAS("chevron-"+direction, nil),
			},
		})))
	})

	lo.ForEach(props.Children, func(child *g.HTMLElement, _ int) {
		if child.ClassList == nil {
			child.ClassList = []string{}
		}
		child.ClassList = append(child.ClassList, carouselItemClass)
	})

	carouselInnerRef := g.CreateRef("carousel-inner-" + uuid.NewString())

	return c.Row(
		g.CE{
			g.Div(g.EB{
				ClassList: []string{"row"},
				Children: g.CE{
					buttons[0],
					g.Div(g.EB{
						ClassList: append([]string{carouselClass}, props.ClassList...),
						Children: g.CE{
							carouselInnerRef.Get(g.Div(g.EB{
								ClassList: []string{carouselInnerClass},
								Children:  props.Children,
							})),
						},
					}),
					buttons[1],
				},
				Script: configureCarousel(carouselInnerRef, buttonRefs),
			}),
		},
		[]string{"flex-wrap", ch.JustifyContent(ch.Content.Center)},
	)
}

func configureCarousel(inner *g.Reference, buttons []*g.Reference) g.JavaScript {
	g.Service("carousel", g.JavaScript(carouselService))

	buttonsJS := fmt.Sprintf(
		`[%s]`,
		strings.Join(
			lo.Map(
				buttons,
				func(ref *g.Reference, _ int) string {
					return string(ref.Javascript())
				},
			),
			",",
		),
	)

	return g.JavaScript(fmt.Sprintf(`
	initCarousel(%s, %s)
	`, inner.Javascript(), buttonsJS))
}
