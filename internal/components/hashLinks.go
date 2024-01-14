package components

import (
	_ "embed"

	g "github.com/zaptross/gorgeous"
)

//go:embed hashLinks.js
var hashLinks string

func Hashlink(children g.CE, hash string, classList []string) *g.HTMLElement {
	g.Service("hashLinks", g.JavaScript(hashLinks))

	// TODO - copy full hashlink to clipboard

	return g.A(g.EB{
		Props: g.Props{
			"name": hash,
		},
		ClassList: classList,
		Children:  children,
	})
}

func GotoHashlink(child *g.HTMLElement, hash string) *g.HTMLElement {
	return link(child, "#"+hash, "hash", true, false)
}
