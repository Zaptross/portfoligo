package classhelpers

import (
	g "github.com/zaptross/gorgeous"
)

func Flex() string {
	className := "flex"
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"display": "flex",
		},
	})

	return className
}

func FlexRow() string {
	className := "flex-row"
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"display":        "flex",
			"flex-direction": "row",
		},
	})

	return className
}

func FlexCol() string {
	className := "flex-col"
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"display":        "flex",
			"flex-direction": "column",
		},
	})

	return className
}

func FlexGrow() string {
	className := "flex-grow"
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"flex-grow": "1",
		},
	})

	return className
}

func FlexWrap() string {
	className := "flex-wrap"
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"flex-wrap": "wrap",
		},
	})

	return className
}
