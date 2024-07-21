package organisms

import (
	"sort"

	"github.com/samber/lo"
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	c "github.com/zaptross/portfoligo/internal/components"
	m "github.com/zaptross/portfoligo/internal/components/molecules"
	t "github.com/zaptross/portfoligo/internal/types"
)

func SeriesNav(p t.PageDetails, allPages []t.PageDetails) *g.HTMLElement {
	series := findSeries(p.Series, allPages)
	previous, next := findAdjacentInSeries(p, series)

	if next == nil && previous == nil || p.Series == "" {
		return g.Empty()
	}

	if next == nil {
		return c.Row(g.CE{m.SeriesNavButton(p.Written, *previous)}, nil)
	}

	if previous == nil {
		return c.Row(g.CE{m.SeriesNavButton(p.Written, *next)}, []string{ch.JustifyContent(ch.Content.End)})
	}

	return c.Row(
		g.CE{
			m.SeriesNavButton(p.Written, *previous),
			m.SeriesNavButton(p.Written, *next),
		},
		[]string{ch.JustifyContent(ch.Content.SpaceBetween)},
	)
}

func findSeries(series string, allPages []t.PageDetails) []t.PageDetails {
	inSeries := lo.Filter(allPages, func(page t.PageDetails, _ int) bool {
		return page.Series == series && page.Series != ""
	})

	return inSeries
}

func findAdjacentInSeries(current t.PageDetails, seriesPages []t.PageDetails) (*t.PageDetails, *t.PageDetails) {
	sort.SliceStable(seriesPages, func(i, j int) bool {
		return seriesPages[i].Written.Before(seriesPages[j].Written)
	})

	for i, page := range seriesPages {
		if page.Title == current.Title {
			var prev, next *t.PageDetails = nil, nil
			if i > 0 {
				prev = &seriesPages[i-1]
			}
			if i < len(seriesPages)-1 {
				next = &seriesPages[i+1]
			}

			return prev, next
		}
	}

	return nil, nil
}
