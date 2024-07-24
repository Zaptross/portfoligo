package pages

import (
	"github.com/samber/lo"
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	m "github.com/zaptross/portfoligo/internal/components/molecules"
	"github.com/zaptross/portfoligo/internal/pages/projects"
	"github.com/zaptross/portfoligo/internal/types"
)

var (
	_ = registerPage(types.PageDetails{
		Title:       "Home",
		Description: "Home page for my portfolio",
		Slug:        "home",
		Type:        types.TYPE_ROOT,
		Content: func(_ types.PageDetails) *g.HTMLElement {
			return g.Div(g.EB{
				ClassList: []string{ch.FlexCol(), ch.JustifyContent(ch.Content.SpaceBetween)},
				Children: []*g.HTMLElement{
					m.SectionHeader(m.SectionHeaderProps{
						TagKey: "welcome",
						Header: g.H2,
						Text:   "Welcome to my portfolio!",
					}),
					a.P(g.EB{
						Text:      "🚧 This site is under construction. 🏗️",
						ClassList: []string{ch.Margin("0.5rem auto")},
					}),
					a.P(g.EB{
						Text:      "I am a full-stack polyglot software engineer with a passion for learning new skills, and this is a collection of my projects and blog posts.",
						ClassList: []string{ch.Margin("0.5rem")},
					}),
					a.P(g.EB{
						Text:      "Here are a few of my favourites projects and posts:",
						ClassList: []string{ch.Margin("0.5rem 0 0.5rem 0.5rem")},
					}),
					a.Col(lo.Map([]types.PageDetails{projects.NoRhythm, projects.SlayYourDragons}, func(page types.PageDetails, _ int) *g.HTMLElement {
						pv := m.Preview(page)
						pv.ClassList = append(pv.ClassList, ch.MarginB("0.5rem"))
						return pv
					}), []string{}),
				}},
			)
		},
	})
)
