package classhelpers

import (
	"strings"
)

// Class Helpers is a collection of functions that return class names for
// common CSS properties. These functions also register the class names with
// the global class registry.

type strReplacer struct {
	str string
}

func (sr *strReplacer) replace(old, new string) {
	sr.str = strings.ReplaceAll(sr.str, old, new)
}
func (sr strReplacer) String() string {
	return sr.str
}

func ClassNameSanitiser(className string) string {
	sr := strReplacer{className}
	sr.replace("%", "pct")
	sr.replace("/", "-")
	sr.replace("(", "-")
	sr.replace(")", "-")
	sr.replace(",", "-")
	sr.replace("#", "-")
	sr.replace(".", "-")
	sr.replace(" ", "_")

	return sr.String()
}
