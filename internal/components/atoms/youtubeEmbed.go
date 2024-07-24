package atoms

import (
	g "github.com/zaptross/gorgeous"
)

func YoutubeEmbed(title string, url string) *g.HTMLElement {
	divClassName := "youtube-embed-container"
	g.Class(&g.CSSClass{
		Selector: "." + divClassName,
		Props: g.CSSProps{
			"display":         "flex",
			"justify-content": "center",
		},
	})

	embedClassName := "youtube-embed"
	g.Class(&g.CSSClass{
		Selector: "." + embedClassName,
		Props: g.CSSProps{
			"aspect-ratio": "16/9",
			"width":        "100%",
		},
	})

	return g.Div(
		g.EB{
			ClassList: []string{divClassName},
			Children: g.CE{
				g.CustomElement(
					"iframe",
					true,
					g.EB{
						ClassList: []string{embedClassName},
						Props: g.Props{
							"src":             url,
							"title":           title,
							"frameborder":     "0",
							"allow":           "accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share",
							"allowfullscreen": "",
						},
					},
				),
			},
		},
	)
}
