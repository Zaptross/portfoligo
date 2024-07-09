package theme

import (
	"fmt"
	r "reflect"

	"github.com/samber/lo"
	T "github.com/zaptross/portfoligo/internal/types"
)

// UseTheme returns a new Theme struct with all fields set to CSS variable references
// eg. Theme{Colors: {Text: {Primary: "var(--colors-text-primary)", Secondary: "var(--colors-text-secondary)", ...}, ...}}
func UseTheme() T.Theme {
	ref := r.New(r.TypeOf(T.Theme{})).Elem()

	recursiveSetField(ref, []r.StructField{})

	return ref.Interface().(T.Theme)
}

func recursiveSetField(ref r.Value, parents []r.StructField) {
	for fi := 0; fi < ref.NumField(); fi++ {
		if !ref.Type().Field(fi).IsExported() {
			continue
		}

		if ref.Field(fi).Kind() == r.Struct {
			recursiveSetField(ref.Field(fi), append(parents, ref.Type().Field(fi)))
		} else {
			parentKeys := lo.Reduce(
				parents,
				func(acc []string, cur r.StructField, _ int) []string {
					return append(acc, cur.Name)
				},
				[]string{},
			)
			setVariableReference(ref.Type().Field(fi).Name, parentKeys, ref.Field(fi))
		}
	}
}

func setVariableReference(fieldName string, parentKeys []string, v r.Value) {
	fn := fmt.Sprintf("var(%s)", getVariableName(fieldName, parentKeys, "", v))

	if v.CanSet() {
		v.SetString(fn)
		return
	}
}
