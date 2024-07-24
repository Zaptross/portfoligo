package atoms

import (
	g "github.com/zaptross/gorgeous"
)

func Fonts() g.CE {
	return g.CE{
		g.Link(g.EB{
			Props: g.Props{
				"rel":  "preconnect",
				"href": "https://fonts.googleapis.com",
			},
		}),
		g.Link(g.EB{
			Props: g.Props{
				"rel":         "preconnect",
				"href":        "https://fonts.gstatic.com",
				"crossorigin": "true",
			},
		}),
		g.Style(g.EB{
			Text: "@import url('https://fonts.googleapis.com/css2?family=Lato:wght@400&display=swap');",
		}),
		g.Style(g.EB{
			Text: "@import url('https://fonts.googleapis.com/css2?family=Roboto:ital,wght@400&display=swap');",
		}),
	}
}
