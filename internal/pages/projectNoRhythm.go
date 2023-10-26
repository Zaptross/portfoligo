package pages

import (
	"time"

	g "github.com/zaptross/gorgeous"
	c "github.com/zaptross/portfoligo/internal/components"
	"github.com/zaptross/portfoligo/internal/types"
)

var (
	_ = registerPage(types.PageDetails{
		Title:       "No Rhythm",
		Description: "A rhythm game made for Ludum Dare 45.",
		Slug:        "no-rhythm",
		Type:        TYPE_PROJECT,
		Written:     time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
		Tags: []string{
			TAG_GAME,
			TAG_LANG_CSHARP,
			TAG_AVAILABLE,
		},
		Content: g.Div(g.EB{
			Children: []*g.HTMLElement{
				c.YoutubeEmbed("NoRhythm Trailer", "https://www.youtube-nocookie.com/embed/arLUT1Lep3s?si=AlY5vs9L0MEj7vsy"),
				c.P(g.EB{
					Children: g.CE{
						c.Emphasis("NoRhythm"),
						g.Text(" is an arcade rhythm game designed and created in three days by a team of two for "),
						c.LinkExternal("Ludum Dare 45", "https://ldjam.com/events/ludum-dare/45"),
						g.Text(". This game was inspired by the theme “Start with nothing” and it's design was "),
						g.Text("driven by the musical talents of my teammate. NoRhythm was created in the Unity game engine with a total of 67 hours of cumulative work from both team members."),
					},
				}),
				c.P(g.EB{
					Text: "My roles in the project were producer, programmer, game designer, and UI designer. I developed and implemented all user interface and gameplay elements, created level design tools for my teammate, and created marketing materials for publishing.",
				}),
				c.P(g.EB{
					Children: g.CE{
						g.Text("NoRhythm was published to "),
						c.LinkExternal("GameJolt", "https://gamejolt.com/games/NoRhythm/443325"),
						g.Text(", it can be played and downloaded or played in the browser from the links below!"),
					},
				}),
			}},
		),
	})
)
