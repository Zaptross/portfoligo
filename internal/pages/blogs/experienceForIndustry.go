package blogs

import (
	"time"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	c "github.com/zaptross/portfoligo/internal/components"
	m "github.com/zaptross/portfoligo/internal/components/molecules"
	"github.com/zaptross/portfoligo/internal/theme"
	"github.com/zaptross/portfoligo/internal/types"
)

var ExperienceForIndustry = types.PageDetails{
	Title:       "Experience for Industry",
	Description: "A blog post analysing and approaching the software and game development job markets for me as a student.",
	Slug:        "experience-for-industry",
	Type:        types.TYPE_BLOG,
	Written:     time.Date(2019, 8, 23, 0, 0, 0, 0, time.UTC),
	Tags: []string{
		types.TAG_STUDENT,
		types.TYPE_BLOG,
		types.TAG_WEB,
	},
	Content: func(p types.PageDetails) *g.HTMLElement {
		t := theme.UseTheme()
		fnc := m.NewFootnoteCollector()

		ue4Market := fnc.Add(m.ReferenceProps{
			Title:     "Content Search - UE4 Marketplace",
			Date:      time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			Retrieved: time.Date(2019, 8, 23, 0, 0, 0, 0, time.UTC),
			Link:      "https://www.unrealengine.com/marketplace/en-US/assets?keywords=feedback",
		})

		gameDevPD := fnc.Add(m.ReferenceProps{
			Title:     "Game Developer: Job Description, Duties and Requirements",
			Date:      time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			Retrieved: time.Date(2019, 8, 23, 0, 0, 0, 0, time.UTC),
			Link:      "https://study.com/articles/Game_Developer_Job_Description_Duties_and_Requirements.html",
		})

		softwareHowTo := fnc.Add(m.ReferenceProps{
			Title:     "How to become a Software Engineer",
			Date:      time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			Retrieved: time.Date(2019, 8, 23, 0, 0, 0, 0, time.UTC),
			Link:      "https://www.seek.com.au/career-guide/role/software-engineer?campaigncode=lrn:skj:sklm:cg:jbd:alpha",
		})

		gamingHowTo := fnc.Add(m.ReferenceProps{
			Title:     "How to Get a Job in Gaming",
			Date:      time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			Retrieved: time.Date(2019, 8, 23, 0, 0, 0, 0, time.UTC),
			Link:      "https://www.adzuna.com.au/blog/how-to-get-a-job-in/how-to-get-a-job-in-gaming/",
		})

		deptJobs := fnc.Add(m.ReferenceProps{
			Title:     "Australian Jobs 2019",
			Date:      time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			Retrieved: time.Date(2024, 7, 15, 0, 0, 0, 0, time.UTC),
			Link:      "https://www.adzuna.com.au/blog/how-to-get-a-job-in/how-to-get-a-job-in-gaming/",
		})

		headerClass := "header-img"
		g.Class(&g.CSSClass{
			Selector: "." + headerClass,
			Props: g.CSSProps{
				"max-height":              "300px",
				"aspect-ratio":            "auto",
				"overflow":                "hidden",
				"border-top-left-radius":  t.Radii.Default,
				"border-top-right-radius": t.Radii.Default,
			},
		})

		sectionOverrideClass := "section-header-override > div"
		g.Class(&g.CSSClass{
			Selector: "." + sectionOverrideClass,
			Props: g.CSSProps{
				"border-top-left-radius":  t.Radii.None,
				"border-top-right-radius": t.Radii.None,
				"margin-top":              "0",
				"margin-right":            "0",
			},
		})

		return g.Div(g.EB{
			Children: []*g.HTMLElement{
				c.Row(g.CE{
					c.Col(g.CE{
						g.Div(g.EB{
							ClassList: []string{headerClass},
							Children: g.CE{
								fnc.RefLink(c.ImgProject(p, "employment-growth-2019.png", []string{ch.W("100%")}), deptJobs),
							},
						}),
						c.SectionHeader(c.SectionHeaderProps{
							TagKey:    "experience",
							Text:      "The Experience",
							Header:    g.H3,
							ClassList: []string{sectionOverrideClass},
						}),
						c.P(g.EB{
							Text: "As the journey of completing a university degree comes to it's long awaited close, the time has come to broach the job market." +
								" Some job research later, and as discussed in more detail below, I found that both the games and software development industries" +
								" look favourably upon many transferable skills.",
						}),
						c.P(g.EB{
							Text: "From reading through these common requirements I noted a number of other skills which I did not have that were similarly transferable." +
								" After brainstorming for a while, I came up with an idea to for a new project which combines several skills into one manageable package.",
						}),
						c.P(g.EB{
							Text: "To further myself professionally, I will undertake a project that will teach me development skills for the Unreal Engine and C++," +
								" expand upon my user interface design experience, and teach me more about web services and API calls.",
						}),
					}, nil),
				}, nil),
				c.Row(g.CE{
					c.Col(g.CE{
						c.SectionHeader(c.SectionHeaderProps{
							TagKey: "industry",
							Text:   "The Industry",
							Header: g.H3,
						}),
						c.P(g.EB{Text: "The games and software development industries seek out similar candidates but with select different skills." +
							" Both industries frequently require employees with a few years of programming experience, understanding of algorithms, demonstrated teamwork," +
							" and independent development experiences.",
						}),
						c.P(g.EB{
							Children: g.CE{
								g.Text("Current games development opportunities tend to require Unity Engine experience, as it has come to be a widely spread engine for" +
									" small and medium sized games on a variety of platforms. These opportunities sometimes require graphics programming, user interface design," +
									" and artificial intelligence experience in addition to more general development skills."),
								fnc.Ref(gamingHowTo),
								g.Text("Many game developer job listings also seek out candidates who have Unreal Engine and C++ experience, which so far is outside of my skill set."),
								fnc.Ref(gameDevPD),
							},
						}),
						c.P(g.EB{
							Children: g.CE{
								g.Text("Software development roles look for candidates who have experience with Databases, web development skills such as JavaScript and HTML," +
									" Quality Assurance testing, documentation, and program maintenance skills."),
								fnc.Ref(softwareHowTo),
							},
						}),
						c.P(g.EB{
							Text: "Many of the listed requirements for such positions I have already met through my studies, such as Unity and C# development skills." +
								" I have demonstrated these through my various game projects, but there are still many ways I could expand upon them to further myself and my" +
								" portfolio in the eyes of potential employers.",
						}),
						c.P(g.EB{
							Text: "Some of these absent skills and experiences in my repertoire are web services and API's, C++, Unreal Engine, player feedback systems," +
								" and publishing works. These are the skills I have chosen to work on for my portfolio and professional development.",
						}),
					}, nil),
				}, nil),
				c.Row(g.CE{
					c.Col(g.CE{
						c.SectionHeader(c.SectionHeaderProps{
							TagKey: "project",
							Text:   "The Project",
							Header: g.H3,
						}),
						c.P(g.EB{
							Text: "To gain experience with my chosen skills I will develop an in-game player feedback system for the Unreal Engine which will collate and" +
								" post player feedback to a specified GitHub project board, then notify developers via Discord or Slack. This project when completed will be" +
								" published to the Unreal Marketplace.",
						}),
						c.P(g.EB{
							Children: g.CE{
								g.Text("The Unreal Engine is a prime choice for this project as in addition to being a desirable skill for games development, it also uses C++" +
									" as its primary language which is in demand in the software development industry. Some research on the Unreal Marketplace shows a gap in the" +
									" market for such a project as there is only one comparable piece of software currently available."),
								fnc.Ref(ue4Market),
							},
						}),
						c.P(g.EB{
							Text: "This tool will be comprised of back-end code and a set of customisable user interface elements in the Unreal Engine which allow the developer to select types of feedback they want to collect. Some example feedback types include: screenshots, player written comments, feedback types (bug, comment, etc.), self reported player moods, severity, system information, and game world location.",
						}),
						c.P(g.EB{
							Text: "These pieces of information are vital to accurately identifying and debugging various software projects. For example, players may find specific and rare locations where the object geometry of a level can allow the player to escape the intended play area, which developers may otherwise not find. By allowing players to quickly report these issues and collecting both screenshots and game world locations, developers can far more reliably identify where problems have occurred. Similarly, players reporting their mood at various points aids designers in identifying what is and isn't creating the player experience they want. This information is typically only gained from in-depth playtesting which takes time and resources to achieve.",
						}),
						c.P(g.EB{
							Text: "This feedback will be posted to a GitHub project board of the developer's choosing, where it will be sorted by feedback type and grouped with similar feedback. These reports will be posted by making calls through the GitHub API, and sorted by project boards containing developer set keywords. Sorting reports this way allows the developers to customisably and easily identify trends within the information they receive.",
						}),
						c.P(g.EB{
							Text: "The developers will be regularly notified about the number and types of reports that are posted to these project boards via Discord or Slack. These notifications will be sent through their respective API's whenever a chosen number of feedback reports have been appended to project boards. In addition to alerting the developers, these notifications will provide some minor analytics of the users recent feedback such as word and report type frequency, average player moods, recent report severity, and game world hot-spots among player reports.",
						}),
						c.P(g.EB{
							Text: "This tool aims to help developers further understand and track both the issues with their game, and what is working for their game. By developing such a tool, I am able to demonstrate key understanding of what information is required to efficiently maintain and improve both software and game projects. This project also allows me to ambitiously expand my skill set to skills I did not find time for during my degree and further catch the eye of potential employers.",
						}),
					}, nil),
				}, nil),
				c.Row(g.CE{
					c.Col(append(g.CE{
						c.SectionHeader(c.SectionHeaderProps{
							TagKey: "references",
							Text:   "References",
							Header: g.H3,
						}),
					},
						fnc.GenerateReferences()...,
					), []string{ch.MxW("100%")}),
				}, nil),
			}},
		)
	},
}
