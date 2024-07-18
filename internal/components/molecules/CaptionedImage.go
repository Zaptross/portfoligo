package molecules

import (
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
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

	return a.ClickZoom(
		g.CustomElement(
			"figure",
			true,
			g.EB{
				ClassList: props.ContainerClasses,
				Children: g.CE{
					ProjectImage(props.Page, props.Src, append([]string{ch.W("100%")}, props.ImgClasses...)),
					g.CustomElement(
						"figcaption",
						true,
						g.EB{
							Text:      props.Caption,
							ClassList: append([]string{ch.FontColor(t.Colors.Text.Primary), ch.TextAlign(ch.Align.Center)}, props.CaptionClasses...),
						},
					),
				},
			},
		),
	)
}
