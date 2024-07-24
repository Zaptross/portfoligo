package projects

import (
	"time"

	g "github.com/zaptross/gorgeous"
	"github.com/zaptross/portfoligo/generated/files"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	m "github.com/zaptross/portfoligo/internal/components/molecules"
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
		return g.Div(g.EB{
			Children: []*g.HTMLElement{
				m.ProjectImage(p, files.PlayerLikeMechAi.MechAi, []string{ch.W("100%")}),
				a.P(g.EB{
					Text: "This project was a university assignment I completed wherein I had to design and implement a decision tree for a game AI to emulate a player" +
						" in a mech battling game.",
				}),
				m.R2C(
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    p,
							Src:     files.PlayerLikeMechAi.BehaviourTree,
							Caption: "Behaviour Tree",
						}),
					},
					[]string{ch.Flex(2)}, g.CE{
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
							Text: "The AI used it's current ammunition, health, and number of known opponents to determine how aggressive to be. When running low on" +
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
					},
					nil,
				),
			}},
		)
	},
}
