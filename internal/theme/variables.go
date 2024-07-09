package theme

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/zaptross/portfoligo/internal/types"
)

type ProcessorFunc func(string, []string, string, reflect.Value) string

// GetThemeVariableDeclaration returns a string of CSS variable declarations
// eg. "--colors-text-primary: #000000;\n--colors-text-secondary: #000000;\n..."
func GetThemeVariableDeclaration(t types.Theme) string {
	return strings.Join(processTheme(t, getVariableDeclaration), "\n")
}

func processTheme(t types.Theme, variableProcessor ProcessorFunc) []string {
	variables := []string{}
	ref := reflect.ValueOf(&t).Elem()

	for fi := 0; fi < ref.NumField(); fi++ {
		if ref.Type().Field(fi).IsExported() {
			variables = append(variables, processToVariables(variableProcessor, ref.Field(fi), ref.Type().Field(fi).Name, []string{})...)
		}
	}

	return variables
}

func processToVariables(variableProcessor ProcessorFunc, v reflect.Value, keyName string, parentKeys []string) []string {
	variables := []string{}

	if v.Kind() == reflect.Struct {
		for fi := 0; fi < v.NumField(); fi++ {
			variables = append(variables, processToVariables(variableProcessor, v.Field(fi), v.Type().Field(fi).Name, append(parentKeys, keyName))...)
		}

		return variables
	}

	return []string{variableProcessor(keyName, parentKeys, v.String(), v)}
}

func getVariableDeclaration(fieldName string, parentKeys []string, value string, v reflect.Value) string {
	return fmt.Sprintf("%s: %s;", getVariableName(fieldName, parentKeys, value, v), value)
}
func getVariableName(fieldName string, parentKeys []string, _ string, _ reflect.Value) string {
	return strings.ToLower(fmt.Sprintf("--%s-%s", strings.Join(parentKeys, "-"), fieldName))
}
