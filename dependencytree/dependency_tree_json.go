package dependencytree

type DependencyTreeJSON map[string]FileDeps

type FileDeps struct {
	Dependencies []string
	Dependents    []string
}

