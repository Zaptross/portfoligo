package molecules

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/samber/lo"
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	c "github.com/zaptross/portfoligo/internal/components"
	"github.com/zaptross/portfoligo/internal/components/atoms"
	"github.com/zaptross/portfoligo/internal/theme"
)

type ReferenceProps struct {
	Title     string
	Link      string
	Date      time.Time
	Retrieved time.Time
	Index     int
}

func Reference(props ReferenceProps) *g.HTMLElement {
	t := theme.UseTheme()
	referenceClass := "footnotes-reference"
	g.Class(&g.CSSClass{
		Selector: "." + referenceClass + " > * + *",
		Include:  true,
		Props: g.CSSProps{
			"padding-left": t.Spacing(1),
		},
	})
	return c.P(g.EB{
		ClassList: []string{referenceClass},
		Children: g.CE{
			tern(props.Index > 0, atoms.Hashlink(g.CE{g.Span(g.EB{Text: fmt.Sprintf("[%d]", props.Index)})}, fmt.Sprintf("[%d]", props.Index), nil), g.Empty()),
			g.Span(g.EB{Text: props.Title}),
			g.Span(g.EB{Text: fmt.Sprintf("(%s)", props.Date.Format("2006"))}),
			tern(!props.Retrieved.IsZero(), g.Span(g.EB{Text: fmt.Sprintf("Retrieved %s, from:", props.Retrieved.Format("2006-01-02"))}), g.Empty()),
			atoms.TextLink(props.Link, props.Link),
		},
	})
}

func Footnote(index int) *g.HTMLElement {
	t := theme.UseTheme()
	hash := fmt.Sprintf("[%d]", index)
	return atoms.GotoHashlink(
		g.Span(g.EB{Text: hash, ClassList: []string{
			ch.MarginL(t.Spacing(1)),
			ch.MarginR(t.Spacing(1)),
			ch.Superscript(),
		}}),
		hash,
	)
}

func tern[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

type FootnoteCollector struct {
	footnotes        []ReferenceProps
	orderedFootnotes []int
}

func NewFootnoteCollector() FootnoteCollector {
	return FootnoteCollector{
		footnotes:        []ReferenceProps{},
		orderedFootnotes: []int{},
	}
}

func (fnc *FootnoteCollector) Add(ref ReferenceProps) int {
	index := lo.IndexOf(fnc.footnotes, ref)

	if index >= 0 {
		return index
	}

	fnc.footnotes = append(fnc.footnotes, ref)

	return len(fnc.footnotes) - 1
}

func (fnc *FootnoteCollector) Ref(index int) *g.HTMLElement {
	return Footnote(fnc.handleReference(index) + 1)
}
func (fnc *FootnoteCollector) RefLink(child *g.HTMLElement, index int) *g.HTMLElement {
	hash := fmt.Sprintf("[%d]", fnc.handleReference(index)+1)
	return atoms.GotoHashlink(
		child,
		hash,
	)
}
func (fnc *FootnoteCollector) handleReference(index int) int {
	var orderedIndex int
	if !lo.Contains(fnc.orderedFootnotes, index) {
		fnc.orderedFootnotes = append(fnc.orderedFootnotes, index)
		orderedIndex = len(fnc.orderedFootnotes) - 1
	}

	return orderedIndex
}

func (fnc *FootnoteCollector) GenerateReferences() []*g.HTMLElement {
	return lo.Map(fnc.orderedFootnotes, fnc.refFromIndex)
}

func (fnc *FootnoteCollector) refFromIndex(index int, order int) *g.HTMLElement {
	ref := fnc.footnotes[index]
	ref.Index = order + 1
	return Reference(ref)
}
