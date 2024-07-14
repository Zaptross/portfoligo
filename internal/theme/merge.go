package theme

import (
	r "reflect"
)

// Merge returns a new TStruct with all fields from the base struct
// overridden by the corresponding fields in the overrides struct
func Merge[TStruct any](base TStruct, overrides TStruct) TStruct {
	ref := r.New(r.TypeOf(base)).Elem()

	recursiveMergeField(ref, r.ValueOf(base), r.ValueOf(overrides))

	return ref.Interface().(TStruct)
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
