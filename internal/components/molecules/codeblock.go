package molecules

import (
	"fmt"
	"html"
	"os"

	g "github.com/zaptross/gorgeous"
	"github.com/zaptross/portfoligo/internal/theme"
)

type CodeblockProps struct {
	Code     string
	FileName string
	FilePath string
	Language string
}

func Codeblock(props CodeblockProps) *g.HTMLElement {
	if (props.Code == "" && props.FilePath == "") || props.Language == "" {
		panic("Codeblock must have a language and code or a file path")
	}

	t := theme.UseTheme()

	g.Class(&g.CSSClass{
		Selector: ".codeblock",
		Props: g.CSSProps{
			"background-color": t.Colors.Background.Primary + " !important",
			"height":           "100%",
			"border":           "1px solid",
			"border-color":     t.Colors.Background.Secondary,
			"padding":          "0.5rem",
			"font-family":      "monospace !important",
		},
	})
	g.Class(&g.CSSClass{
		Selector: "pre",
		Include:  true,
		Props: g.CSSProps{
			"background":  "none !important",
			"text-shadow": "none !important",
		},
	})
	g.Class(&g.CSSClass{
		Selector: "code",
		Include:  true,
		Props: g.CSSProps{
			"color":       t.Colors.Text.Secondary + " !important",
			"background":  "none !important",
			"text-shadow": "none !important",
		},
	})
	g.Class(&g.CSSClass{
		Selector: ".operator",
		Include:  true,
		Props: g.CSSProps{
			"background": "none !important",
		},
	})

	var codeString string
	if props.Code != "" {
		codeString = fmt.Sprintf(`//%s\n%s`, props.FileName, props.Code)
	} else {
		codeString = loadCodeFile(props)
	}

	return g.Div(
		g.EB{
			Children: g.CE{
				g.Pre(g.EB{Children: g.CE{
					g.CustomElement("code", true, g.EB{
						Children:  g.CE{g.RawText(codeString)},
						ClassList: []string{"language-" + props.Language},
					}),
				}}),
			},
			ClassList: []string{"codeblock"},
		},
	)
}

func loadCodeFile(props CodeblockProps) string {
	raw, err := os.ReadFile(props.FilePath)

	if err != nil {
		panic(err)
	}

	codeString := html.EscapeString(string(raw))

	return "// " + props.FileName + "\n" + codeString
}
