package industryproject

import (
	"time"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	c "github.com/zaptross/portfoligo/internal/components"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	m "github.com/zaptross/portfoligo/internal/components/molecules"
	"github.com/zaptross/portfoligo/internal/generated/files"
	"github.com/zaptross/portfoligo/internal/types"
)

var Redesign = types.PageDetails{
	Title:       "Industry Project: Redesign",
	Description: "The second week of working on the industry project, spent debugging the screenshot and redesigning.",
	Slug:        "industry-project-redesign",
	Type:        types.TYPE_PROJECT,
	Written:     time.Date(2019, 10, 9, 0, 0, 0, 0, time.UTC),
	Series:      types.SERIES_INDUSTRY_PROJECT,
	Tags: []string{
		types.TAG_STUDENT,
		types.TAG_TOOL_UNREAL,
		types.TAG_WEB,
		types.TAG_DESIGN,
	},
	Content: func(p types.PageDetails) *g.HTMLElement {
		return g.Div(g.EB{
			Children: []*g.HTMLElement{
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Caption: "Burndown Chart for Week 3-4",
					Src:     files.IndustryProjectRedesign.Week34BurndownRedesign,
					ContainerClasses: []string{
						ch.MxH("fit-content"),
					},
				}),
				a.P(g.EB{
					Text: "As the chart above shows, this week I worked for less hours than planned but achieved significantly less progress than intended. This lack of" +
						" progress is due to unforeseen issues with making API calls in C++, and the amount of time spent tackling the bug of screenshots not displaying.",
				}),
				a.P(g.EB{
					Text: "The images below show the final state of the project at the end of these two weeks: all proposed user interface elements have been" +
						" successfully implemented, and a new solution for managing the API calls has been designed.",
				}),
				c.Row(g.CE{
					m.CaptionedImage(m.CaptionedImageProps{
						Page:    p,
						Caption: "All UI Implemented",
						Src:     files.IndustryProjectRedesign.UnrealAllUiImplemented,
						ContainerClasses: []string{
							ch.MxH("fit-content"),
						},
					}),
					m.CaptionedImage(m.CaptionedImageProps{
						Page:    p,
						Caption: "Server Design",
						Src:     files.IndustryProjectRedesign.ServerDesign,
						ContainerClasses: []string{
							ch.MxH("fit-content"),
						},
					}),
				}, nil),
				c.SectionHeader(c.SectionHeaderProps{
					Text:   "Sentiment Radio-Buttons",
					TagKey: "sentiment",
					Header: g.H4,
				}),
				a.P(g.EB{
					Text: "Radio Buttons proved to be more complex than I had originally anticipated, as they are not a default UI feature within the Unreal Engine." +
						" So instead, I constructed my own radio-buttons using check boxes and a surprising amount of blueprinting. The final implementation below will" +
						" report the user's sentiment in a range of -2 to 2, as depicted by emotive faces in the mock-up.",
				}),
				c.Row(g.CE{
					m.CaptionedImage(m.CaptionedImageProps{
						Page:    p,
						Caption: "Mockup Layout",
						Src:     files.IndustryProjectRedesign.ProtoLayout3,
					}),
					m.CaptionedImage(m.CaptionedImageProps{
						Page:    p,
						Caption: "Checkbox Implementation",
						Src:     files.IndustryProjectRedesign.UnrealBlueprintCheckboxes,
					}),
				}, nil),
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Caption: "Checkboxes in Action",
					Src:     files.IndustryProjectRedesign.UnrealCheckboxes,
				}),
				c.SectionHeader(c.SectionHeaderProps{
					TagKey: "screenshot",
					Text:   "Getting the Screenshot Working",
					Header: g.H4,
				}),
				a.P(g.EB{
					Text: "This was to be the first of two unfortunate hurdles: the screenshot wasn't appearing on the UI but was demonstrably updating correctly on the" +
						" debugging block in the level. I spent a significant amount of time scouring the internet for answers and experimenting with different" +
						" variations on the blueprints, before finally arriving at the solution.",
				}),
				c.Row(g.CE{
					m.CaptionedImage(m.CaptionedImageProps{
						Page:             p,
						Caption:          "Before Capture",
						Src:              files.IndustryProjectRedesign.UnrealBeforeCapture,
						ContainerClasses: []string{ch.Flex(1)},
					}),
					m.CaptionedImage(m.CaptionedImageProps{
						Page:             p,
						Caption:          "After Capture",
						Src:              files.IndustryProjectRedesign.UnrealAfterCapture,
						ContainerClasses: []string{ch.Flex(1)},
					}),
				}, nil),
				a.P(g.EB{
					Text: "The solution was something that I had experimented with as part of the problem solving process, but hadn't seen due to inexperience with" +
						" the Unreal Engine: the UI can't render a render texture, but it can render a UI material. With that revelation, it became clear that " +
						"the rest of my blueprint was working!",
				}),
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Caption: "Screenshot Displaying in UI",
					Src:     files.IndustryProjectRedesign.UnrealScreenshotInUi,
				}),
				c.SectionHeader(c.SectionHeaderProps{
					TagKey: "c-plus-plus",
					Text:   "C++ and PC Info",
					Header: g.H4,
				}),
				a.P(g.EB{
					Text: "The last element of my proposed UI design was retrieving and optionally sending some PC info to help developers. To do this, I started down" +
						" the path of learning C++ to write my own blueprint functions.",
				}),
				a.P(g.EB{
					Text: "Thankfully, C++ has some built-in functions for retrieving PC details, so my first foray into C++ was mostly learning from examples.",
				}),
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Src:     files.IndustryProjectRedesign.UnrealBlueprintGetPcInfo,
					Caption: "PC Details Blueprint",
				}),
				a.P(g.EB{
					Text: "Thus, as seen below, my proposed UI was all implemented and functioning as intended.",
				}),
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Src:     files.IndustryProjectRedesign.UnrealAllUiImplemented,
					Caption: "All UI Implemented",
				}),
				c.SectionHeader(c.SectionHeaderProps{
					TagKey: "c-plus-plus-api",
					Text:   "C++ and API Research",
					Header: g.H4,
				}),
				a.P(g.EB{
					Children: g.CE{
						g.Text("The last part of my week was to explore my chosen API's and begin creating API calls in C++ to be executed from inside the game." +
							" This quickly ground to a halt as it became clear that making these API calls wouldn't take the form that I was familiar with. In other projects" +
							" I have primarily used URL based calls, like those of the "),
						a.TextLink("YouTube API", "https://developers.google.com/youtube/v3/docs/channels/list"),
						g.Text(". The GitHub and Imgur APIs are both written primarily for use with "),
						a.TextLink("CURL", "https://en.wikipedia.org/wiki/CURL"),
						g.Text(" which I am not familiar with using in code."),
					},
				}),
				a.P(g.EB{
					Children: g.CE{
						g.Text("So, I turned to finding a C++ library which could be used to make these calls more easily. While looking for GitHub "),
						a.TextLink("libraries", "https://developer.github.com/v3/libraries/"),
						g.Text(" I discovered that there were many language options, but there was not a library for C++. Instead, I began to look for a CURL library for C++ and found "),
						a.TextLink("curlPP", "http://www.curlpp.org/"),
						g.Text(" which seemed to be what I needed."),
					},
				}),
				a.P(g.EB{
					Children: g.CE{
						g.Text("I downloaded and attempted to install curlPP by bringing it's code into the source files of my Unreal project, and found that the Unreal" +
							" Engine didn't recognise them. After some searching online, I found a "),
						a.TextLink("few", "https://forums.unrealengine.com/development-discussion/c-gameplay-programming/28934-can-t-include-custom-library"),
						a.TextLink("answer", "https://wiki.unrealengine.com/Linking_Static_Libraries_Using_The_Build_System"),
						a.TextLink("posts", "https://answers.unrealengine.com/questions/218616/how-do-i-add-thirdparty-library.html"),
						g.Text(" detailing ways to include libraries into Unreal projects. While following these instructions, I found myself continually hitting build" +
							" errors in the project. I attempted to separately build the curlPP files to include them, on the recommendation of a different answer post, to no avail."),
					},
				}),
				a.P(g.EB{
					Text: "Finally, I attempted to find examples of creating HTTP header requests from the ground up in C++. Many of these examples suggested using a" +
						" library first, but I found some which had example code written with only the standard C++ libraries. While attempting to implement some of their" +
						" code I found that I continued to run into build errors, leading to the realisation that writing this all in C++ manually would be " +
						"time-inefficient and continue to hold back development.",
				}),
				c.SectionHeader(c.SectionHeaderProps{
					TagKey: "redesign",
					Text:   "Redesign: Server",
					Header: g.H4,
				}),
				a.P(g.EB{
					Text: "After discussing the C++ issues with a friend, I came to the conclusion that a slight rework of my original idea would cut down on the" +
						" problems I was facing and speed up development. Originally, it was intended for the game to GET and POST to the different API's directly to cut-down on developer setup time.",
				}),
				a.P(g.EB{
					Text: "The approach I now intend to take is to create a back-end web-server written in JavaScript which will handle the various API calls, and " +
						"cut down the number of calls made directly from the game client. This change means that instead of making six or more unique API calls in C++," +
						" I will only be writing two: one to post feedback information to the server, and one to post the image to the server.",
				}),
				a.P(g.EB{
					Text: "I believe that this redesign will reduce the remaining time on the API calls by approximately three hours due to my familiarity with JavaScript and the availability of API libraries.",
				}),
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Src:     files.IndustryProjectRedesign.ServerDesign,
					Caption: "Server Design",
				}),
				a.P(g.EB{
					Children: g.CE{
						g.Text("The above diagram demonstrates the intended flow of the tool from the user through the game and server to the various APIs. The game" +
							" will send the feedback and the screenshot to the server, where the screenshot will be saved in an "),
						a.TextLink("Amazon S3 Bucket", "https://aws.amazon.com/s3/"),
						g.Text(" for later reference. This screenshot’s URL will then be retrieved and passed to the feedback section."),
					},
				}),
				a.P(g.EB{
					Text: "While the screenshot uploads, the server will simultaneously request a project board of the submitted feedback type and retrieve a" +
						" summary card from that board. The player's feedback will then be posted as a new card to that board with the screenshot URL," +
						" and the summary card will be updated to include the new feedback.",
				}),
				a.P(g.EB{
					Text: "Finally, the server will perform some analysis on the feedback data to find hot-spots in levels where players have been sending reports most" +
						" and the average sentiment of reports in that area. The server will then post that analysis with the summary information to a discord" +
						" server of the developer's choosing at regular intervals.",
				}),
				a.P(g.EB{
					Children: g.CE{
						g.Text("This rework has a few benefits. Firstly, many APIs have official or third-party "),
						a.TextLink("libraries", "https://developer.github.com/v3/libraries/"),
						g.Text(" for making use of the API significantly simpler. This will speed up development by cutting down the amount of work required to implement" +
							" and test each call. Secondly, this removes the security risk of publishing work with API keys to the public as the remote server will have these" +
							" keys instead. Finally, this reduces processing and networking load done in the game client. Originally, it was intended to perform some analysis" +
							" and generate summaries from within the game client before posting, where now the server will bear this load."),
					},
				}),
				c.SectionHeader(c.SectionHeaderProps{
					TagKey: "next-time",
					Text:   "For Next Week",
					Header: g.H4,
				}),
				a.P(g.EB{
					Text: "I believe that this redesign will reduce the remaining time on the API calls by approximately three hours due to my familiarity with JavaScript and the availability of API libraries.",
				}),
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Src:     files.IndustryProjectRedesign.Week5BurndownPlan,
					Caption: "Burndown Plan for Week 5",
				}),
			},
		})
	},
}
