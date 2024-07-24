package industryproject

import (
	"time"

	g "github.com/zaptross/gorgeous"
	"github.com/zaptross/portfoligo/generated/files"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	m "github.com/zaptross/portfoligo/internal/components/molecules"
	"github.com/zaptross/portfoligo/internal/types"
)

var Recap = types.PageDetails{
	Title:       "Industry Project: Recap",
	Description: "Looking back over the industry project, what went well and what could have gone better.",
	Slug:        "industry-project-recap",
	Type:        types.TYPE_BLOG,
	Written:     time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
	Series:      types.SERIES_INDUSTRY_PROJECT,
	Tags: []string{
		types.TAG_STUDENT,
		types.TAG_WEB,
		types.TAG_TOOL_UNREAL,
		types.TAG_DESIGN,
		types.TAG_TOOL_AWS,
	},
	Content: func(p types.PageDetails) *g.HTMLElement {
		// t := theme.UseTheme()
		fnc := m.NewFootnoteCollector()

		softwareJobs := fnc.Add(a.ReferenceProps{
			Title:     "Software Developer C++ Search",
			Date:      time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
			Retrieved: time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
			Link:      "https://www.seek.com.au/software-developer-C%2B%2B-jobs",
		})

		gameDesignerJobs := fnc.Add(a.ReferenceProps{
			Title:     "Game designer Jobs",
			Date:      time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
			Retrieved: time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
			Link:      "https://www.glassdoor.com.au/Job/game-designer-jobs-SRCH_KO0,13.htm",
		})

		softwareHowTo := fnc.Add(a.ReferenceProps{
			Title:     "How to become a Software Engineer",
			Date:      time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
			Retrieved: time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
			Link:      "https://www.seek.com.au/career-guide/role/software-engineer?campaigncode=lrn:skj:sklm:cg:jbd:alpha",
		})

		fnc.Add(a.ReferenceProps{
			Title:     "Top skills for software engineers",
			Date:      time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
			Retrieved: time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
			Link:      "https://www.monster.com/career-advice/article/software-engineer-skills",
		})

		mitLicence := fnc.Add(a.ReferenceProps{
			Title:     "The MIT License | Open Source Initiative",
			Date:      time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
			Retrieved: time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
			Link:      "https://opensource.org/licenses/MIT",
		})

		fnc.Add(a.ReferenceProps{
			Title:     "Content Search – UE4 Marketplace",
			Date:      time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
			Retrieved: time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
			Link:      "https://www.unrealengine.com/marketplace/en-US/assets?keywords=feedback",
		})

		fnc.Add(a.ReferenceProps{
			Title:     "Unity Asset Store – Feedback Content Search",
			Date:      time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
			Retrieved: time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
			Link:      "https://assetstore.unity.com/?q=feedback&orderBy=0",
		})

		fnc.Add(a.ReferenceProps{
			Title:     "Game Developer: Job Description, Duties and Requirements",
			Date:      time.Date(2019, 8, 23, 0, 0, 0, 0, time.UTC),
			Retrieved: time.Date(2019, 8, 23, 0, 0, 0, 0, time.UTC),
			Link:      "https://study.com/articles/Game_Developer_Job_Description_Duties_and_Requirements.html",
		})

		fnc.Add(a.ReferenceProps{
			Title:     "How to Get a Job in Gaming",
			Date:      time.Date(2019, 8, 23, 0, 0, 0, 0, time.UTC),
			Retrieved: time.Date(2019, 8, 23, 0, 0, 0, 0, time.UTC),
			Link:      "https://www.adzuna.com.au/blog/how-to-get-a-job-in/how-to-get-a-job-in-gaming/",
		})

		fnc.Add(a.ReferenceProps{
			Title:     "Australian Jobs 2019",
			Date:      time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			Retrieved: time.Date(2024, 7, 15, 0, 0, 0, 0, time.UTC),
			Link:      "https://www.professions.org.au/wp-content/uploads/Australian-Jobs-2019_AustGovt_DJSB.pdf",
		})

		return g.Div(g.EB{
			Children: []*g.HTMLElement{
				m.SectionHeader(m.SectionHeaderProps{
					TagKey: "project",
					Text:   "The Project",
					Header: g.H3,
				}),
				a.P(g.EB{Text: "The development process of FlexiFeedback has been a thoroughly enjoyable one. Though it had it's fair share of troubles and has not" +
					" yet met all of its outlined goals, I am proud of what I have accomplished."}),
				m.R2C(
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    Beginnings,
							Caption: "Mockup Layout",
							Src:     files.IndustryProjectBeginnings.ProtoLayout1,
						}),
					},
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    Redesign,
							Caption: "Current Layout",
							Src:     files.IndustryProjectRedesign.UnrealAllUiImplemented,
						}),
					},
					nil,
				),
				a.P(g.EB{Text: "Though only functional and not polished, all the elements of my proposed user interface have been implemented and tested. Process" +
					" of creating these elements worked towards the next intended outcomes of the project: gaining experience in the Unreal Engine, blueprints, and C++."}),
				m.R3C(
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    Redesign,
							Caption: "Checkbox Implementation",
							Src:     files.IndustryProjectRedesign.UnrealBlueprintCheckboxes,
						}),
					},
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    p,
							Caption: "C++ HTTP Post Request",
							Src:     files.IndustryProjectRecap.HttpPostRequest,
						}),
					},
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    p,
							Caption: "Blueprint Node of C++ Function",
							Src:     files.IndustryProjectRecap.UnrealBlueprintHttpNode,
						}),
					},
					nil,
				),
				a.P(g.EB{Text: "While originally designed to make a variety of API requests directly from the game, I came to the conclusion that reworking the" +
					" project to include a back-end server would instead speed up development."}),
				m.R2C(
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    p,
							Caption: "JavaScript Handler for Feedback",
							Src:     files.IndustryProjectRecap.JsFeedbackHandler,
						}),
					},
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    Redesign,
							Caption: "Server Design",
							Src:     files.IndustryProjectRedesign.ServerDesign,
						}),
					},
					nil,
				),
				a.P(g.EB{
					Children: g.CE{
						g.Text("This redesign led to exploring a variety of MIT licenced"),
						fnc.Ref(mitLicence),
						g.Text(" libraries for JavaScript, which did significantly speed up development. This server allowed me to achieve my goal of posting these reports to GitHub's Project Boards, and inspired the creation of an administration page for developers."),
					},
				}),
				m.R2C(
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    p,
							Caption: "Creating New Project Board Columns",
							Src:     files.IndustryProjectRecap.CreatingProjectboardColumns,
						})},
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    p,
							Caption: "Administration Page",
							Src:     files.IndustryProjectRecap.AdminPage,
						})},
					nil,
				),
				a.P(g.EB{Text: "Through the creation and implementation of this server, and the C++ code I had written for the feedback panel, I was finally able to" +
					" achieve my primary goal: I was able to send feedback from the game to the project board."}),
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Caption: "Full Demo",
					Src:     files.IndustryProjectRecap.FullDemo,
				}),
				a.P(g.EB{
					Text: "After this was successfully implemented the time had come to submit my project, and thus I wasn't able to complete all of my outlined" +
						" goals. In the allotted time I was not able to implement Discord or Slack summary posts as I had intended, nor was I able to produce a " +
						"fully polished version of the user interface.",
				}),
				a.P(g.EB{
					Text: "Due to these missing features, I chose not to submit this application to the Unreal Marketplace as intended. In it's current state it" +
						" could be published as it does successfully perform its core functionality, but it is lacking adequate depth of features and its unpolished " +
						"user interface would be unappealing to is intended market.",
				}),
				m.SectionHeader(m.SectionHeaderProps{
					TagKey: "industry",
					Text:   "The Industry",
					Header: g.H3,
				}),
				a.P(g.EB{
					Text: "The original intent of this project was to gain industry relevant experience with new tools to further myself professionally and become " +
						"a more desirable applicant. To achieve this, I had decided to create this project and gain experience with the Unreal Engine, C++, and Blueprints.",
				}),
				a.P(g.EB{Children: g.CE{
					g.Text("These tools are still in demand in their respective industries with hundreds of results when searching for C++ and software development jobs"),
					fnc.Ref(softwareJobs),
					g.Text(", and hundreds more results when searching for game designer jobs."),
					fnc.Ref(gameDesignerJobs),
				}}),
				a.P(g.EB{Children: g.CE{
					g.Text("These tools are still in demand in their respective industries with hundreds of results when searching for C++ and software development jobs"),
					fnc.Ref(softwareHowTo),
					g.Text(". The redesign also demonstrates my willingness to learn new techniques and use varied technologies to achieve my goals, which will " +
						"endear me to potential employers looking for adaptable employees."),
				}}),
				a.P(g.EB{
					Text: "I believe that this project positions me well for future interviews in the games industry by demonstrating key skills and understanding" +
						" of development with the Unreal Engine, blueprints, C++. Furthermore, it has made me a more desirable candidate for employment within the software " +
						"development industry by expanding upon my web development skills and allowing me to demonstrate adaptability with new technologies.",
				}),
				m.SectionHeader(m.SectionHeaderProps{
					TagKey: "references",
					Text:   "References",
					Header: g.H3,
				}),
				a.Col(fnc.GenerateReferences(), nil),
			}},
		)
	},
}
