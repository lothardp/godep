package dependency_tree

import (
	"fmt"
	"strings"
)

type DependencyTree struct {
	Root Module
}

type Module struct {
	Name     string
	Parent   *Module
	Children map[string]Module
	Files    map[string]File
}

type File struct {
	Name         string
	ParentModule *Module
	Dependencies []*File
	Dependents   []*File
}

func BuildDependencyTree(dtJSON DependencyTreeJSON) *DependencyTree {
	dt := DependencyTree{}

	dt.Root = Module{Name: ".", Parent: nil, Children: make(map[string]Module), Files: make(map[string]File)}

	for file, deps := range dtJSON {
		dt.addFile(file, deps)
	}

	return &dt
}

func (dt *DependencyTree) addFile(fileName string, deps FileDeps) {
	file := dt.getOrCreateFile(fileName)
	dt.addDeps(file, deps)
}

// Returns a pointer to the file in the tree
// If the file is not present in the tree, it creates it,
// but does not add any dependencies
func (dt *DependencyTree) getOrCreateFile(fileName string) *File {
	parentModule := dt.getOrCreateParentModule(fileName)

	file, ok := parentModule.Files[fileName]

	if !ok {
		parentModule.addFile(fileName)
	}

	return &file
}

// Adds the dependencies of the file (pointers to Files, they
// may not be present in the tree)
func (dt *DependencyTree) addDeps(file *File, deps FileDeps) {
    for _, dep := range deps.Dependencies {
        depFile := dt.getOrCreateFile(dep)
        file.Dependencies = append(file.Dependencies, depFile)
        depFile.Dependents = append(depFile.Dependents, file)
    }
}

// Returns a pointer to the module in the tree going through
// the path of the file (creating the modules if theyre not present)
func (dt *DependencyTree) getOrCreateParentModule(fileName string) *Module {
	pathOfModules := getPathOfModules(fileName)[1:] // Remove root
    currentModule := &dt.Root

    for _, module := range pathOfModules {
        currentModule = currentModule.getOrCreateChildModule(module)
    }

    return currentModule
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

// UTILS
// They could be moved out of here

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
