package molecules

import (
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	c "github.com/zaptross/portfoligo/internal/components"
	a "github.com/zaptross/portfoligo/internal/components/atoms"
	"github.com/zaptross/portfoligo/internal/theme"
	"github.com/zaptross/portfoligo/internal/types"
)

type CaptionedImageProps struct {
	Caption          string
	Src              string
	Page             types.PageDetails
	ContainerClasses []string
	ImgClasses       []string
	CaptionClasses   []string
}

func CaptionedImage(props CaptionedImageProps) *g.HTMLElement {
	t := theme.UseTheme()

	return g.CustomElement(
		"figure",
		true,
		g.EB{
			ClassList: props.ContainerClasses,
			Children: g.CE{
				a.ClickZoom(
					c.Col(g.CE{
						ProjectImage(props.Page, props.Src, append([]string{ch.W("100%"), ch.H("100%"), ch.MxW("80vw"), ch.MxH("80vh"), ch.Aspect("preserve")}, props.ImgClasses...)),
						g.CustomElement(
							"figcaption",
							true,
							g.EB{
								Text:      props.Caption,
								ClassList: append([]string{ch.FontColor(t.Colors.Text.Secondary), ch.TextAlign(ch.Align.Center), ch.PadT(t.Spacing(1))}, props.CaptionClasses...),
							},
						),
					}, nil),
					nil,
				),
			},
		},
	)
}
