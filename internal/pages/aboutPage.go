package pages

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	m "github.com/zaptross/portfoligo/internal/components/molecules"
	o "github.com/zaptross/portfoligo/internal/components/organisms"
	"github.com/zaptross/portfoligo/internal/components/specific"
	"github.com/zaptross/portfoligo/internal/theme"
	"github.com/zaptross/portfoligo/internal/types"
)

var (
	_ = registerPage(types.PageDetails{
		Title:       "About",
		Description: "About Me",
		Slug:        "about",
		Type:        types.TYPE_ROOT,
		Content: func(_ types.PageDetails) *g.HTMLElement {
			t := theme.UseTheme()
			return g.Div(g.EB{
				ClassList: []string{ch.FlexCol(), ch.JustifyContent(ch.Content.SpaceBetween)},
				Children: []*g.HTMLElement{
					specific.AboutBlock(),
					aboutPageWorkingInSoftwareBlock(),
					m.SectionHeader(m.SectionHeaderProps{
						TagKey: "skills",
						Text:   "Skills",
						Header: g.H2,
					}),
					a.P(g.EB{
						Text: "I have experience with a wide range of languages and tools, including:",
					}),
					o.SkillsGrid(o.SkillsGridProps{
						Skills: []string{
							types.TAG_LANG_GO, types.TAG_LANG_CSHARP, types.TAG_LANG_TS, types.TAG_LANG_REACT, types.TAG_TOOL_JAVA, types.TAG_TOOL_PYTHON,
							types.TAG_LANG_RUST, types.TAG_TOOL_DATADOG,
							types.TAG_TOOL_AWS, types.TAG_TOOL_DOCKER, types.TAG_TOOL_KUBE, types.TAG_TOOL_DOTNET, types.TAG_TOOL_NODE,
							types.TAG_TOOL_FIGMA, types.TAG_TOOL_UNITY, types.TAG_TOOL_UNREAL, types.TAG_GAME + " " + types.TAG_DESIGN,
						},
						MapIcons: o.SkillsMappingDefaults(),
					}),
					a.P(g.EB{
						Text: "And hobbyist experience with:",
					}),
					o.SkillsGrid(o.SkillsGridProps{
						Skills:   []string{types.TAG_ELECTRONICS, types.TAG_TOOL_KICAD, types.TAG_TOOL_FREECAD, types.TAG_TOOL_PHOTOSHOP},
						MapIcons: o.SkillsMappingDefaults(),
					}),
					m.SectionHeader(m.SectionHeaderProps{
						TagKey: "kind-words",
						Text:   "Kind Words People Have Said About Me",
						Header: g.H2,
					}),
					a.AuthorQuoteText(
						"Matt displays the intelligence, creativity and ingenuity of a seasoned veteran, despite [being in] his first formal role in software engineering. His passion and determination to do [things] the right way shines in all that he does. That skill and style is exceeded only by his humanity. Pricey is simply one of the most kind, compassionate and generous humans I have ever known - an example to others and one that others thus try to live to. It was an honour to exist beside him in our team. If you have the fortune to get his attention, I strongly recommend him!",
						"Damian, AppSec Lead at Swyftx",
						t.Colors.Pallette.Green,
						nil,
					),
					a.AuthorQuoteText(
						"Matthew has demonstrated exceptional skills as a team player, working well with others to complete assigned projects to a high quality. Matthew would be a tremendous asset for your company and has my highest recommendation.",
						"Robert, Sessional Academic at QUT",
						t.Colors.Pallette.Cyan,
						nil,
					),
					a.AuthorQuoteText(
						"Matt has always worked above expectations as a part of our team [and] has always been one of the most pleasant individuals I've had under my employ.",
						"Dale, Store Manager at Dominos",
						t.Colors.Pallette.Blue,
						nil,
					),
					m.SectionHeader(m.SectionHeaderProps{
						TagKey: "contact",
						Text:   "Getting in Touch",
						Header: g.H2,
					}),
					a.P(g.EB{
						Text: "If you'd like to get in touch, my preferred method of contact is via the linkedin button below, or in the top right.",
					}),
					a.Row(g.CE{
						a.LinkIcon(
							a.FAB("linkedin", g.CSSProps{"font-size": "2rem"}),
							"https://www.linkedin.com/in/mpdd/",
						),
					}, []string{ch.MarginZA()}),
				}},
			)
		},
	})
)

func aboutPageWorkingInSoftwareBlock() *g.HTMLElement {
	text, script := specific.YearsInSoftware()
	return a.P(g.EB{
		Children: g.CE{
			g.Span(g.EB{
				Text:   fmt.Sprintf("I have been working in software for %s and I have experience with a variety of languages and frameworks in both professional and hobbyist capacities. See below for more details.", text),
				Script: script,
			}),
		},
	})
}
