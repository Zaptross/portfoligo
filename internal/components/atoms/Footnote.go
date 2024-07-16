package atoms

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	"github.com/zaptross/portfoligo/internal/theme"
)

func Footnote(index int) *g.HTMLElement {
	t := theme.UseTheme()
	hash := fmt.Sprintf("[%d]", index)
	return GotoHashlink(
		g.Span(g.EB{Text: hash, ClassList: []string{
			ch.MarginL(t.Spacing(1)),
			ch.MarginR(t.Spacing(1)),
			ch.Superscript(),
		}}),
		hash,
	)
}
