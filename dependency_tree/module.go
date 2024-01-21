package dependency_tree

import (
	"fmt"
    "strings"
)

type Module struct {
	Name     string
	Parent   *Module
	Children map[string]Module
	Files    map[string]File
}

func (m *Module) getOrCreateChildModule(moduleName string) *Module {
	childModule, ok := m.Children[moduleName]

	if !ok {
		m.addChildModule(moduleName)
		childModule = m.Children[moduleName]
	}

	return &childModule
}

func (m *Module) addChildModule(moduleName string) {
	fmt.Println("Adding module", moduleName, "to", m.Name)
	m.Children[moduleName] = Module{
		Name:     moduleName,
		Parent:   m,
		Children: make(map[string]Module),
		Files:    make(map[string]File),
	}
}

func (m *Module) addFile(fileName string) {
	cleanFileName := strings.TrimPrefix(fileName, "./")
	m.Files[cleanFileName] = File{
		Name:         cleanFileName,
		ParentModule: m,
		Dependencies: make([]*File, 0),
		Dependents:   make([]*File, 0),
	}
}
