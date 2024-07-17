package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var desiredDirectories = []string{
	"blog",
	"projects",
}

func filesMain() {
	setupDirectories()

	fps := getPublicFilesInSegments()
	groups := groupBySubpath(fps)
	lo.ForEach(groups, func(group [][]string, _ int) {
		writeGroupToFile(group)
		println("Generated file for", group[0][1])
	})
}

func setupDirectories() {
	err := os.Mkdir("../internal/generated", 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	err = os.Mkdir("../internal/generated/files", 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
}

func getPublicFilesInSegments() [][]string {
	filePathsInSegments := [][]string{}

	err := filepath.Walk("../cmd/public", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strContainsAny(path, desiredDirectories) {
			return nil
		}

		if !info.IsDir() {
			trimmedPath := trimPath(path)
			segments := strings.Split(trimmedPath, string(os.PathSeparator))
			filePathsInSegments = append(filePathsInSegments, segments)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return filePathsInSegments
}

func groupBySubpath(paths [][]string) [][][]string {
	groupedPaths := make([][][]string, 0)

	for _, path := range paths {
		if len(groupedPaths) == 0 {
			groupedPaths = append(groupedPaths, [][]string{path})
			continue
		}

		if groupedPaths[len(groupedPaths)-1][0][1] == path[1] {
			groupedPaths[len(groupedPaths)-1] = append(groupedPaths[len(groupedPaths)-1], path)
		} else {
			groupedPaths = append(groupedPaths, [][]string{path})
		}
	}

	return groupedPaths
}

func writeGroupToFile(paths [][]string) {
	fileName := kebabToCamelCase(paths[0][1]) + ".go"

	file, err := os.Create("../internal/generated/files/" + fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.WriteString("package files\n\n")
	file.WriteString(createStructFromGroup(paths) + "\n\n")
	file.WriteString(createVarFromGroup(paths) + "\n")
}

func createStructFromGroup(paths [][]string) string {
	var structFields []string

	for _, path := range paths {
		structFields = append(structFields, "\t// "+path[len(path)-1])
		structFields = append(structFields, "\t"+createFieldName(path[2:])+" string")
	}

	return "type " + createStructName(paths[0][1]) + " struct {\n" + strings.Join(structFields, "\n") + "\n}"
}
func createVarFromGroup(paths [][]string) string {
	var structFields []string

	for _, path := range paths {
		structFields = append(structFields, "\t"+createFieldName(path[2:])+": \"/"+path[len(path)-1]+"\",")
	}

	return "var " + kebabToCamelCase(paths[0][1]) + " = " + createStructName(paths[0][1]) + "{\n" + strings.Join(structFields, "\n") + "\n}"
}

func createStructName(path string) string {
	return lowerFirstChar(kebabToCamelCase(path) + "Files")
}
func createFieldName(path []string) string {
	return stripExtension(kebabToCamelCase(strings.Join(path, "-")))
}

func kebabToCamelCase(s string) string {
	segments := strings.Split(s, "-")
	for i, segment := range segments {
		segments[i] = cases.Title(language.English).String(segment)
	}

	return strings.Join(segments, "")
}
func stripExtension(s string) string {
	return strings.TrimSuffix(s, filepath.Ext(s))
}
func lowerFirstChar(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}
func trimPath(path string) string {
	return strings.Replace(path, "../cmd/public/", "", 1)
}

func strContainsAny(s string, substrings []string) bool {
	for _, substring := range substrings {
		if strings.Contains(s, substring) {
			return true
		}
	}

	return false
}