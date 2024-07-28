package specific

import (
	g "github.com/zaptross/gorgeous"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	m "github.com/zaptross/portfoligo/internal/components/molecules"
)

func AboutBlock() *g.HTMLElement {
	return a.Col(g.CE{
		m.SectionHeader(m.SectionHeaderProps{
			TagKey: "about",
			Text:   "Hi 👋, I'm Matt.",
			Header: g.H1,
		}),
		a.P(g.EB{
			Text: "I am a full-stack polyglot software engineer with a passion for learning new skills and technologies.",
		}),
		a.P(g.EB{
			Text: "In my spare time I love to play video- and board- games, develop fun ideas and projects, and work on hobbyist electronics projects.",
		}),
	}, nil)
}
