package projects

import (
	"time"

	g "github.com/zaptross/gorgeous"
	"github.com/zaptross/portfoligo/generated/files"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	c "github.com/zaptross/portfoligo/internal/components"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	m "github.com/zaptross/portfoligo/internal/components/molecules"
	o "github.com/zaptross/portfoligo/internal/components/organisms"
	"github.com/zaptross/portfoligo/internal/types"
)

var AshesOfTheVeil = types.PageDetails{
	Title:       "Ashes of the Veil",
	Description: "Ashes of the Veil is a first person investigation game, developed and published as the capstone university project of a team of five.",
	Slug:        "ashes-of-the-veil",
	Type:        types.TYPE_PROJECT,
	Written:     time.Date(2019, 8, 19, 0, 0, 0, 0, time.UTC),
	Tags: []string{
		types.TAG_STUDENT,
		types.TAG_GAME,
		types.TAG_TOOL_UNITY,
		types.TAG_LANG_CSHARP,
		types.TAG_AVAILABLE,
	},
	Content: func(p types.PageDetails) *g.HTMLElement {
		return g.Div(g.EB{
			Children: []*g.HTMLElement{
				c.YoutubeEmbed("Ashes of the Veil Trailer", "https://www.youtube-nocookie.com/embed/9jO1RffzHp4"),
				a.P(g.EB{Text: p.Description}),
				a.P(g.EB{Text: "During the course of development, I undertook the roles of technical lead, quality assurance tester, and producer."}),
				c.Row(g.CE{
					c.Col(g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    p,
							Caption: "Project Management Spreadsheet",
							Src:     files.AshesOfTheVeil.ProjectManagementSheet,
						}),
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    p,
							Caption: "Quality Assurance Spreadsheet",
							Src:     files.AshesOfTheVeil.QualityAssuranceSheet,
						}),
					}, []string{ch.MxW("45%"), ch.Flex(2)}),
					c.Col(g.CE{
						a.P(g.EB{
							Text: "As the producer, I created and managed our project tracking spreadsheet, organised meetings with the other team members, and" +
								" presented our final piece to the course co-ordinator and a panel of local game developers.",
						}),
						a.P(g.EB{
							Text: "As the quality assurance tester, I also created our QA spreadsheet, before running full and partial QA tests on our list of 500 items.",
						}),
						a.P(g.EB{
							Text: "As technical lead on the project, I was programming some main features of the game and managing the two other programmers in our team." +
								" I designed and implemented the lens mechanic of the game, whereby the magical lens allows the player to see into the past, see residue from" +
								" where others had used magic, and see omens of peril from the future highlighted in text. I designed and implemented the shader on the lens" +
								" which only shows the highlighted elements within the view of the lens.",
						}),
					}, []string{ch.Flex(3)}),
				}, nil),
				a.P(g.EB{
					ClassList: []string{ch.TextAlign(ch.Align.Center)},
					Children: g.CE{
						g.Text("You can find the project on "),
						a.TextLink("GameJolt", "https://gamejolt.com/games/Ashes_Of_The_Veil/372720"),
						g.Text(" and also on "),
						a.TextLink("IndieDB", "https://www.indiedb.com/games/ashes-of-the-veil"),
						g.Text("."),
					},
				}),
				o.Carousel(o.CarouselProps{
					ClassList: []string{ch.Aspect("16/9")},
					Children: g.CE{
						m.ProjectImage(p, files.AshesOfTheVeil.Note, nil),
						m.ProjectImage(p, files.AshesOfTheVeil.RedLensDemo, nil),
						m.ProjectImage(p, files.AshesOfTheVeil.BlueLensDemo, nil),
					},
				}),
			}},
		)
	},
}
