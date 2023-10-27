package main

import (
	"fmt"
	"strings"

	g "github.com/zaptross/gorgeous"
	c "github.com/zaptross/portfoligo/internal/components"
	"github.com/zaptross/portfoligo/internal/pages"
)

func main() {
	registerClasses()

	// Render distribution files
	createDistDirectories()
	copyPublicToDist()

	g.Class(&g.CSSClass{
		Include:  true,
		Selector: ":root",
		Props: g.CSSProps{
			"font-size": "22px",
		},
	})

	println("===  Rendering pages\t===")
	for _, pages := range pages.GetAllPagesByTypes() {
		for _, page := range pages {
			println(fmt.Sprintf("%s%s@ %s", page.Title, strings.Repeat(" ", 2*8-len(page.Title)), page.GetRelativeURL()))
			rendered := g.RenderStatic(
				g.Document(
					c.Head(page),
					c.Layout(page),
				),
			)

			writeRenderedHTML(page.GetRelativeURL(), rendered)
		}
	}
}
