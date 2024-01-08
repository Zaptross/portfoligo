package pages

import (
	"github.com/samber/lo"
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	c "github.com/zaptross/portfoligo/internal/components"
	p "github.com/zaptross/portfoligo/internal/provider"
	"github.com/zaptross/portfoligo/internal/types"
)

var (
	_ = registerPage(types.PageDetails{
		Title:       "Home",
		Description: "Home page for my portfolio",
		Slug:        "home",
		Type:        TYPE_ROOT,
		Content: func(_ types.PageDetails) *g.HTMLElement {
			t := p.ThemeProvider.GetTheme()
			return g.Div(g.EB{
				ClassList: []string{ch.FlexCol(), ch.JustifyContent(ch.Content.SpaceBetween)},
				Children: []*g.HTMLElement{
					g.H1(g.EB{
						Text:      "Welcome to my portfolio!",
						ClassList: []string{ch.FontColor(t.Green), ch.Margin("0.5rem auto")},
					}),
					c.P(g.EB{
						Text:      "🚧 This site is under construction. 🏗️",
						ClassList: []string{ch.Margin("0.5rem auto")},
					}),
					c.P(g.EB{
						Text:      "I am a full-stack polyglot software engineer with a passion for learning new skills, and this is a collection of my projects and blog posts.",
						ClassList: []string{ch.Margin("0.5rem")},
					}),
					c.P(g.EB{
						Text:      "Here are a few of my favourites projects and posts:",
						ClassList: []string{ch.Margin("0.5rem 0 0.5rem 0.5rem")},
					}),
					c.Col(lo.Map([]types.PageDetails{NoRhythm, SlayYourDragons}, func(page types.PageDetails, _ int) *g.HTMLElement {
						pv := c.Preview(page)
						pv.ClassList = append(pv.ClassList, ch.MarginB("0.5rem"))
						return pv
					}), []string{}),
				}},
			)
		},
	})
)
