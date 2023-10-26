package main

import (
	"os"

	g "github.com/zaptross/gorgeous"
)

func createDistDirectories() {
	os.Mkdir("dist", 0755)
	os.Mkdir("dist/public", 0755)
	os.Mkdir("dist/blog", 0755)
	os.Mkdir("dist/project", 0755)
}

func writeRenderedHTML(path string, rendered *g.RenderedHTML) {
	path = "dist" + path

	os.Mkdir(path, 0755)

	indexPath := path + "/index.html"
	os.WriteFile(indexPath, []byte(rendered.Document), 0644)
	if len(rendered.Style) > 0 {
		stylePath := path + "/style.css"
		os.WriteFile(stylePath, []byte(rendered.Style), 0644)
	}
	if len(rendered.Script) > 0 {
		scriptPath := path + "/script.js"
		os.WriteFile(scriptPath, []byte(rendered.Script), 0644)
	}
}

func copyPublicToDist() {
	files, err := os.ReadDir("public")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileContents, fcErr := os.ReadFile("public/" + file.Name())

		if fcErr != nil {
			panic(fcErr)
		}

		os.WriteFile("dist/public/"+file.Name(), fileContents, 0644)
	}
}
