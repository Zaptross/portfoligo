package classhelpers

import (
	g "github.com/zaptross/gorgeous"
)

func Margin(m string) string {
	className := ClassNameSanitiser("m-" + m)
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"margin": m,
		},
	})

	return className
}

func MarginL(m string) string {
	className := ClassNameSanitiser("ml-" + m)
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"margin-left": m,
		},
	})

	return className
}
func MarginR(m string) string {
	className := ClassNameSanitiser("mr-" + m)
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"margin-right": m,
		},
	})

	return className
}
func MarginT(m string) string {
	className := ClassNameSanitiser("mt-" + m)
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"margin-top": m,
		},
	})

	return className
}
func MarginB(m string) string {
	className := ClassNameSanitiser("mb-" + m)
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"margin-bottom": m,
		},
	})

	return className
}

func MarginZA() string {
	return Margin("0 auto")
}
func MarginAZ() string {
	return Margin("auto 0")
}
func MarginAA() string {
	return Margin("auto")
}
