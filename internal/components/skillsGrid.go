package components

import (
	"github.com/samber/lo"
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	p "github.com/zaptross/portfoligo/internal/provider"
)

type SkillMapping struct {
	Icon   string
	IconFn func(string, g.CSSProps) *g.HTMLElement
}

type SkillsGridProps struct {
	Skills []string

	// For skills that aren't directly related to an icon, we can use this map to
	// map the skill to an icon name.
	MapIcons map[string]SkillMapping
}

func SkillsGrid(pr SkillsGridProps) *g.HTMLElement {
	if pr.MapIcons == nil {
		pr.MapIcons = map[string]SkillMapping{}
	}

	t := p.ThemeProvider.GetTheme()
	skillGridClass := "skills-grid"
	g.Class(&g.CSSClass{
		Selector: "." + skillGridClass,
		Props: g.CSSProps{
			"display":               "grid",
			"grid-template-columns": "repeat(auto-fit, minmax(100px, 1fr))",
			"grid-gap":              "4rem",
			"margin":                "1rem",
			"align-items":           "center",
			"justify-items":         "center",
		},
	})

	fontColorClass := ch.FontColor(t.Base1)
	g.Class(&g.CSSClass{
		Selector: "a:hover > ." + fontColorClass,
		Props: g.CSSProps{
			"transition": "color 0.25s ease-in-out transform 0.25s ease-in-out",
			"color":      t.Violet,
			"transform":  "scale(1.05)",
		},
	})

	return g.Div(g.EB{
		ClassList: []string{skillGridClass},
		Children: lo.Map(pr.Skills, func(skill string, _ int) *g.HTMLElement {
			icon := skill
			iconFn := FAB
			if mapping, ok := pr.MapIcons[skill]; ok {
				iconFn = mapping.IconFn

				if mapping.Icon != "" {
					icon = mapping.Icon
				}
			}

			return LinkNav(
				Col(
					g.CE{
						iconFn(icon, g.CSSProps{"font-size": "2rem"}),
						g.Text(skill),
					},
					[]string{fontColorClass, ch.AlignItems(ch.Align.Center)},
				), "/search/?q="+skill)
		}),
	})
}
