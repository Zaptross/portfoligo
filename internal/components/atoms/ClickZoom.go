package atoms

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	"github.com/zaptross/portfoligo/internal/theme"
)

func ClickZoom(element *g.HTMLElement) *g.HTMLElement {
	t := theme.UseTheme()

	cz := "click-zoom"
	g.Class(&g.CSSClass{
		Selector: "." + cz,
		Include:  true,
		Props: g.CSSProps{
			"position":         "fixed",
			"z-index":          t.ZIndices.Overlay,
			"top":              "0",
			"left":             "0",
			"background-color": "rgba(0, 0, 0, 0.2)",
			"display":          "flex",
			"place-items":      "center",
			"place-content":    "center",
		},
	})
	g.Class(&g.CSSClass{
		Selector: "." + cz + " > *",
		Include:  true,
		Props: g.CSSProps{
			"top":        "10vh",
			"left":       "10vw",
			"width":      "80%",
			"height":     "80%",
			"max-width":  "80vw",
			"max-height": "80vh",
		},
	})

	return g.Div(g.EB{
		Script: g.JavaScript(fmt.Sprintf(`
		const clickZoomClass = "%s";
		const removeClickZoom = () => document.querySelectorAll("." + clickZoomClass).forEach(e => e.classList.remove(clickZoomClass))
		thisElement.addEventListener("click", () => thisElement.classList.contains(clickZoomClass) ? removeClickZoom() : thisElement.classList.add(clickZoomClass));
		document.addEventListener("keydown", (e) => {
			if (e.key === "Escape") {
				removeClickZoom();
			}
		});
	`, cz)),
		ClassList: []string{ch.H("100%"), ch.W("100%")},
		Children: []*g.HTMLElement{
			element,
		},
	})
}
