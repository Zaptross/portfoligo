package molecules

import (
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	u "github.com/zaptross/portfoligo/internal/utils"
)

func R2C(
	col1Classes []string, col1 g.CE,
	col2Classes []string, col2 g.CE,
	containerClasses []string,
) *g.HTMLElement {
	flex1 := []string{ch.Flex(1)}
	return a.Row(g.CE{
		a.Col(col1, u.Tern(col1Classes == nil, flex1, col1Classes)),
		a.Col(col2, u.Tern(col2Classes == nil, flex1, col2Classes)),
	}, containerClasses)
}
