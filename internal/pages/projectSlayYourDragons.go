package pages

import (
	"time"

	g "github.com/zaptross/gorgeous"
	c "github.com/zaptross/portfoligo/internal/components"
	"github.com/zaptross/portfoligo/internal/types"
)

var (
	_ = registerPage(types.PageDetails{
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
		Content: func() *g.HTMLElement {
			return g.Div(g.EB{
				Children: []*g.HTMLElement{
					c.P(g.EB{
						Children: g.CE{
							c.Emphasis("Slay Your Dragons"),
							g.Text(" is an applied game where you, a daring adventurer, seek out to slay the dragons which plague your homelands."),
						},
					}),
					c.P(g.EB{Text: "You set up your assignments or projects as boss monsters in the world and as you work towards them, you fight and slay them in the gamified world."}),
					c.P(g.EB{Text: "This applied game was designed and built in a team of three to motivate young adults to start projects earlier."}),
					c.P(g.EB{Text: "Within the team I undertook the role of programmer, and contributed to both the game and menu designs."}),
					c.P(g.EB{Text: "I implemented the game logic for creating a boss, battling a boss, and purchasing items in the game's store with in-game currency."}),
				}},
			)
		},
	})
)
