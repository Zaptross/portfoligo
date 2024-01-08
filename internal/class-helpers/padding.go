package classhelpers

import (
	g "github.com/zaptross/gorgeous"
)

func Pad(p string) string {
	className := ClassNameSanitiser("pad-" + p)
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"padding": p,
		},
	})

	return className
}

func PadL(p string) string {
	className := ClassNameSanitiser("padl-" + p)
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"padding-left": p,
		},
	})

	return className
}
func PadR(p string) string {
	className := ClassNameSanitiser("padr-" + p)
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"padding-right": p,
		},
	})

	return className
}
func PadT(p string) string {
	className := ClassNameSanitiser("padt-" + p)
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"padding-top": p,
		},
	})

	return className
}
func PadB(p string) string {
	className := ClassNameSanitiser("padb-" + p)
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"padding-bottom": p,
		},
	})

	return className
}
