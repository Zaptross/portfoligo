package pages

import (
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	c "github.com/zaptross/portfoligo/internal/components"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	m "github.com/zaptross/portfoligo/internal/components/molecules"
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
					m.SectionHeader(m.SectionHeaderProps{
						TagKey:    "about",
						Text:      "Hi 👋, I'm Matt.",
						Header:    g.H1,
						ClassList: []string{ch.MarginL("1rem")},
					}),
					a.P(g.EB{
						Text:      "I am a full-stack polyglot software engineer with a passion for learning new skills and technologies.",
						ClassList: []string{ch.Margin("0.25rem 1rem")},
					}),
					a.P(g.EB{
						Text:      "In my spare time I love to play videogames and boardgames, develop fun ideas and projects, and work on hobbyist electronics projects.",
						ClassList: []string{ch.Margin("0.25rem 1rem")},
					}),
					m.SectionHeader(m.SectionHeaderProps{
						TagKey:    "skills",
						Text:      "Skills",
						Header:    g.H2,
						ClassList: []string{ch.MarginL("1rem")},
					}),
					a.P(g.EB{
						Text:      "I have experience with a wide range of languages and tools, including:",
						ClassList: []string{ch.MarginL("1rem")},
					}),
					c.SkillsGrid(c.SkillsGridProps{
						Skills: []string{
							types.TAG_LANG_GO, types.TAG_LANG_CSHARP, types.TAG_LANG_TS, types.TAG_LANG_REACT, types.TAG_LANG_RUST,
							types.TAG_TOOL_AWS, types.TAG_TOOL_DOCKER, types.TAG_TOOL_KUBE, types.TAG_TOOL_DOTNET, types.TAG_TOOL_NODE,
							types.TAG_TOOL_FIGMA, types.TAG_TOOL_UNITY, types.TAG_TOOL_UNREAL, types.TAG_GAME + " " + types.TAG_DESIGN,
						},
						MapIcons: map[string]c.SkillMapping{
							types.TAG_TOOL_KUBE:                       {IconFn: a.SimpleIcon, Filter: true},
							types.TAG_LANG_TS:                         {IconFn: a.SimpleIcon, Filter: true},
							types.TAG_LANG_CSHARP:                     {IconFn: a.SimpleIcon, Filter: true},
							types.TAG_TOOL_UNREAL:                     {IconFn: a.SimpleIcon, Filter: true},
							types.TAG_TOOL_DOTNET:                     {IconFn: a.SimpleIcon, Filter: true},
							(types.TAG_GAME + " " + types.TAG_DESIGN): {IconFn: a.FAS, Icon: "gamepad"},
						},
					}),
					a.P(g.EB{
						Text:      "And hobbyist experience with:",
						ClassList: []string{ch.MarginL("1rem")},
					}),
					c.SkillsGrid(c.SkillsGridProps{
						Skills: []string{types.TAG_ELECTRONICS, types.TAG_TOOL_KICAD, types.TAG_TOOL_FREECAD, types.TAG_TOOL_PHOTOSHOP},
						MapIcons: map[string]c.SkillMapping{
							types.TAG_ELECTRONICS:    {IconFn: a.FAS, Icon: "microchip"},
							types.TAG_TOOL_KICAD:     {IconFn: a.SimpleIcon, Filter: true},
							types.TAG_TOOL_FREECAD:   {IconFn: a.FAS, Icon: "cog"},
							types.TAG_TOOL_PHOTOSHOP: {IconFn: a.SimpleIcon, Filter: true, Icon: "adobephotoshop"},
						},
					}),
					m.SectionHeader(m.SectionHeaderProps{
						TagKey:    "kind-words",
						Text:      "Kind Words People Have Said About Me",
						Header:    g.H2,
						ClassList: []string{ch.MarginL("1rem")},
					}),
					a.AuthorQuoteText(
						"Matt displays the intelligence, creativity and ingenuity of a seasoned veteran, despite [being in] his first formal role in software engineering. His passion and determination to do [things] the right way shines in all that he does. That skill and style is exceeded only by his humanity. Pricey is simply one of the most kind, compassionate and generous humans I have ever known - an example to others and one that others thus try to live to. It was an honour to exist beside him in our team. If you have the fortune to get his attention, I strongly recommend him!",
						"Damian, AppSec Lead at Swyftx",
						t.Colors.Pallette.Green,
						[]string{},
					),
					a.AuthorQuoteText(
						"Matthew has demonstrated exceptional skills as a team player, working well with others to complete assigned projects to a high quality. Matthew would be a tremendous asset for your company and has my highest recommendation.",
						"Robert, Sessional Academic at QUT",
						t.Colors.Pallette.Cyan,
						[]string{},
					),
					a.AuthorQuoteText(
						"Matt has always worked above expectations as a part of our team [and] has always been one of the most pleasant individuals I've had under my employ.",
						"Dale, Store Manager at Dominos",
						t.Colors.Pallette.Blue,
						[]string{},
					),
					m.SectionHeader(m.SectionHeaderProps{
						TagKey:    "contact",
						Text:      "Getting in Touch",
						Header:    g.H2,
						ClassList: []string{ch.MarginL("1rem")},
					}),
					a.P(g.EB{
						Text:      "If you'd like to get in touch, my preferred method of contact is via the linkedin button below, or in the top right.",
						ClassList: []string{ch.MarginL("1rem")},
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
