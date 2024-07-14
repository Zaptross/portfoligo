package themeselector

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
	"github.com/zaptross/portfoligo/internal/theme"
)

// ThemeSelector is a component that allows the user to select a theme.
// It returns a select element with the available themes as options.
func ThemeSelector() *g.HTMLElement {
	t := theme.UseTheme()
	options := g.CE{}

	for _, t := range theme.AllThemes {
		options = append(options,
			g.Option(g.EB{
				Text: t.Selector(),
				Props: g.Props{
					"value": t.Selector(),
				},
			}),
		)
	}

	themeSelectorClass := "theme-selector"
	g.Class(&g.CSSClass{
		Selector: "." + themeSelectorClass,
		Props: g.CSSProps{
			"padding":          t.Spacing(0) + " " + t.Spacing(1),
			"margin":           t.Spacing(0) + " " + t.Spacing(1),
			"border-radius":    t.Spacing(1),
			"border-color":     t.Colors.Text.Secondary,
			"border-width":     t.Borders.Thin,
			"border-style":     t.Borders.Style.Default,
			"background-color": t.Colors.Background.Secondary,
			"color":            t.Colors.Text.Secondary,
		},
	})

	return g.Select(g.EB{
		Children:  options,
		ClassList: []string{themeSelectorClass},
		Props: g.Props{
			"id":       "theme-selector",
			"onchange": "localStorage.setItem('theme-selector', document.body.className = event.target.value);",
		},
		Script: g.JavaScript(fmt.Sprintf(`
		if (localStorage.getItem('theme-selector')) {
			const selectedTheme = localStorage.getItem('theme-selector');
			document.body.className = selectedTheme;
			thisElement.value = selectedTheme;
			return;
		}

		// ensure the theme matches the user's preference
		if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
			const selectedTheme = '%s';
			document.body.className = selectedTheme;
			thisElement.value = selectedTheme;
			localStorage.setItem('theme-selector', selectedTheme);
		}
		if (window.matchMedia && window.matchMedia('(prefers-color-scheme: light)').matches) {
			const selectedTheme = '%s';
			document.body.className = selectedTheme;
			thisElement.value = selectedTheme;
			localStorage.setItem('theme-selector', selectedTheme);
		}
			`,
			theme.VSCodeDark.Selector(),
			theme.VSCodeLight.Selector(),
		)),
	})
}
