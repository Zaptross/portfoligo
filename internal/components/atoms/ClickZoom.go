package atoms

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	"github.com/zaptross/portfoligo/internal/theme"
)

func ClickZoom(element *g.HTMLElement, classList []string) *g.HTMLElement {
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
			"background-color": "rgba(0, 0, 0, 0.4)",
			"display":          "flex",
			"place-items":      "center",
			"place-content":    "center",
		},
	})
	g.Class(&g.CSSClass{
		Selector: "." + cz + " > *",
		Include:  true,
		Props: g.CSSProps{
			"place-items":   "center",
			"place-content": "center",
			"width":         "80%",
			"height":        "80%",
			"max-width":     "80vw",
			"max-height":    "80vh",
		},
	})
	g.Class(&g.CSSClass{
		Selector: "." + cz + " > * > *",
		Include:  true,
		Props: g.CSSProps{
			"width": "fit-content !important",
		},
	})

	return g.Div(g.EB{
		Script: g.JavaScript(fmt.Sprintf(`
		const clickZoomClass = "%s";
		const removeClickZoom = () => document.querySelectorAll("." + clickZoomClass).forEach(e => {
			e.classList.remove(clickZoomClass);
			e.parentElement.style.height = "";
		})
		thisElement.addEventListener("click", () => {
			const hasClickZoom = thisElement.classList.contains(clickZoomClass)
			if (hasClickZoom ) {
				removeClickZoom()
			} else {
				thisElement.parentElement.style.height = thisElement.parentElement.clientHeight + "px"
				thisElement.classList.add(clickZoomClass)
			}
		});
		document.addEventListener("keydown", (e) => {
			if (e.key === "Escape") {
				removeClickZoom();
			}
		});
	`, cz)),
		ClassList: tern(classList != nil, classList, []string{ch.H("100%"), ch.W("100%")}),
		Children: []*g.HTMLElement{
			element,
		},
	})
}

func tern[T any](condition bool, a T, b T) T {
	if condition {
		return a
	}
	return b
}
