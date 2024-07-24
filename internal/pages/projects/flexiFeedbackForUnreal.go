package projects

import (
	"time"

	g "github.com/zaptross/gorgeous"
	"github.com/zaptross/portfoligo/generated/files"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	m "github.com/zaptross/portfoligo/internal/components/molecules"
	industryproject "github.com/zaptross/portfoligo/internal/pages/blogs/industry-project"
	"github.com/zaptross/portfoligo/internal/types"
)

var FlexiFeedbackForUnreal = types.PageDetails{
	Title:       "FlexiFeedback for Unreal",
	Description: "FlexiFeedback for Unreal is an in-game feedback panel for Unreal Engine 4 that allows developers to collect feedback from players.",
	Slug:        "flexi-feedback-for-unreal",
	Type:        types.TYPE_PROJECT,
	Written:     time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
	Tags: []string{
		types.TAG_STUDENT,
		types.TAG_GAME,
		types.TAG_WEB,
		types.TAG_DESIGN,
		types.TAG_TOOL_UNREAL,
		types.TAG_TOOL_AWS,
	},
	Content: func(p types.PageDetails) *g.HTMLElement {
		return g.Div(g.EB{
			Children: []*g.HTMLElement{
				a.YoutubeEmbed("FlexiFeedback for Unreal Demonstration Video", "https://www.youtube-nocookie.com/embed/br8MqCWPmtk"),
				a.P(g.EB{Children: g.CE{
					g.Text("FlexiFeedback For Unreal is an in-game feedback panel that allows players to send reports of many kinds from the game to the developers "),
					a.TextLink("GitHub Project Board", "https://help.github.com/en/github/managing-your-work-on-github/about-project-boards"),
					g.Text(". This project was designed to expand upon my skills as a developer and gain industry relevant experience."),
				}}),
				m.R2C(
					nil, g.CE{
						a.P(g.EB{
							Text: "FlexiFeedback is comprised of three major parts: the in-game feedback panel, an administrative web page, and a back-end server which" +
								" runs the page and handles feedback reports.",
						}),
						a.P(g.EB{
							Text: "The feedback panel gathers information directly reported by the player such as what type of report they are making, how they feel about" +
								" the issue, any comments they have, and optionally information about their gaming system.",
						}),
					},
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    industryproject.Redesign,
							Caption: "Feedback Panel Layout",
							Src:     files.IndustryProjectRedesign.UnrealAllUiImplemented,
						}),
					},
					nil,
				),
				m.R2C(
					nil, g.CE{
						a.P(g.EB{
							Text: "When players attempt to send partially complete information, or if an error occurs while attempting to send that information, the " +
								"feedback panel will inform them of the issue before allowing them to move on.",
						}),
						a.P(g.EB{
							Text: "When feedback is completed it is then sent to the back-end server for processing.",
						}),
					},
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    p,
							Caption: "Submit Warning Dialog",
							Src:     files.FlexiFeedbackForUnreal.SubmitWarningDialog,
						}),
					},
					nil,
				),
				m.R2C(
					nil, g.CE{
						a.P(g.EB{
							Children: g.CE{
								g.Text("The back-end server for FlexiFeedback is written in JavaScript, and accepts feedback reports formatted as JSON objects. This data" +
									" is then processed before being cached in a locally run "),
								a.TextLink("Redis", "https://redis.io/"),
								g.Text(" instance and stored in an "),
								a.TextLink("S3 Bucket.", "https://aws.amazon.com/s3/"),
							},
						}),
						a.P(g.EB{
							Text: "Finally, these reports are posted to the appropriate columns on a GitHub’s Project Board, alongside regularly updated summaries of each column.",
						}),
					},
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    industryproject.Recap,
							Caption: "Blueprint Node of HTTP Request",
							Src:     files.IndustryProjectRecap.UnrealBlueprintHttpNode,
						}),
					},
					nil,
				),
				m.R2C(
					nil, g.CE{
						a.P(g.EB{
							Text: "The last part of FlexiFeedback is an administrator's page hosted on the back-end server which regularly retrieves summary information " +
								"from the server, and provides them with an easy automated way to set-up their project board for FlexiFeedback use.",
						}),
					},
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    industryproject.Recap,
							Caption: "Admin Page",
							Src:     files.IndustryProjectRecap.AdminPage,
						}),
					},
					nil,
				),
				m.SectionHeader(m.SectionHeaderProps{
					Text:   "Where do I find FlexiFeedback?",
					TagKey: "where-to-find",
					Header: g.H4,
				}),
				a.P(g.EB{Children: g.CE{
					g.Text("You can read more about FlexiFeedback's development in my blog post "),
					a.TextLink("Industry Project: Recap", industryproject.Recap.GetRelativeURL()),
					g.Text(" and more about the design and building process in "),
					a.TextLink("my blog posts here.", "/search/?q="+types.SERIES_INDUSTRY_PROJECT),
				}}),
				a.P(g.EB{
					Text: "As of the time of writing, FlexiFeedback is unavailable for download or purchase, as it is still in active development. Some planned features" +
						" include: more developer tools on the administration page, sending and storing a screenshot of the player's view when they send a report, Discord " +
						"and other application integrations, and a Unity Engine release.",
				}),
			}},
		)
	},
}
