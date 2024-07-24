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

var KowloonExigency = types.PageDetails{
	Title:       "Kowloon Exigency",
	Description: "Kowloon Exigency is a short top-down shooter made as a level design excercise in a small team",
	Slug:        "kowloon-exigency",
	Type:        types.TYPE_PROJECT,
	Written:     time.Date(2019, 11, 2, 0, 0, 0, 0, time.UTC),
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
				m.ProjectImage(p, files.KowloonExigency.Title, []string{ch.W("100%")}),
				a.P(g.EB{
					Text: "Kowloon Exigency is a short top-down shooter made in a team of four, focusing on creating an intriguing and visually complex level design.",
				}),
				a.P(g.EB{
					Text: "Within this team of four I undertook the roles of narrative designer, version control manager, cinematographer, and level designer.",
				}),
				m.SectionHeader(m.SectionHeaderProps{
					TagKey: "levels",
					Text:   "Level Design",
					Header: g.H3,
				}),
				m.R2C(
					[]string{ch.Flex(1), ch.JustifyContent(ch.Align.Center)}, g.CE{
						a.P(g.EB{Text: "In my role as a level designer, I created four level tiles for the project."}),
						a.P(g.EB{Text: "These tiles included a slum-like street, a run-down office corridor, a busy office space, and a penthouse balcony where the final boss battle takes place."}),
						a.P(g.EB{Text: "Below is a gallery of these tiles."}),
					},
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    p,
							Caption: "Four Level Tiles",
							Src:     files.KowloonExigency.FourTiles,
						}),
					},
					nil),
				m.R2C(
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{Page: p, Caption: "Slum Street Tile", Src: files.KowloonExigency.SlumTile}),
					},
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{Page: p, Caption: "Tile During Play", Src: files.KowloonExigency.SlumTilePlay}),
					},
					nil,
				),
				m.R2C(
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{Page: p, Caption: "Office Tile", Src: files.KowloonExigency.OfficeTile}),
					},
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{Page: p, Caption: "Tile Overhead", Src: files.KowloonExigency.OfficeTileOverhead}),
					},
					nil,
				),
				m.R2C(
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{Page: p, Caption: "Corridor Tile", Src: files.KowloonExigency.CorridorTile}),
					},
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{Page: p, Caption: "Penthouse Balcony", Src: files.KowloonExigency.PenthouseTile}),
					},
					nil,
				),
				m.SectionHeader(m.SectionHeaderProps{
					TagKey: "cinematography",
					Text:   "Cutscenes, Camera, and Narrative",
					Header: g.H3,
				}),
				m.R2C(
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{
							Page:    p,
							Caption: "Elevator Unlocked",
							Src:     files.KowloonExigency.CutsceneElevator,
						}),
					},
					nil, g.CE{
						a.P(g.EB{
							Text: "In my roles as cinematographer and narrative designer I developed extra features for our game's camera. These features included" +
								" changing it's position to ensure the player's avatar is always in view, moving to show cutscenes, and shifting perspective at the end of a level.",
						}),
						a.P(g.EB{
							Text: "In addition to developing the camera, I also implemented the cutscenes I designed. These cutscenes were designed to fulfill both narrative" +
								" and gameplay functions, by conveying story and guiding the player towards their objectives.",
						}),
						a.P(g.EB{
							Text: "Below is a gallery of my work on the game's cutscenes and camera.",
						}),
					},
					nil),
				m.R2C(
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{Page: p, Caption: "Keeping Player Onscreen", Src: files.KowloonExigency.CameraKeepPlayerOnscreen}),
					},
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{Page: p, Caption: "Dossier Narrative Introduction", Src: files.KowloonExigency.CutsceneDossier}),
					},
					nil),
				m.R2C(
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{Page: p, Caption: "Police Blockade", Src: files.KowloonExigency.CutsceneBlockade}),
					},
					nil, g.CE{
						m.CaptionedImage(m.CaptionedImageProps{Page: p, Caption: "End of Level Perspective", Src: files.KowloonExigency.CutsceneLevelEnd}),
					},
					nil),
			}},
		)
	},
}
