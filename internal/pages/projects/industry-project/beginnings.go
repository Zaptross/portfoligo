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

var Beginnings = types.PageDetails{
	Title:       "Industry Project: Beginnings",
	Description: "The first week of working on the industry project, spent learning Unreal Engine and planning.",
	Slug:        "industry-project-beginnings",
	Type:        types.TYPE_PROJECT,
	Written:     time.Date(2019, 9, 23, 0, 0, 0, 0, time.UTC),
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
					Caption: "Burndown Chart for Week 2",
					Src:     files.IndustryProjectBeginnings.BurndownChart,
					ContainerClasses: []string{
						ch.MxH("fit-content"),
					},
				}),
				a.P(g.EB{
					Children: g.CE{
						g.Text("Unfortunately, due to other work the week prior, this project had to be started in what was intended to be the second week of development " +
							"and is a catch-up week. Though I didn't completely catch up this week, I managed to make up about one third of the last week's time. " + "For more detail, see the "),
						a.TextLink("project board", "https://github.com/Zaptross/flexifeedback/projects/1"),
						g.Text("."),
					},
				}),
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Caption: "Testing UI Elements",
					Src:     files.IndustryProjectBeginnings.TestingUnrealUi,
					ContainerClasses: []string{
						ch.MxH("fit-content"),
					},
				}),
				a.P(g.EB{Text: "This gif demonstrates the state of the UI elements of the project at the end of the week. So far, the UI is opening and closing, " + "laid out roughly according to the design, and the dropdown and buttons are functional."}),
				c.SectionHeader(c.SectionHeaderProps{
					Text:   "Learning Unreal",
					TagKey: "learning-unreal",
					Header: g.H4,
				}),
				a.P(g.EB{
					Children: g.CE{
						g.Text("The beginning steps of working on my industry project is starting to learn how to interact with the Unreal Engine. I decided to start" + " my journey into Unreal by exploring it's built-in tutorial system "),
						a.TextLink("through the lens of unity", "https://docs.unrealengine.com/en-US/GettingStarted/FromUnity/index.html"),
						g.Text(", which guided me through all the elements of the editor."),
					},
				}),
				c.Row(g.CE{
					m.CaptionedImage(m.CaptionedImageProps{
						Page:             p,
						Caption:          "Pulsing Darkness",
						Src:              files.IndustryProjectBeginnings.UnrealShaderWiggle,
						ContainerClasses: []string{ch.Flex(1)},
					}),
					m.CaptionedImage(m.CaptionedImageProps{
						Page:             p,
						Caption:          "Black Hole Shader",
						Src:              files.IndustryProjectBeginnings.UnrealShaderBlackHole,
						ContainerClasses: []string{ch.Flex(1)},
					}),
					a.P(g.EB{
						ClassList: []string{ch.Flex(2)},
						Children: g.CE{
							g.Text("During this process, I found myself experimenting with the material editor borne of my familiarity with Unity's "),
							a.TextLink("new rendering pipelines", "https://unity.com/srp/universal-render-pipeline"),
							g.Text(",  and made a couple of cool effects."),
						},
					}),
				}, nil),
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Caption: "Unreal Engine Shader Graph Editor",
					Src:     files.IndustryProjectBeginnings.UnrealShaderGraph,
					ContainerClasses: []string{
						ch.MxH("fit-content"),
					},
				}),
				c.Row(g.CE{
					m.CaptionedImage(m.CaptionedImageProps{
						Page:             p,
						Caption:          "Bouncy Chair",
						Src:              files.IndustryProjectBeginnings.UnrealChairBounce,
						ContainerClasses: []string{ch.Flex(1)},
					}),
					m.CaptionedImage(m.CaptionedImageProps{
						Page:             p,
						Caption:          "Bouncy Chair Blueprint",
						Src:              files.IndustryProjectBeginnings.UnrealChairBounceBlueprint,
						ContainerClasses: []string{ch.Flex(1)},
					}),
					a.P(g.EB{
						ClassList: []string{ch.Flex(2)},
						Text:      "I then began to experiment with scene objects and blueprints, deciding to make an object interact with the world in unintuitive ways.",
					}),
				}, nil),
				c.SectionHeader(c.SectionHeaderProps{
					TagKey: "prototyping-ui",
					Text:   "Building the Prototype UI",
					Header: g.H4,
				}),
				a.P(g.EB{
					Text: "The next step for the week was beginning the prototyping process for the UI elements of the project, as per my initial mock-ups. This " +
						"process expanded upon my rough understanding of the Unreal Engine's Blueprints, and by the end of the process I felt very comfortable programming visually.",
				}),
				c.Row(g.CE{
					m.CaptionedImage(m.CaptionedImageProps{
						Page:             p,
						Src:              files.IndustryProjectBeginnings.ProtoLayout1,
						Caption:          "Mockup Button",
						ContainerClasses: []string{ch.Flex(1)},
					}),
					m.CaptionedImage(m.CaptionedImageProps{
						Page:             p,
						Src:              files.IndustryProjectBeginnings.ProtoLayout2,
						Caption:          "Mockup Layout",
						ContainerClasses: []string{ch.Flex(1)},
					}),
				}, nil),
				c.Row(g.CE{
					m.CaptionedImage(m.CaptionedImageProps{
						Page:             p,
						Src:              files.IndustryProjectBeginnings.UnrealTestingUiElements,
						Caption:          "Testing UI Elements",
						ContainerClasses: []string{ch.Flex(1)},
					}),
					m.CaptionedImage(m.CaptionedImageProps{
						Page:             p,
						Src:              files.IndustryProjectBeginnings.UnrealInputFieldsClearing,
						Caption:          "Input Fields Clearing",
						ContainerClasses: []string{ch.Flex(1)},
					}),
				}, nil),
				c.SectionHeader(c.SectionHeaderProps{
					TagKey: "screenshot",
					Text:   "Taking the Screenshot",
					Header: g.H4,
				}),
				a.P(g.EB{
					Text: "Taking the screenshot felt like the perfect next step as conceptually it was the next most complex part of the design. This proved itself" +
						" quickly true, as I encountered difficulties in trying to keep frame overhead low, mismatched type errors, and determining which type of camera to use.",
				}),
				a.P(g.EB{
					Text: "I discovered that the SceneCapture2D camera was the camera which would serve my purpose best, and that it was not able to be spawned and" +
						" destroyed in code. This was solved by creating a blueprint of the object, finding it by class, then manually capturing the scene.",
				}),
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Src:     files.IndustryProjectBeginnings.UnrealScreenshotBlueprint,
					Caption: "Screenshot Blueprint Function",
					ContainerClasses: []string{
						ch.MxH("fit-content"),
					},
				}),
				a.P(g.EB{
					Text: "Finally, when I had reworked my screenshot functions, I tested it by creating a material from the render texture and applying it to a cube in" +
						" the scene demonstrating that the rendertexture is only updated when the submit feedback panel is opened.",
				}),
				m.CaptionedImage(m.CaptionedImageProps{
					Page:    p,
					Src:     files.IndustryProjectBeginnings.UnrealScreenshot,
					Caption: "Screenshot Functionality",
					ContainerClasses: []string{
						ch.MxH("fit-content"),
					},
				}),
			},
		})
	},
}
