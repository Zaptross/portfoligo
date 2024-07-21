package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/samber/lo"
	g "github.com/zaptross/gorgeous"
	c "github.com/zaptross/portfoligo/internal/components"
	o "github.com/zaptross/portfoligo/internal/components/organisms"
	"github.com/zaptross/portfoligo/internal/pages"
	"github.com/zaptross/portfoligo/internal/types"
)

func main() {
	globalClasses()
	setupThemes()

	// Render distribution files
	createDistDirectories()
	generateRSS()
	copyPublicToDist("")

	g.Class(&g.CSSClass{
		Include:  true,
		Selector: ":root",
		Props: g.CSSProps{
			"font-size":   "20px",
			"font-family": "'Lato', sans-serif",
		},
	})

	displayTab := 2
	longestName := 0
	lo.ForEach(pages.GetAllPages(), func(page types.PageDetails, _ int) {
		if len(page.Title) > longestName {
			longestName = len(page.Title)
		}
	})

	allPages := pages.GetAllPages()

	println("===  Rendering pages\t===")
	for _, pages := range pages.GetAllPagesByTypes() {
		for _, page := range pages {
			spacing := int(math.Ceil(float64(longestName)/float64(displayTab)))*displayTab - len(page.Title)
			println(fmt.Sprintf("%s%s-> %s", page.Title, strings.Repeat(" ", spacing), page.GetRelativeURL()))
			rendered := g.RenderStatic(
				g.Document(
					c.Head(page),
					o.Layout(page, allPages),
				),
			)

			writeRenderedHTML(page.GetRelativeURL(), rendered)
		}
	}
}
