package pages

import (
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	c "github.com/zaptross/portfoligo/internal/components"
	p "github.com/zaptross/portfoligo/internal/provider"
	"github.com/zaptross/portfoligo/internal/types"
)

var (
	_ = registerPage(types.PageDetails{
		Title:       "About",
		Description: "About Me",
		Slug:        "about",
		Type:        types.TYPE_ROOT,
		Content: func(_ types.PageDetails) *g.HTMLElement {
			t := p.ThemeProvider.GetTheme()
			return g.Div(g.EB{
				ClassList: []string{ch.FlexCol(), ch.JustifyContent(ch.Content.SpaceBetween)},
				Children: []*g.HTMLElement{
					c.SectionHeader(c.SectionHeaderProps{
						TagKey:    "about",
						Text:      "Hi 👋, I'm Matt.",
						Header:    g.H1,
						Color:     t.Yellow,
						ClassList: []string{ch.MarginL("1rem")},
					}),
					c.P(g.EB{
						Text:      "I am a full-stack polyglot software engineer with a passion for learning new skills and technologies.",
						ClassList: []string{ch.Margin("0.25rem 1rem")},
					}),
					c.P(g.EB{
						Text:      "In my spare time I love to play videogames and boardgames, develop fun ideas and projects, and work on hobbyist electronics projects.",
						ClassList: []string{ch.Margin("0.25rem 1rem")},
					}),
					c.SectionHeader(c.SectionHeaderProps{
						TagKey:    "skills",
						Text:      "Skills",
						Header:    g.H2,
						Color:     t.Yellow,
						ClassList: []string{ch.MarginL("1rem")},
					}),
					c.P(g.EB{
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
							types.TAG_TOOL_KUBE:                       {IconFn: c.SimpleIcon},
							types.TAG_LANG_TS:                         {IconFn: c.SimpleIcon},
							types.TAG_LANG_CSHARP:                     {IconFn: c.SimpleIcon},
							types.TAG_TOOL_UNREAL:                     {IconFn: c.SimpleIcon},
							types.TAG_TOOL_DOTNET:                     {IconFn: c.SimpleIcon},
							(types.TAG_GAME + " " + types.TAG_DESIGN): {IconFn: c.FAS, Icon: "gamepad"},
						},
					}),
					c.P(g.EB{
						Text:      "And hobbyist experience with:",
						ClassList: []string{ch.MarginL("1rem")},
					}),
					c.SkillsGrid(c.SkillsGridProps{
						Skills: []string{types.TAG_ELECTRONICS, types.TAG_TOOL_KICAD, types.TAG_TOOL_FREECAD, types.TAG_TOOL_PHOTOSHOP},
						MapIcons: map[string]c.SkillMapping{
							types.TAG_ELECTRONICS:    {IconFn: c.FAS, Icon: "microchip"},
							types.TAG_TOOL_KICAD:     {IconFn: c.SimpleIcon},
							types.TAG_TOOL_FREECAD:   {IconFn: c.FAS, Icon: "cog"},
							types.TAG_TOOL_PHOTOSHOP: {IconFn: c.SimpleIcon, Icon: "adobephotoshop"},
						},
					}),
					c.SectionHeader(c.SectionHeaderProps{
						TagKey:    "kind-words",
						Text:      "Kind Words People Have Said About Me",
						Header:    g.H2,
						Color:     t.Yellow,
						ClassList: []string{ch.MarginL("1rem")},
					}),
					c.AuthorQuote(
						"Matt displays the intelligence, creativity and ingenuity of a seasoned veteran, despite [being in] his first formal role in software engineering. His passion and determination to do [things] the right way shines in all that he does. That skill and style is exceeded only by his humanity. Pricey is simply one of the most kind, compassionate and generous humans I have ever known - an example to others and one that others thus try to live to. It was an honour to exist beside him in our team. If you have the fortune to get his attention, I strongly recommend him!",
						"Damian, AppSec Lead at Swyftx",
						t.Green,
						[]string{},
					),
					c.AuthorQuote(
						"Matthew has demonstrated exceptional skills as a team player, working well with others to complete assigned projects to a high quality. Matthew would be a tremendous asset for your company and has my highest recommendation.",
						"Robert, Sessional Academic at QUT",
						t.Cyan,
						[]string{},
					),
					c.AuthorQuote(
						"Matt has always worked above expectations as a part of our team [and] has always been one of the most pleasant individuals I've had under my employ.",
						"Dale, Store Manager at Dominos",
						t.Blue,
						[]string{},
					),
					c.SectionHeader(c.SectionHeaderProps{
						TagKey:    "contact",
						Text:      "Getting in Touch",
						Header:    g.H2,
						Color:     t.Yellow,
						ClassList: []string{ch.MarginL("1rem")},
					}),
					c.P(g.EB{
						Text:      "If you'd like to get in touch, my preferred method of contact is via the linkedin button below, or in the top right.",
						ClassList: []string{ch.MarginL("1rem")},
					}),
					c.Row(g.CE{
						c.LinkIcon(
							c.FAB("linkedin", g.CSSProps{"font-size": "2rem"}),
							"https://www.linkedin.com/in/mpdd/",
						),
					}, []string{ch.MarginZA()}),
				}},
			)
		},
	})
)
