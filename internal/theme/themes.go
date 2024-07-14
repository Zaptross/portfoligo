package theme

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
	"github.com/zaptross/portfoligo/internal/types"
)

var AllThemes = []types.Theme{
	SolarizedDark,
	SolarizedLight,
	VSCodeDark,
	VSCodeLight,
}

func initTheme(selector string, t types.Theme) types.Theme {
	return Merge(types.BaseTheme, t).Setup(selector)
}

func SetupTheme(t types.Theme) {
	g.Class(&g.CSSClass{
		Selector: "." + t.Selector(),
		Include:  true,
		Raw: g.CSS(
			fmt.Sprintf(
				`
.%s {
	%s
}
				`,
				t.Selector(),
				GetThemeVariableDeclaration(t),
			)),
	})
}
