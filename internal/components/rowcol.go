package components

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

func Col(children g.CE, classList []string) *g.HTMLElement {
		return g.Div(g.EB{
		ClassList: append([]string{ColClass()}, classList...),
		Children:  children,
	})
}

func ColClass() string {
	colClassName := "col"
	g.Class(&g.CSSClass{
		Selector: "." + colClassName,
		Props: g.CSSProps{
			"display":        "flex",
			"flex-direction": "column",
		},
	})

	return colClassName
}