package dependencytree

type DependencyTree struct {
	Root Module
}

func NewDependencyTree(depsFile string) *DependencyTree {
	djJSON, err := OpenJSONFile(depsFile)

	if err != nil {
		panic(err)
	}

	return BuildDependencyTree(djJSON)
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
