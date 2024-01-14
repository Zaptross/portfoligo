package components

import (
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	p "github.com/zaptross/portfoligo/internal/provider"
)

const backToTopScript = `
function checkShowBackToTop() {
	// Get the height of the document
	const backToTopButton = document.getElementById("back-to-top-button");
	const html = document.documentElement;
	const height = Math.max(body.scrollHeight, body.offsetHeight, html.clientHeight, html.scrollHeight, html.offsetHeight);

	// If the page is longer than 1 screen and a small margin, show the back to top button
	if (height > window.innerHeight * 1.01) {
		backToTopButton.classList.remove("hidden");
	} else {
		backToTopButton.classList.add("hidden");
	}
}
`

func BackToTop() *g.HTMLElement {
	g.Service("backToTop", g.JavaScript(backToTopScript))
	t := p.ThemeProvider.GetTheme()
	className := "back-to-top-button"
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"background-color": t.Blue,
			"padding":          "0.5rem",
			"font-size":        "0.8rem",
			"font-color":       t.Base03,
			"width":            "max-content",
			"margin":           "0 auto",
			"border-radius":    "0.5rem",
			"border":           "0.1rem solid " + t.Blue,
			"cursor":           "pointer",
			"transition":       "background-color 0.25s ease-in-out",
		},
	})
	g.Class(&g.CSSClass{
		Selector: "." + className + ":hover",
		Props: g.CSSProps{
			"background-color": t.Cyan,
			"border-color":     t.Cyan,
		},
	})

	hidden := ch.Hidden()
	return g.Button(g.EB{
		Children: g.CE{
			g.Text("Back to top"),
			FAS("turn-up", g.CSSProps{"margin-left": "0.3rem"}),
		},
		Id:        "back-to-top-button",
		OnClick:   g.JavaScript(`window.scroll({ top: 0, left: 0, behavior: "smooth" });`),
		ClassList: []string{hidden, className},
		Script:    "checkShowBackToTop()",
	})
}
