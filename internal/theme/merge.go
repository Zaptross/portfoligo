package theme

import (
	r "reflect"

	T "github.com/zaptross/portfoligo/internal/types"
)

// Merge returns a new Theme struct with all fields from the base Theme struct
// overridden by the corresponding fields in the overrides Theme struct
func Merge(base T.Theme, overrides T.Theme) T.Theme {
	ref := r.New(r.TypeOf(T.Theme{})).Elem()

	recursiveMergeField(ref, r.ValueOf(base), r.ValueOf(overrides))

	return ref.Interface().(T.Theme)
}

func recursiveMergeField(out, base, overrides r.Value) {
	for fi := 0; fi < out.NumField(); fi++ {
		if !out.Type().Field(fi).IsExported() {
			continue
		}

		if out.Field(fi).Kind() == r.Struct {
			recursiveMergeField(out.Field(fi), base.Field(fi), overrides.Field(fi))
		} else {
			if overrides.Field(fi).String() != "" {
				out.Field(fi).Set(overrides.Field(fi))
			} else {
				out.Field(fi).Set(base.Field(fi))
			}
		}
	}
}
