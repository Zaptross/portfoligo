package components

import (
	g "github.com/zaptross/gorgeous"
)

func faIcon(iconSet string, icon string, css g.CSSProps) *g.HTMLElement {
	return g.I(g.EB{
		Style: 	 css,
		ClassList: []string{"fa-" + iconSet, "fa-" + icon},
	})
}

func FAS(icon string, css g.CSSProps) *g.HTMLElement {
	return faIcon("solid", icon, css)
}

func FAB(icon string, css g.CSSProps) *g.HTMLElement {
	return faIcon("brands", icon, css)
}

func FAR(icon string, css g.CSSProps) *g.HTMLElement {
	return faIcon("regular", icon, css)
}
