package pages

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	m "github.com/zaptross/portfoligo/internal/components/molecules"
	o "github.com/zaptross/portfoligo/internal/components/organisms"
	"github.com/zaptross/portfoligo/internal/components/specific"
	industryproject "github.com/zaptross/portfoligo/internal/pages/blogs/industry-project"
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
					specific.AboutBlock(),
					homePageWorkingInSoftwareBlock(),
					homePageProfessionalExperienceSkillGrid(),
					homePageSeeAboutPageCallout(),
					m.SectionHeader(m.SectionHeaderProps{
						TagKey: "favourites",
						Header: g.H3,
						Text:   "A few of my favourites projects and posts, and why I love them:",
					}),
					o.PreviewExplained("My work on FlexiFeedback for Unreal was the project that set me on my current career path, and I am proud of my younger self's direction with the project and view to the future.", industryproject.Recap),
					o.PreviewExplained("This was my first game jam entry, made with a friend and we had a blast! It's left me with an ongoing desire to participate in more jams!", projects.NoRhythm),
					o.PreviewExplained("To this day I am especially proud of the slum and office tiles I created for Kowloon Exigency as a display of rich level design.", projects.KowloonExigency),
				}},
			)
		},
	})
)

func homePageSeeAboutPageCallout() *g.HTMLElement {
	return a.P(g.EB{
		Children: g.CE{
			g.Span(g.EB{Text: "See my "}),
			a.TextLink("about page", "/about"),
			g.Span(g.EB{Text: " for a more detailed breakdown of my skills."}),
		},
	})
}

func homePageProfessionalExperienceSkillGrid() *g.HTMLElement {
	skills := []string{
		types.TAG_LANG_TS,
		types.TAG_LANG_REACT,
		types.TAG_LANG_CSHARP,
		types.TAG_TOOL_DOTNET,
		types.TAG_TOOL_JAVA,
		types.TAG_TOOL_PYTHON,
		types.TAG_TOOL_AWS,
		types.TAG_TOOL_DATADOG,
	}

	return o.SkillsGrid(o.SkillsGridProps{
		Skills:   skills,
		MapIcons: o.SkillsMappingDefaults(),
		Style: g.CSSProps{
			"grid-template-columns": "repeat(auto-fit, minmax(35px, 1fr))",
		},
	})
}

func homePageWorkingInSoftwareBlock() *g.HTMLElement {
	text, script := specific.YearsInSoftware()
	return a.P(g.EB{
		Children: g.CE{
			g.Span(g.EB{
				Text:   fmt.Sprintf("I have been working in software for %s and I have experience with a variety of languages and frameworks, including professional experience with: ", text),
				Script: script,
			}),
		},
	})
}
