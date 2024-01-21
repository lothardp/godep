package dependency_tree

type File struct {
	Name         string
	ParentModule *Module
	Dependencies []*File
	Dependents   []*File
}
