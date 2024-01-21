package dependency_tree

import (
	"testing"
)

var (
    // slice of {{file}, {dependencies}}
	inputJSON = [][][]string{
		{{"a"}, {"b", "c", "d", "e", "mod1/a", "mod2/a", "mod3/b", "mod4/submod1/a"}},
		{{"b"}, {"c", "d", "e", "mod1/a", "mod2/a", "mod3/b", "mod4/submod1/a"}},
		{{"c"}, {"d", "e", "mod1/a", "mod2/a", "mod3/b", "mod4/submod1/a"}},
		{{"d"}, {"e", "mod1/a", "mod2/a", "mod3/b", "mod4/submod1/a"}},
		{{"e"}, {"mod1/a", "mod2/a", "mod3/b", "mod4/submod1/a"}},
		{{"mod1/a"}, {"mod1/b", "mod1/c", "mod1/d", "mod1/e", "mod2/a", "mod3/b", "mod4/submod1/a"}},
		{{"mod1/b"}, {"mod1/c", "mod1/d", "mod1/e", "mod2/a", "mod3/b", "mod4/submod1/a"}},
		{{"mod1/c"}, {"mod1/d", "mod1/e", "mod2/a", "mod3/b", "mod4/submod1/a"}},
		{{"mod1/d"}, {"mod1/e", "mod2/a", "mod3/b", "mod4/submod1/a"}},
		{{"mod1/e"}, {"mod2/a", "mod3/b", "mod4/submod1/a"}},
		{{"mod2/a"}, {"mod3/b", "mod4/submod1/a"}},
		{{"mod3/b"}, {"mod4/submod1/a"}},
		{{"mod4/submod1/a"}, {"mod4/submod1/b"}},
	}
)

func createDTJSON(inputJSON [][][]string) DependencyTreeJSON {
	dtJSON := make(DependencyTreeJSON)

	for _, fileDeps := range inputJSON {
		fileName := fileDeps[0][0]
		deps := fileDeps[1]

		dtJSON[fileName] = FileDeps{Dependencies: deps, Dependents: []string{}}
	}

	for fileName, deps := range dtJSON {
		for _, dep := range deps.Dependencies {
            tmp := dtJSON[dep]
            tmp.Dependents = append(tmp.Dependents, fileName)
            dtJSON[dep] = tmp
		}
	}

	return dtJSON
}

func TestBuildDependencyTree(t *testing.T) {
    jsonMap := createDTJSON(inputJSON)

	dt := BuildDependencyTree(jsonMap)

	if dt.Root.Name != "." {
		t.Errorf("Root name should be '.', got %s", dt.Root.Name)
	}

	if len(dt.Root.Files) != 5 {
		t.Errorf("Root should have 5 Files, got %d", len(dt.Root.Files))
	}

	if len(dt.Root.Children) != 4 {
		t.Errorf("Root should have 4 Children Modules, got %d", len(dt.Root.Children))
	}
}
