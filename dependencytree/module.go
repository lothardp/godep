package dependencytree

import (
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

func (dt *DependencyTree) GetAllModules() []*Module {
	return dt.Root.getAllModules()
}

func (m *Module) getAllModules() []*Module {
	modules := make([]*Module, 0)

	modules = append(modules, m)

	for _, child := range m.Children {
		modules = append(modules, child.getAllModules()...)
	}

	return modules
}

func (dt *DependencyTree) GetAllFiles() []*File {
	return dt.Root.getAllFiles()
}

func (m *Module) getAllFiles() []*File {
	files := make([]*File, 0)

	for _, file := range m.Files {
		files = append(files, &file)
	}

	for _, child := range m.Children {
		files = append(files, child.getAllFiles()...)
	}

	return files
}
