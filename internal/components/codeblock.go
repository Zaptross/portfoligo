package components

import (
	"fmt"
	"html"
	"os"

	g "github.com/zaptross/gorgeous"
	prv "github.com/zaptross/portfoligo/internal/provider"
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

	theme := prv.ThemeProvider.GetTheme()

	g.Class(&g.CSSClass{
		Selector: ".codeblock",
		Props: g.CSSProps{
			"background-color": theme.Base02 + " !important",
			"height":           "100%",
			"border":           "1px solid",
			"border-color":     theme.Green,
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
			"color":       theme.Base1 + " !important",
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
