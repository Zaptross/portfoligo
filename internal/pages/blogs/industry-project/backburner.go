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

var Backburner = types.PageDetails{
	Title:       "Industry Project: Backburner",
	Description: "Week five of working on the industry project implementing Redis, AWS S3, and more.",
	Slug:        "industry-project-backburner",
	Type:        types.TYPE_BLOG,
	Written:     time.Date(2019, 10, 21, 0, 0, 0, 0, time.UTC),
	Series:      types.SERIES_INDUSTRY_PROJECT,
	Tags: []string{
		types.TAG_STUDENT,
		types.TAG_WEB,
		types.TAG_DESIGN,
		types.TAG_TOOL_AWS,
	},
	Content: func(p types.PageDetails) *g.HTMLElement {
		t := theme.UseTheme()
		return g.Div(g.EB{
			Children: []*g.HTMLElement{
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Caption: "Burndown Chart for Week 6",
					Src:     files.IndustryProjectBackburner.Week6Burndown,
				}),
				a.P(g.EB{
					Text: "This week this project had to take a back seat to other projects with sooner deadlines, and I only accomplished part of what I had originally planned.",
				}),
				a.P(g.EB{
					Text: "This week, I finished implementing Redis for caching and Amazon's S3 for long-term and cross-instance storage. In addition, I added a " +
						"administration page with a login page to allow developers to have a simple centralised control system for their server.",
				}),
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Caption: "Demonstrating Caching, Long Term Storage, and Admin Page",
					Src:     files.IndustryProjectBackburner.CachingLtsLogin,
				}),
				m.SectionHeader(m.SectionHeaderProps{
					TagKey: "redis-and-s3",
					Text:   "Redis and S3",
					Header: g.H3,
				}),
				a.P(g.EB{
					Text: "This week I spend the majority of my time working on the Redis and S3 code to cache data locally and store data across instances in the " +
						"long term. During this process I again encountered race conditions, which slowed development. Both my Redis and S3 functions were returning " +
						"undefined objects instead of the data they were supposed to fetch.",
				}),
				a.P(g.EB{
					Children: g.CE{
						g.Text("I solved this issue by sourcing and implementing an alternative library for Redis: "),
						a.TextLink("Async-Redis", "https://www.npmjs.com/package/async-redis"),
						g.Text(" This library promisifies the functionality of the standard Redis library, such that it is usable with async/await."),
					},
				}),
				a.P(g.EB{
					Text: "With this sorted, I looked into different libraries to perform the same promisification of the AWS-SDK. After some time of looking, I was " +
						"not able to find a library to do this for me, so I experimented with restructuring my code. Shortly into restructuring, I discovered that the " +
						"AWS-SDK has promise methods for many of its functions, which allowed me to neatly solve the race condition issue with S3.",
				}),
				m.SectionHeader(m.SectionHeaderProps{
					TagKey: "admin-page",
					Text:   "Admin Page",
					Header: g.H3,
				}),
				a.P(g.EB{
					Text: "While working on the Redis and S3 code, I found myself wishing for a simpler way to run my testing code. I slowly came to the conclusion " +
						"that a simple Web page served by the server might be the best option to fill that need.",
				}),
				a.P(g.EB{
					Text: "This page setup was relatively quick and easy to implement, as I was able to repurpose work from other projects. I added a page to ask for " +
						"login details, and check those details on the server side before sending a response to the browser. When polished, this will act as an " +
						"administrative tool for developers to manage the server from a user-interface instead of a command line.",
				}),
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Caption: "Demo of Admin Page",
					Src:     files.IndustryProjectBackburner.AdminPage,
				}),
				m.SectionHeader(m.SectionHeaderProps{
					TagKey: "until-submission",
					Text:   "Until Submission",
					Header: g.H3,
				}),
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Caption: "Burndown Plan for Week 7",
					Src:     files.IndustryProjectBackburner.Week7BurndownPlan,
				}),
				a.AuthorQuote(
					g.Text("The plan for the next week includes finishing the GitHub implementation, adding the Discord API calls, then writing the C++ calls to "+
						"contact this server from the Feedback Panel. Beyond this point, I plan to move on to polishing the Feedback Panel with some textures and "+
						"animations, and documenting the installation procedures for both the server and the panel. Ideally, I will end this next week with all necessary "+
						"features implemented and polished."),
					a.TextLink("Industry Project: Server", Server.GetRelativeURL()),
					t.Colors.Pallette.Green,
					nil,
				),
				a.P(g.EB{
					Text: "For the remaining time until submission, I will be undertaking the same tasks as last week, as I was not able to see them to completion within the week.",
				}),
			}},
		)
	},
}
