package theme

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/samber/lo"
	"github.com/zaptross/portfoligo/internal/types"
)

type ProcessorFunc func(string, []string, string, reflect.Value, map[string]string) string

// GetThemeVariableDeclaration returns a string of CSS variable declarations
// eg. "--colors-text-primary: #000000;\n--colors-text-secondary: #000000;\n..."
func GetThemeVariableDeclaration(t types.Theme) string {
	t.Colors.Icons = t.Colors.Text
	t.Colors.Filters = Merge(getPalletteFilters(t), t.Colors.Filters)

	return strings.Join(processTheme(t, getVariableDeclaration, getPalletteMap(t.Colors.Pallette)), "\n")
}

func processTheme(t types.Theme, variableProcessor ProcessorFunc, palMap map[string]string) []string {
	variables := []string{}
	ref := reflect.ValueOf(&t).Elem()

	for fi := 0; fi < ref.NumField(); fi++ {
		if ref.Type().Field(fi).IsExported() {
			variables = append(variables, processToVariables(variableProcessor, ref.Field(fi), ref.Type().Field(fi).Name, []string{}, palMap)...)
		}
	}

	return variables
}

func processToVariables(variableProcessor ProcessorFunc, v reflect.Value, keyName string, parentKeys []string, palMap map[string]string) []string {
	variables := []string{}
	pks := append(parentKeys, keyName)

	if v.Kind() == reflect.Struct {
		for fi := 0; fi < v.NumField(); fi++ {
			variables = append(variables, processToVariables(variableProcessor, v.Field(fi), v.Type().Field(fi).Name, pks, palMap)...)
		}

		return variables
	}

	cssVar := variableProcessor(keyName, parentKeys, v.String(), v, palMap)

	if lo.Contains(parentKeys, "Icons") {
		cssVar = strings.Replace(cssVar, "pallette", "filters", 1)
	}

	return []string{cssVar}
}

func getVariableDeclaration(fieldName string, parentKeys []string, value string, v reflect.Value, palMap map[string]string) string {
	vn := getVariableName(fieldName, parentKeys, value, v)
	if palMap[value] != "" && !strings.Contains(palMap[value], vn) {
		return fmt.Sprintf("%s: %s;", vn, palMap[value])
	}
	return fmt.Sprintf("%s: %s;", vn, value)
}
func getVariableName(fieldName string, parentKeys []string, _ string, _ reflect.Value) string {
	return strings.ToLower(fmt.Sprintf("--%s-%s", strings.Join(parentKeys, "-"), fieldName))
}

func getPalletteFilters(t types.Theme) types.Pallette {
	filters := reflect.New(reflect.TypeOf(types.Pallette{})).Elem()
	colours := reflect.ValueOf(t.Colors.Pallette)

	for fi := 0; fi < filters.NumField(); fi++ {
		// Convert each colour that exists to a filter
		if colours.Field(fi).String() != "" {
			filters.Field(fi).SetString(GetFiltersForHex(colours.Field(fi).String(), 1))
		}
	}

	// apply changes, overriding calculated filters with user defined filters
	return filters.Interface().(types.Pallette)
}

func getPalletteMap(p types.Pallette) map[string]string {
	v := UseTheme().Colors.Pallette
	m := map[string]string{}

	rp := reflect.ValueOf(p)
	rv := reflect.ValueOf(v)
	for fi := 0; fi < rp.NumField(); fi++ {
		// eg: m["#000000"] = "--colors-pallette-base03"
		m[rp.Field(fi).String()] = rv.Field(fi).String()
	}

	return m
}
