package molecules

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

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
	notReferenced, _ := lo.Difference(
		lo.Map(fnc.footnotes, func(_ atoms.ReferenceProps, i int) int { return i }),
		fnc.orderedFootnotes,
	)

	slices.SortFunc(notReferenced, func(i, j int) int {
		return strings.Compare(fnc.footnotes[i].Title, fnc.footnotes[j].Title)
	})

	return append(
		lo.Map(fnc.orderedFootnotes, fnc.refFromIndex),
		lo.Map(notReferenced, func(ref int, _ int) *g.HTMLElement { return atoms.Reference(fnc.footnotes[ref]) })...,
	)
}

func (fnc *FootnoteCollection) refFromIndex(index int, order int) *g.HTMLElement {
	ref := fnc.footnotes[index]
	ref.Index = order + 1
	return atoms.Reference(ref)
}
