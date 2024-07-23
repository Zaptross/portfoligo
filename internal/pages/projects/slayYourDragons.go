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

var SlayYourDragons = types.PageDetails{
	Title:       "Slay Your Dragons",
	Description: "An applied game designed to increase the frequency of positive studying behaviors.",
	Slug:        "slay-your-dragons",
	Type:        types.TYPE_PROJECT,
	Written:     time.Date(2019, 8, 20, 0, 0, 0, 0, time.UTC),
	Tags: []string{
		types.TAG_STUDENT,
		types.TAG_GAME,
		types.TAG_DESIGN,
		types.TAG_LANG_CSHARP,
		types.TAG_TOOL_UNITY,
	},
	Content: func(p types.PageDetails) *g.HTMLElement {
		return g.Div(g.EB{
			Children: []*g.HTMLElement{
				a.Row(g.CE{
					a.Col(g.CE{
						a.P(g.EB{
							Children: g.CE{
								c.Emphasis("Slay Your Dragons"),
								g.Text(" is an applied game where you, a daring adventurer, seek out to slay the dragons which plague your homelands. You set up your assignments or projects as boss monsters in the world and as you work towards them, you fight and slay them in the gamified world."),
							},
						}),
						a.P(g.EB{Text: "This applied game was designed and built in a team of three to motivate young adults to start projects earlier. Within the team I undertook the role of programmer, and contributed to both the game and menu designs. I implemented the game logic for creating a boss, battling a boss, and purchasing items in the game's store with in-game currency."}),
					}, nil),
				}, nil),
				o.Carousel(o.CarouselProps{
					ClassList: []string{ch.W("clamp(200px,80vw,343px)")},
					Children: g.CE{
						m.ProjectImage(p, files.SlayYourDragons.BossScreenTutorial, nil),
						m.ProjectImage(p, files.SlayYourDragons.BossScreenLayout, nil),
						m.ProjectImage(p, files.SlayYourDragons.BossSetupDemo, nil),
					},
				}),
			}},
		)
	},
}
