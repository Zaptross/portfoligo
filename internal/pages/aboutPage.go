package pages

import (
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	c "github.com/zaptross/portfoligo/internal/components"
	p "github.com/zaptross/portfoligo/internal/provider"
	"github.com/zaptross/portfoligo/internal/types"
)

var (
	_ = registerPage(types.PageDetails{
		Title:       "About",
		Description: "About Me",
		Slug:        "about",
		Type:        TYPE_ROOT,
		Content: func(_ types.PageDetails) *g.HTMLElement {
			t := p.ThemeProvider.GetTheme()
			return g.Div(g.EB{
				ClassList: []string{ch.FlexCol(), ch.JustifyContent(ch.Content.SpaceBetween)},
				Children: []*g.HTMLElement{
					g.H1(g.EB{
						Text:      "Hi, I'm Matt.",
						ClassList: []string{ch.Margin("0.5rem auto"), ch.FontColor(t.Green)},
					}),
					c.P(g.EB{
						Text: "I am a full-stack polyglot software engineer with a passion for learning new skills and technologies.",
					}),
					c.P(g.EB{
						Text: "Some incredibly kind things people have said about me:",
					}),
					c.AuthorQuote(
						"Matt displays the intelligence, creativity and ingenuity of a seasoned veteran, despite [being in] his first formal role in software engineering. His passion and determination to do [things] the right way shines in all that he does. That skill and style is exceeded only by his humanity. Pricey is simply one of the most kind, compassionate and generous humans I have ever known - an example to others and one that others thus try to live to. It was an honour to exist beside him in our team. If you have the fortune to get his attention, I strongly recommend him!",
						"Damian, AppSec Lead at Swyftx",
						t.Green,
						[]string{},
					),
					c.AuthorQuote(
						"Matthew has demonstrated exceptional skills as a team player, working well with others to complete assigned projects to a high quality. Matthew would be a tremendous asset for your company and has my highest recommendation.",
						"Robert, Sessional Academic at QUT",
						t.Cyan,
						[]string{},
					),
					c.AuthorQuote(
						"Matt has always worked above expectations as a part of our team [and] has always been one of the most pleasant individuals I've had under my employ.",
						"Dale, Store Manager at Dominos",
						t.Blue,
						[]string{},
					),
				}},
			)
		},
	})
)
