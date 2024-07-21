package molecules

import (
	"fmt"
	"time"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	c "github.com/zaptross/portfoligo/internal/components"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	"github.com/zaptross/portfoligo/internal/theme"
	t "github.com/zaptross/portfoligo/internal/types"
	u "github.com/zaptross/portfoligo/internal/utils"
)

func SeriesNavButton(written time.Time, otherPage t.PageDetails) *g.HTMLElement {
	t := theme.UseTheme()

	after := "series-nav-after"
	g.Class(&g.CSSClass{
		Selector: "." + after + "::after",
		Props: g.CSSProps{
			"content": "' —'",
		},
	})

	before := "series-nav-before"
	g.Class(&g.CSSClass{
		Selector: "." + before + "::before",
		Props: g.CSSProps{
			"content": "'— '",
		},
	})

	directionText := u.Tern(otherPage.Written.After(written), "Next", "Earlier")
	directionClass := u.Tern(otherPage.Written.After(written), after, before)
	directionColumnClasses := u.Tern(otherPage.Written.After(written), []string{ch.TextAlign(ch.Align.End)}, nil)

	return a.LinkNav(
		c.Col(g.CE{
			a.P(g.EB{
				Text:      fmt.Sprintf("%s in %s", directionText, otherPage.Series),
				ClassList: []string{ch.MarginB(t.Spacing(1)), directionClass},
			}),
			g.Em(g.EB{
				Text:      otherPage.Title,
				ClassList: []string{ch.MarginT(t.Spacing(1))},
			}),
		}, directionColumnClasses),
		otherPage.GetRelativeURL(),
	)
}
