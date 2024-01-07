package pages

import (
	"time"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	c "github.com/zaptross/portfoligo/internal/components"
	"github.com/zaptross/portfoligo/internal/types"
)

var (
	SlayYourDragons = registerPage(types.PageDetails{
		Title:       "Slay Your Dragons",
		Description: "An applied game for increasing frequency of positive studying behaviors.",
		Slug:        "slay-your-dragons",
		Type:        TYPE_PROJECT,
		Written:     time.Date(2019, 8, 20, 0, 0, 0, 0, time.UTC),
		Tags: []string{
			TAG_GAME,
			TAG_LANG_CSHARP,
			TAG_DESIGN,
		},
		Content: func(p types.PageDetails) *g.HTMLElement {
			return g.Div(g.EB{
				Children: []*g.HTMLElement{
					c.Row(g.CE{
						c.Col(g.CE{
							c.P(g.EB{
								Children: g.CE{
									c.Emphasis("Slay Your Dragons"),
									g.Text(" is an applied game where you, a daring adventurer, seek out to slay the dragons which plague your homelands. You set up your assignments or projects as boss monsters in the world and as you work towards them, you fight and slay them in the gamified world."),
								},
							}),
							c.P(g.EB{Text: "This applied game was designed and built in a team of three to motivate young adults to start projects earlier. Within the team I undertook the role of programmer, and contributed to both the game and menu designs. I implemented the game logic for creating a boss, battling a boss, and purchasing items in the game's store with in-game currency."}),
						}, nil),
					}, nil),
					c.Row(g.CE{
						c.Carousel(c.CarouselProps{
							ClassList: []string{ch.W("clamp(200px,80vw,343px)")},
							Children: g.CE{
								c.ImageRel(p, "boss-screen-tutorial.png", nil),
								c.ImageRel(p, "boss-screen-layout.png", nil),
								c.ImageRel(p, "boss-setup-demo.gif", nil),
							},
						}),
					}, []string{"flex-wrap", ch.JustifyContent(ch.Content.Center)}),
					// }, []string{"justify-space-around", "flex-wrap"}),
				}},
			)
		},
	})
)
