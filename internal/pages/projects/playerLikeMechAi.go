package projects

import (
	"time"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	c "github.com/zaptross/portfoligo/internal/components"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	"github.com/zaptross/portfoligo/internal/theme"
	"github.com/zaptross/portfoligo/internal/types"
)

var PlayerLikeMechAi = types.PageDetails{
	Title:       "Player-Like Mech AI",
	Description: "A university project to create a player-like AI for a mech game using behaviour trees.",
	Slug:        "player-like-mech-ai",
	Type:        types.TYPE_PROJECT,
	Written:     time.Date(2019, 8, 22, 0, 0, 0, 0, time.UTC),
	Tags: []string{
		types.TAG_STUDENT,
		types.TAG_GAME,
		types.TAG_TOOL_UNITY,
		types.TAG_LANG_CSHARP,
	},
	Content: func(p types.PageDetails) *g.HTMLElement {
		t := theme.UseTheme()

		twoCols := "two-columns"
		mediaAspect := "(max-aspect-ratio: 7/9)"
		g.Media(mediaAspect, "."+twoCols, g.CSSProps{
			"flex-direction": "column-reverse",
		})
		g.Media(mediaAspect, "."+twoCols+">.col", g.CSSProps{
			"max-width": "100%",
		})
		g.Media(mediaAspect, "."+twoCols+">.col + .col", g.CSSProps{
			"margin-left": "0",
		})

		return g.Div(g.EB{
			Children: []*g.HTMLElement{
				c.ImgProject(p, "mech-ai.gif", []string{ch.W("100%")}),
				a.P(g.EB{
					Text: "This project was a university assignment I completed wherein I had to design and implement a decision tree for a game AI to emulate a player" +
						" in a mech battling game.",
				}),
				c.Row(g.CE{
					c.Col(g.CE{
						c.ImgProject(p, "behaviour-tree.png", nil),
					}, []string{ch.FlexGrow(1), ch.MxW("45%")}),
					c.Col(g.CE{
						a.P(g.EB{
							Children: g.CE{
								g.Text("For this project, I implemented behaviour functions in the main AI script and ran all decision logic from the behaviour tree, "),
								a.TextLink("using PandaBT", "https://assetstore.unity.com/packages/tools/behavior-ai/panda-bt-free-33057"),
								g.Text(" a behaviour tree library for Unity."),
							},
						}),
						a.P(g.EB{
							Text: "Part of being a player-like AI meant that I could only use information that a player could reasonably have had access to, like line of" +
								" sight to enemies, guessing where enemies were going when leaving line of sight, and only attempting to gather pickups that they had seen before.",
						}),
						a.P(g.EB{
							Text: "The AI used it’s current ammunition, health, and number of known opponents to determine how aggressive to be. When running low on" +
								" ammunition and health, the need for pickups greatly increases. Players also tend to try and disengage when being targeted by multiple" +
								" aggressors, which this AI does also.",
						}),
						a.P(g.EB{
							Text: "The AI was implemented in C# using the Unity game engine, built from a base game provided by the teaching staff at the university.",
						}),
						a.P(g.EB{
							Children: g.CE{
								g.Text("You can find my "),
								a.TextLink("decision tree and AI script", "https://github.com/Zaptross/ExampleWork/tree/master/C%23/Game%20AI"),
								g.Text(" on my GitHub."),
							},
						}),
					}, []string{ch.MarginL(t.Spacing(4)), ch.FlexGrow(2)}),
				}, []string{twoCols}),
			}},
		)
	},
}
