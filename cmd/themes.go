package main

import (
	g "github.com/zaptross/gorgeous"
	t "github.com/zaptross/portfoligo/internal/theme"
)

func setupThemes() {
	for _, theme := range t.AllThemes {
		t.SetupTheme(theme)
	}

	// Setup themes to cleanly switch between them
	g.Media("(prefers-reduced-motion: no-preference)", "div", g.CSSProps{
		"transition": "0.25s ease-in-out all",
	})
}
