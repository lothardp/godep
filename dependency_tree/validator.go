package dependency_tree

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ValidateJSON(fileName string) error {
	jsonMap, err := openJSONFile(fileName)

	if err != nil {
		return err
	}

	err = validateDependencies(jsonMap)

	if err != nil {
		return err
	}

	println("JSON dependency file is valid")
	return nil
}

func validateDependencies(jsonMap DependencyTreeJSON) error {
	for file, deps := range jsonMap {
		for _, dependency := range deps.Dependencies {
			dep, ok := jsonMap[dependency]
			if !ok {
				fmt.Printf("Dependency %s of %s is not present in the JSON file (should be external)\n", dependency, file)
			} else if !contains(dep.Dependents, file) {
				return fmt.Errorf("File %s is not present in dependency %s dependents\n", file, dependency)
			}
		}

		for _, dependent := range deps.Dependents {
			dep, ok := jsonMap[dependent]
			if !ok {
				return fmt.Errorf("Dependent %s of %s is not present in the JSON file\n", dependent, file)
			} else if !contains(dep.Dependencies, file) {
				return fmt.Errorf("File %s is not present in dependent %s dependencies\n", file, dependent)
			}
		}
	}
	return nil
}

func contains(slice []string, element string) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}
	return false
}

func openJSONFile(fileName string) (DependencyTreeJSON, error) {
	jsonFile, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully opened ", fileName)
	defer jsonFile.Close()

	jsonBytes, _ := ioutil.ReadAll(jsonFile)

	var jsonMap DependencyTreeJSON

	err = json.Unmarshal(jsonBytes, &jsonMap)
	if err != nil {
		return nil, err
	}

	return jsonMap, nil
}
