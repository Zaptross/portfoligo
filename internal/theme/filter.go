package theme

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

var tolerance = 1.0
var invertRange = [2]float64{0, 1}
var invertStep = 0.1
var sepiaRange = [2]float64{0, 1}
var sepiaStep = 0.1
var saturateRange = [2]float64{5, 100}
var saturateStep = 5
var hueRotateRange = [2]float64{0, 360}
var hueRotateStep = 1
var possibleColors []Color
var cache = readFilterCache()

type Color struct {
	Filters Filters
	Color   [3]float64
}

type Filters struct {
	Invert    float64
	Sepia     float64
	Saturate  float64
	HueRotate float64
}

// hexToRGB converts a hex color string to an RGB array in the range [0, 255]
func hexToRGB(hex string) [3]int {
	var r, g, b int
	fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
	return [3]int{r, g, b}
}

func GetFiltersForHex(hex string, tolerance float64) string {
	return getNewColor(hex, tolerance)
}

func filtersToCSS(filters Filters) string {
	return fmt.Sprintf(
		"invert(%f%%) sepia(%f%%) saturate(%f%%) hue-rotate(%fdeg)",
		filters.Invert*100,
		filters.Sepia*100,
		filters.Saturate,
		filters.HueRotate,
	)
}

func invertBlack(i float64) [3]float64 {
	return [3]float64{i * 255, i * 255, i * 255}
}

func filter(m [9]float64, c [3]float64) [3]float64 {
	return [3]float64{
		clamp(m[0]*c[0] + m[1]*c[1] + m[2]*c[2]),
		clamp(m[3]*c[0] + m[4]*c[1] + m[5]*c[2]),
		clamp(m[6]*c[0] + m[7]*c[1] + m[8]*c[2]),
	}
}

func clamp(value float64) float64 {
	if value > 255 {
		return 255
	} else if value < 0 {
		return 0
	}
	return value
}

func generateColors() []Color {
	var possibleColors []Color

	invert := invertRange[0]
	for invert <= invertRange[1] {
		sepia := sepiaRange[0]
		for sepia <= sepiaRange[1] {
			saturate := saturateRange[0]
			for saturate <= saturateRange[1] {
				hueRotate := hueRotateRange[0]
				for hueRotate <= hueRotateRange[1] {
					invertColor := invertBlack(invert)
					sepiaColor := filter(sepiaMatrix(sepia), invertColor)
					saturateColor := filter(saturateMatrix(saturate), sepiaColor)
					hueRotateColor := filter(hueRotateMatrix(hueRotate), saturateColor)

					colorObject := Color{
						Filters: Filters{
							Invert:    invert,
							Sepia:     sepia,
							Saturate:  saturate,
							HueRotate: hueRotate,
						},
						Color: hueRotateColor,
					}

					possibleColors = append(possibleColors, colorObject)

					hueRotate += float64(hueRotateStep)
				}
				saturate += float64(saturateStep)
			}
			sepia += sepiaStep
		}
		invert += invertStep
	}

	return possibleColors
}

func getFilters(targetColor [3]int, localTolerance float64) Filters {
	possibleColors = possibleColors[:0] // Reset possibleColors slice

	if len(possibleColors) == 0 {
		possibleColors = generateColors()
	}

	for _, color := range possibleColors {
		if math.Abs(color.Color[0]-float64(targetColor[0])) < localTolerance &&
			math.Abs(color.Color[1]-float64(targetColor[1])) < localTolerance &&
			math.Abs(color.Color[2]-float64(targetColor[2])) < localTolerance {
			return color.Filters
		}
	}

	localTolerance += tolerance
	return getFilters(targetColor, localTolerance)
}

func sepiaMatrix(s float64) [9]float64 {
	return [9]float64{
		(0.393 + 0.607*(1-s)), (0.769 - 0.769*(1-s)), (0.189 - 0.189*(1-s)),
		(0.349 - 0.349*(1-s)), (0.686 + 0.314*(1-s)), (0.168 - 0.168*(1-s)),
		(0.272 - 0.272*(1-s)), (0.534 - 0.534*(1-s)), (0.131 + 0.869*(1-s)),
	}
}

func saturateMatrix(s float64) [9]float64 {
	return [9]float64{
		0.213 + 0.787*s, 0.715 - 0.715*s, 0.072 - 0.072*s,
		0.213 - 0.213*s, 0.715 + 0.285*s, 0.072 - 0.072*s,
		0.213 - 0.213*s, 0.715 - 0.715*s, 0.072 + 0.928*s,
	}
}

func hueRotateMatrix(d float64) [9]float64 {
	cos := math.Cos(d * math.Pi / 180)
	sin := math.Sin(d * math.Pi / 180)
	a00 := 0.213 + cos*0.787 - sin*0.213
	a01 := 0.715 - cos*0.715 - sin*0.715
	a02 := 0.072 - cos*0.072 + sin*0.928

	a10 := 0.213 - cos*0.213 + sin*0.143
	a11 := 0.715 + cos*0.285 + sin*0.140
	a12 := 0.072 - cos*0.072 - sin*0.283

	a20 := 0.213 - cos*0.213 - sin*0.787
	a21 := 0.715 - cos*0.715 + sin*0.715
	a22 := 0.072 + cos*0.928 + sin*0.072

	return [9]float64{
		a00, a01, a02,
		a10, a11, a12,
		a20, a21, a22,
	}
}

func getNewColor(color string, tolerance float64) string {
	if cache[color] != "" {
		return cache[color]
	}
	// exit early if the color is a variable
	if strings.Contains(color, "var") {
		cache["var"] = color
		return color
	}
	defer writeFilterCache()

	filters := getFilters(hexToRGB(color), tolerance)
	filtersCSS := fmt.Sprintf(
		"invert(%d%%) sepia(%d%%) saturate(%d%%) hue-rotate(%ddeg)",
		int(filters.Invert*100),
		int(filters.Sepia*100),
		int(filters.Saturate*100),
		int(filters.HueRotate),
	)

	cache[color] = filtersCSS

	return filtersCSS
}

func readFilterCache() map[string]string {
	cache := make(map[string]string)

	file, err := os.ReadFile(".filter_cache.json")
	if err != nil {
		return cache
	}

	json.Unmarshal(file, &cache)

	return cache
}

func writeFilterCache() {
	file, err := os.Create(".filter_cache.json")
	if err != nil {
		return
	}

	json, err := json.Marshal(cache)
	if err != nil {
		return
	}

	io.WriteString(file, string(json))
}
