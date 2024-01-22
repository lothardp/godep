package dependencytree

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type DependencyTreeJSON map[string]FileDeps

type FileDeps struct {
	Dependencies []string
	Dependents   []string
}

func OpenJSONFile(fileName string) (DependencyTreeJSON, error) {
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
