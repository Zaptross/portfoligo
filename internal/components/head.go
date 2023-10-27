package components

import (
	g "github.com/zaptross/gorgeous"
	"github.com/zaptross/portfoligo/internal/types"
)

func Head(md types.PageDetails) *g.HTMLElement {
	return g.Head(g.EB{
		Children: append(append(append(Meta(md), Favicon()...), PrismJS()...), OpenGraph(md)...),
	})
}

func Meta(md types.PageDetails) []*g.HTMLElement {
	return []*g.HTMLElement{
		g.Meta(g.EB{
			Props: g.Props{
				"charset": "UTF-8",
			},
		}),
		g.Meta(g.EB{
			Props: g.Props{
				"name":    "viewport",
				"content": "width=device-width, initial-scale=1.0",
			},
		}),
		g.Title(g.EB{
			Text: md.Title,
		}),
		g.Link(g.EB{
			Props: g.Props{
				"rel":  "stylesheet",
				"type": "text/css",
				"href": "style.css",
			},
		}),
		g.Script(g.EB{
			Props: g.Props{
				"type": "text/javascript",
				"src":  "script.js",
			},
		}),
		g.Style(g.EB{
			Text: "@import url('https://fonts.googleapis.com/css2?family=Lato:wght@400&display=swap');",
		}),
	}
}

func OpenGraph(md types.PageDetails) g.CE {
	return g.CE{
		openGraphMeta("og:title", md.Title),
		openGraphMeta("description", md.Description),
		openGraphMeta("og:description", md.Description),
		// openGraphMeta("og:image", "https://dystrophygame.com/og-image.png"),
	}
}

func openGraphMeta(name, content string) *g.HTMLElement {
	return g.Meta(g.EB{
		Props: g.Props{
			"name":    name,
			"content": content,
		},
	})
}

func PrismJS() []*g.HTMLElement {
	return []*g.HTMLElement{
		// Add syntax highlighting for code blocks with Prism
		g.Link(g.EB{
			Props: g.Props{
				"rel":  "stylesheet",
				"type": "text/css",
				"href": "https://cdnjs.cloudflare.com/ajax/libs/prism/1.23.0/themes/prism.min.css",
			},
		}),
		loadScript("https://cdnjs.cloudflare.com/ajax/libs/prism/1.23.0/components/prism-core.min.js"),
		loadScript("https://cdnjs.cloudflare.com/ajax/libs/prism/1.23.0/plugins/autoloader/prism-autoloader.min.js"),
	}
}

func loadScript(src string) *g.HTMLElement {
	return g.Script(g.EB{
		Props: g.Props{
			"type": "text/javascript",
			"src":  src,
		},
	})
}
