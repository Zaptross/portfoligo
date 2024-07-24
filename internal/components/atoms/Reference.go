package atoms

import (
	"fmt"
	"time"

	g "github.com/zaptross/gorgeous"
	"github.com/zaptross/portfoligo/internal/theme"
	u "github.com/zaptross/portfoligo/internal/utils"
)

type ReferenceProps struct {
	Title     string
	Link      string
	Date      time.Time
	Retrieved time.Time
	Index     int
}

func Reference(props ReferenceProps) *g.HTMLElement {
	t := theme.UseTheme()
	referenceClass := "footnotes-reference"
	g.Class(&g.CSSClass{
		Selector: "." + referenceClass + " > * + *",
		Include:  true,
		Props: g.CSSProps{
			"padding-left": t.Spacing(1),
		},
	})
	return P(g.EB{
		ClassList: []string{referenceClass},
		Children: g.CE{
			u.Tern(props.Index > 0, Hashlink(g.CE{g.Span(g.EB{Text: fmt.Sprintf("[%d]", props.Index)})}, fmt.Sprintf("[%d]", props.Index), nil), g.Empty()),
			g.Span(g.EB{Text: props.Title}),
			g.Span(g.EB{Text: fmt.Sprintf("(%s)", props.Date.Format("2006"))}),
			u.Tern(!props.Retrieved.IsZero(), g.Span(g.EB{Text: fmt.Sprintf("Retrieved %s, from:", props.Retrieved.Format("2006-01-02"))}), g.Empty()),
			TextLink(props.Link, props.Link),
		},
	})
}
