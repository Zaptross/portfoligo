package components

import (
	g "github.com/zaptross/gorgeous"
)

func Row(children g.CE) *g.HTMLElement {
	rowClassName := "row"
	g.Class(&g.CSSClass{
		Selector: "." + rowClassName,
		Props: g.CSSProps{
			"display":        "flex",
			"flex-direction": "row",
		},
	})

	return g.Div(g.EB{
		ClassList: []string{rowClassName},
		Children:  children,
	})
}

func Col(children g.CE) *g.HTMLElement {
	colClassName := "col"
	g.Class(&g.CSSClass{
		Selector: "." + colClassName,
		Props: g.CSSProps{
			"display":        "flex",
			"flex-direction": "column",
		},
	})

	return g.Div(g.EB{
		ClassList: []string{colClassName},
		Children:  children,
	})
}
