package main

import (
	"github.com/samber/lo"
)

func main() {
	println("gen.go")

	setupDirectories()

	fps := getPublicFilesInSegments()
	groups := groupBySubpath(fps)
	lo.ForEach(groups, func(group [][]string, _ int) {
		writeGroupToFile(group)
	})
}
