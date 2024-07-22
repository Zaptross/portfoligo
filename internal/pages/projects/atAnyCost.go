package projects

import (
	"time"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	c "github.com/zaptross/portfoligo/internal/components"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	m "github.com/zaptross/portfoligo/internal/components/molecules"
	files "github.com/zaptross/portfoligo/internal/generated/files"
	"github.com/zaptross/portfoligo/internal/theme"
	"github.com/zaptross/portfoligo/internal/types"
)

var AtAnyCost = types.PageDetails{
	Title:       "At Any Cost",
	Description: "At Any Cost is a science-fiction tabletop role-playing game setting, written in a student team.",
	Slug:        "at-any-cost",
	Type:        types.TYPE_PROJECT,
	Written:     time.Date(2019, 9, 28, 0, 0, 0, 0, time.UTC),
	Tags: []string{
		types.TAG_STUDENT,
		types.TAG_GAME,
		types.TAG_WRITING,
	},
	Content: func(p types.PageDetails) *g.HTMLElement {
		t := theme.UseTheme()

		return g.Div(g.EB{
			Children: []*g.HTMLElement{
				c.Row(g.CE{
					c.Col(g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Caption: "At Any Cost",
							Page:    p,
							Src:     files.AtAnyCost.Cover,
						}),
						m.CaptionedImage(m.CaptionedImageProps{
							Caption: "The Empyrean Bullion Faction",
							Page:    p,
							Src:     files.AtAnyCost.EmpyreanBullion,
						}),
						c.Row(g.CE{
							m.CaptionedImage(m.CaptionedImageProps{
								Caption:          "Attribution",
								Page:             p,
								Src:              files.AtAnyCost.AttributionPage,
								ContainerClasses: []string{ch.Flex(1), ch.Margin(t.Spacing(2))},
							}), m.CaptionedImage(m.CaptionedImageProps{
								Caption:          "Contents Page 1",
								Page:             p,
								Src:              files.AtAnyCost.TableOfContents,
								ContainerClasses: []string{ch.Flex(1), ch.Margin(t.Spacing(2))},
							}),
							m.CaptionedImage(m.CaptionedImageProps{
								Caption:          "Contents Page 2",
								Page:             p,
								Src:              files.AtAnyCost.TableOfContents2,
								ContainerClasses: []string{ch.Flex(1), ch.Margin(t.Spacing(2))},
							}),
							m.CaptionedImage(m.CaptionedImageProps{
								Caption:          "Ship Rules",
								Page:             p,
								Src:              files.AtAnyCost.ShipRules,
								ContainerClasses: []string{ch.Flex(1), ch.Margin(t.Spacing(2))},
							}),
						}, []string{ch.Margin(t.Spacing(2))}),
					}, []string{ch.Flex(1)}),
					c.Col(g.CE{
						a.P(g.EB{
							Children: g.CE{
								c.Emphasis("At Any Cost"),
								g.Text(" is a science-fiction tabletop role-playing game setting built upon the "),
								a.TextLink("Fate Accelerated", "https://evilhat.com/product/fate-accelerated-edition/"),
								g.Text(" TTRPG system."),
							},
						}),
						a.AuthorQuote(
							"In the year 13812, mega-corporations have long divided the solar system up amongst each other. In this corporate government "+
								"landscape tensions are running high as inter-corporate vies for power strangle competition in the face of dwindling new resources. When a new "+
								"solar system is discovered a power vacuum was created, drawing the attention of these mega-corporations to it. Whether it be via the means of "+
								"force, negotiation or sabotage, the colonization of the Tyche System means life or death for a corporation. As the risk of outright war ever looms"+
								" over everyone involved, the only thing stopping it from escalating is the thin veil of peace that the masses still believe in.",
							"At Any Cost, Page 7",
							t.Colors.Pallette.Green,
							[]string{ch.Margin("unset")},
						),
						a.P(g.EB{
							Text: "For the month long duration of the project, I took on the roles of producer, editor, and a writer.",
						}),
						a.P(g.EB{
							Text: "The world of At Any Cost was created from by combining the team's individual world concepts. When filtering down our ideas to their most" +
								" primitive form we found a unifying theme - humanity. Specifically, how humanity will fare in the face of a new space era, and the political " +
								"structures underlying a technologically advanced multi-planetary society.",
						}),
						a.P(g.EB{
							Text: "In my role as a writer, I took on the tasks of designing and writing two of the nine factions in our setting including the Empyrean" +
								" Bullion. This faction's purpose was twofold: to drive exploration in the second solar system through mining opportunities, and to create" +
								" corporate ties which then drew lines of alliance in the sand.",
						}),
						a.P(g.EB{
							Text: "In addition to writing factions, I undertook the writing of the two solar systems: the Sol System and the Tyche System. " +
								"The Sol System is the re-contextualization of our own solar system within the world of At Any Cost. Each planet has been colonised by" +
								" corporations and average citizens over the last two millenia.",
						}),
						a.P(g.EB{
							Text: "The Tyche System is an entirely fictional solar system which is the closest to the Sol System. This system was designed as a mostly" +
								" uninfluenced solar system to allow a greater variety of stories to be told in At Any Cost.",
						}),
						a.P(g.EB{
							Text: "The last of my writing on the book took place in the combat and character advancement sections. Here I altered the Fate Accelerated base" +
								" combat system to add a combat variant for space ships, and associated character sheets. Secondly, I added a variant character advancement " +
								"system which drives characters to focus on personal goals more than combat, as is often the case in traditional TTRPG's.",
						}),
						a.P(g.EB{
							Text: "Finally in my role as producer and editor, it was my job to manage team meetings, create and maintain the project tracking spreadsheet," +
								" and to perform the full editing passes of the document.",
						}),
						a.P(g.EB{
							Text: "As yet this project is unavailable for public download, but when it does become available, a link will be added here.",
						}),
					}, []string{ch.Flex(1)}),
				}, nil),
				c.Col(g.CE{
					c.Row(g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    p,
							Caption: "Inner Sol System",
							Src:     files.AtAnyCost.SolSystem1,
						}),
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    p,
							Caption: "Outer Sol System",
							Src:     files.AtAnyCost.SolSystem2,
						}),
					}, nil),
					c.Row(g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    p,
							Caption: "Inner Tyche System",
							Src:     files.AtAnyCost.TycheSystem1,
						}),
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    p,
							Caption: "Outer Tyche System",
							Src:     files.AtAnyCost.TycheSystem2,
						}),
					}, nil),
				}, nil),
			}},
		)
	},
}
