package molecules

import (
	"github.com/samber/lo"
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
	return RNC(
		[][]string{col1Classes, col2Classes},
		[]g.CE{col1, col2},
		containerClasses,
	)
}

func R3C(
	col1Classes []string, col1 g.CE,
	col2Classes []string, col2 g.CE,
	col3Classes []string, col3 g.CE,
	containerClasses []string,
) *g.HTMLElement {
	return RNC(
		[][]string{col1Classes, col2Classes, col3Classes},
		[]g.CE{col1, col2, col3},
		containerClasses,
	)
}

func RNC(
	colClasses [][]string,
	cols []g.CE,
	containerClasses []string,
) *g.HTMLElement {
	flex1 := []string{ch.Flex(1)}

	children := g.CE{}
	lo.ForEach(cols, func(col g.CE, i int) {
		children = append(children, a.Col(col, u.Tern(colClasses[i] == nil, flex1, colClasses[i])))
	})

	return a.Row(children, containerClasses)
}
