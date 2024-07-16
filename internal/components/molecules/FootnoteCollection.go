package molecules

import (
	_ "embed"
	"fmt"

	"github.com/samber/lo"
	g "github.com/zaptross/gorgeous"
	"github.com/zaptross/portfoligo/internal/components/atoms"
)

type FootnoteCollection struct {
	footnotes        []atoms.ReferenceProps
	orderedFootnotes []int
}

func NewFootnoteCollector() FootnoteCollection {
	return FootnoteCollection{
		footnotes:        []atoms.ReferenceProps{},
		orderedFootnotes: []int{},
	}
}

func (fnc *FootnoteCollection) Add(ref atoms.ReferenceProps) int {
	index := lo.IndexOf(fnc.footnotes, ref)

	if index >= 0 {
		return index
	}

	fnc.footnotes = append(fnc.footnotes, ref)

	return len(fnc.footnotes) - 1
}

func (fnc *FootnoteCollection) Ref(index int) *g.HTMLElement {
	return atoms.Footnote(fnc.handleReference(index) + 1)
}
func (fnc *FootnoteCollection) RefLink(child *g.HTMLElement, index int) *g.HTMLElement {
	hash := fmt.Sprintf("[%d]", fnc.handleReference(index)+1)
	return atoms.GotoHashlink(
		child,
		hash,
	)
}
func (fnc *FootnoteCollection) handleReference(index int) int {
	var orderedIndex int
	if !lo.Contains(fnc.orderedFootnotes, index) {
		fnc.orderedFootnotes = append(fnc.orderedFootnotes, index)
		orderedIndex = len(fnc.orderedFootnotes) - 1
	}

	return orderedIndex
}

func (fnc *FootnoteCollection) GenerateReferences() []*g.HTMLElement {
	return lo.Map(fnc.orderedFootnotes, fnc.refFromIndex)
}

func (fnc *FootnoteCollection) refFromIndex(index int, order int) *g.HTMLElement {
	ref := fnc.footnotes[index]
	ref.Index = order + 1
	return atoms.Reference(ref)
}
