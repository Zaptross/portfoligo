package organisms

import (
	"fmt"
	"os"

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
		g.Script(g.EB{
			Props: g.Props{
				"src":         "https://kit.fontawesome.com/653251a210.js",
				"crossorigin": "anonymous",
			},
		}),
		g.Style(g.EB{
			Text: "@import url('https://fonts.googleapis.com/css2?family=Lato:wght@400&display=swap');",
		}),
	}
}

func OpenGraph(md types.PageDetails) g.CE {
	// TODO - update to correct domain
	domain := "https://portfolio.zaptross.com"
	ogElements := g.CE{
		openGraphMeta("og:title", md.Title),
		openGraphMeta("description", md.Description),
		openGraphMeta("og:description", md.Description),
		openGraphMeta("og:type", "website"),
		openGraphMeta("og:url", fmt.Sprintf("%s%s", domain, md.GetRelativeURL())),
	}

	relativeImagePath := fmt.Sprintf("/public/preview/%s.png", md.Slug)
	_, err := os.Stat(fmt.Sprintf(".%s", relativeImagePath))

	if err == nil {
		ogElements = append(
			ogElements,
			openGraphMeta("og:image", fmt.Sprintf("%s%s", domain, relativeImagePath)),
			openGraphMeta("og:image:width", "300"),
			openGraphMeta("og:image:height", "200"),
		)
	}

	return ogElements
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
