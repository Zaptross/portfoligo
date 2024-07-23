package industryproject

import (
	"time"

	g "github.com/zaptross/gorgeous"
	"github.com/zaptross/portfoligo/generated/files"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	m "github.com/zaptross/portfoligo/internal/components/molecules"
	"github.com/zaptross/portfoligo/internal/theme"
	"github.com/zaptross/portfoligo/internal/types"
)

var Server = types.PageDetails{
	Title:       "Industry Project: Server",
	Description: "The fourth week of working on the industry project, spent integrating with Github and Redis.",
	Slug:        "industry-project-server",
	Type:        types.TYPE_BLOG,
	Written:     time.Date(2019, 10, 14, 0, 0, 0, 0, time.UTC),
	Series:      types.SERIES_INDUSTRY_PROJECT,
	Tags: []string{
		types.TAG_STUDENT,
		types.TAG_WEB,
		types.TAG_DESIGN,
	},
	Content: func(p types.PageDetails) *g.HTMLElement {
		t := theme.UseTheme()
		return g.Div(g.EB{
			Children: []*g.HTMLElement{
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Caption: "Burndown Chart for Week 5",
					Src:     files.IndustryProjectServer.Week5Burndown,
				}),
				a.P(g.EB{
					Text: "This last week I spent working on the back-end server as described in last week's update. This mostly went well, but is taking longer than expected.",
				}),
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Caption: "Working Feedback Submission",
					Src:     files.IndustryProjectServer.FeedbackSubmissionWorking,
				}),
				m.SectionHeader(m.SectionHeaderProps{
					TagKey: "what-made-it-in",
					Text:   "What made it in?",
					Header: g.H3,
				}),
				a.AuthorQuote(
					g.Text("For next week, I am planning to build the back-end server and test the GitHub and Discord API calls through it, write and test calls to the "+
						"back-end server in C++, and link those calls to the Feedback Panel in the Unreal Engine."),
					a.TextLink("Industry Project: Redesign", Redesign.GetRelativeURL()),
					t.Colors.Pallette.Green,
					nil,
				),
				a.P(g.EB{Children: g.CE{
					g.Text("So far of my previous list, GitHub API through "),
					a.TextLink("Octokit REST", "https://www.npmjs.com/package/@octokit/rest"),
					g.Text(" has been mostly implemented (as seen in the demonstration video above). All of the necessary API calls have been written and are in the" +
						" process of being logically structured and tested. While writing and testing these calls, it became clear that a setup command and a local data" +
						" cache would become necessary very quickly. To that end, I have begun implementing "),
					a.TextLink("Redis", "https://www.npmjs.com/package/redis"),
					g.Text(" as a caching service, and writing a server endpoint to set up a project board with minimal developer work."),
				}}),
				m.SectionHeader(m.SectionHeaderProps{
					TagKey: "whats-not-in",
					Text:   "What's not in?",
					Header: g.H3,
				}),
				a.P(g.EB{
					Text: "From the above quote, Discord's API calls have not yet been implemented but have begun to be investigated, and the C++ code" +
						" necessary to link the Unreal Feedback Panel to the back-end server has not yet been written either.",
				}),
				m.SectionHeader(m.SectionHeaderProps{
					TagKey: "next-week",
					Text:   "Plan for next week",
					Header: g.H3,
				}),
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Caption: "Burndown Chart for Week 6",
					Src:     files.IndustryProjectServer.Week6BurndownPlan,
				}),
				a.P(g.EB{
					Text: "The plan for the next week includes finishing the GitHub implementation, adding the Discord API calls, then writing the C++ calls" +
						" to contact this server from the Feedback Panel. Beyond this point, I plan to move on to polishing the Feedback Panel with some textures and " +
						"animations, and documenting the installation procedures for both the server and the panel. Ideally, I will end this next week with all " +
						"necessary features implemented and polished.",
				}),
			}},
		)
	},
}
