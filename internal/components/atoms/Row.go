package atoms

import (
	g "github.com/zaptross/gorgeous"
)

func Row(children g.CE, classList []string) *g.HTMLElement {
	return g.Div(g.EB{
		ClassList: append([]string{RowClass()}, classList...),
		Children:  children,
	})
}

func RowClass() string {
	rowClassName := "row"
	g.Class(&g.CSSClass{
		Selector: "." + rowClassName,
		Props: g.CSSProps{
			"display":        "flex",
			"flex-direction": "row",
		},
	})

	return rowClassName
}
