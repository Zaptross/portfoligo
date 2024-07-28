package organisms

import (
	"github.com/samber/lo"
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	"github.com/zaptross/portfoligo/internal/theme"
	"github.com/zaptross/portfoligo/internal/types"
)

type SkillMapping struct {
	Icon   string
	IconFn func(string, g.CSSProps) *g.HTMLElement
	Filter bool
}

type SkillsGridProps struct {
	Skills []string

	// For skills that aren't directly related to an icon, we can use this map to
	// map the skill to an icon name.
	MapIcons map[string]SkillMapping

	// Allow style overrides
	Style g.CSSProps
}

func SkillsGrid(pr SkillsGridProps) *g.HTMLElement {
	if pr.MapIcons == nil {
		pr.MapIcons = map[string]SkillMapping{}
	}

	t := theme.UseTheme()
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

	fontColorClass := "sg-font-color"
	g.Class(&g.CSSClass{
		Selector: "a > ." + fontColorClass,
		Props: g.CSSProps{
			"transition": "none",
			"color":      t.Colors.Text.Primary,
		},
	})
	g.Class(&g.CSSClass{
		Selector: "a:hover > ." + fontColorClass,
		Props: g.CSSProps{
			"color":     t.Colors.Text.Hover,
			"transform": "scale(1.05)",
		},
	})

	iconColorClass := "sg-icon-color"
	g.Class(&g.CSSClass{
		Selector: "a > ." + iconColorClass + " > *",
		Props: g.CSSProps{
			"transition": "none",
		},
	})
	g.Class(&g.CSSClass{
		Selector: "a > ." + iconColorClass + " > img",
		Props: g.CSSProps{
			"filter": t.Colors.Icons.Primary,
		},
	})
	g.Class(&g.CSSClass{
		Selector: "a:hover > ." + iconColorClass,
		Props: g.CSSProps{
			"transform": "scale(1.05)",
		},
	})
	g.Class(&g.CSSClass{
		Selector: "a:hover > ." + iconColorClass + " > img",
		Props: g.CSSProps{
			"filter": t.Colors.Icons.Hover,
		},
	})

	textClass := ch.FontColor(t.Colors.Text.Primary)
	g.Class(&g.CSSClass{
		Selector: "a:hover > * > ." + textClass,
		Props: g.CSSProps{
			"color": t.Colors.Text.Hover,
		},
	})

	return g.Div(g.EB{
		Style:     pr.Style,
		ClassList: []string{skillGridClass},
		Children: lo.Map(pr.Skills, func(skill string, _ int) *g.HTMLElement {
			icon := skill
			iconFn := a.FAB
			needsFilter := false
			if mapping, ok := pr.MapIcons[skill]; ok {
				iconFn = mapping.IconFn
				needsFilter = mapping.Filter

				if mapping.Icon != "" {
					icon = mapping.Icon
				}
			}

			colClass := fontColorClass
			if needsFilter {
				colClass = iconColorClass
			}

			return a.LinkNav(
				a.Col(
					g.CE{
						iconFn(icon, g.CSSProps{"font-size": "2rem"}),
						g.Span(g.EB{Text: skill, ClassList: []string{textClass}}),
					},
					[]string{colClass, ch.AlignItems(ch.Align.Center)},
				), "/search/?q="+skill)
		}),
	})
}

func SkillsMappingDefaults() map[string]SkillMapping {
	return map[string]SkillMapping{
		types.TAG_TOOL_KUBE:                       {IconFn: a.SimpleIcon, Filter: true},
		types.TAG_LANG_TS:                         {IconFn: a.SimpleIcon, Filter: true},
		types.TAG_LANG_CSHARP:                     {IconFn: a.SimpleIcon, Filter: true},
		types.TAG_TOOL_UNREAL:                     {IconFn: a.SimpleIcon, Filter: true},
		types.TAG_TOOL_DOTNET:                     {IconFn: a.SimpleIcon, Filter: true},
		types.TAG_TOOL_DATADOG:                    {IconFn: a.SimpleIcon, Filter: true},
		(types.TAG_GAME + " " + types.TAG_DESIGN): {IconFn: a.FAS, Icon: "gamepad"},
	}
}
