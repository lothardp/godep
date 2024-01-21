package dependencytree

import (
	"strings"
)

func getPathOfModules(fileName string) []string {
	fileName = strings.TrimPrefix(fileName, "./")
	path := []string{"."}
	lastMod := "."

	for i, pathElement := range strings.Split(fileName, "/") {
		if i == len(fileName)-1 {
			break
		}

		lastMod = lastMod + "/" + pathElement
		path = append(path, lastMod)
	}

	return path
}
