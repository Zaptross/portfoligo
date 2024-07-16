package classhelpers

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
)

func DisplayFlex() string {
	className := "flex"
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"display": "flex",
		},
	})

	return className
}

func Flex(i int) string {
	digit := fmt.Sprintf("%d", i)
	className := "flex-" + digit
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"flex": digit,
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

func FlexGrow(i int) string {
	digit := fmt.Sprintf("%d", i)
	className := "flex-grow-" + digit
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"flex-grow": digit,
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
