"""
This script is used to transpile a dependency tree in the json format of the JS 
depency cruiser library (https://github.com/sverweij/dependency-cruiser/blob/main/doc/cli.md#json)
into a dependency tree in our own json format.

Our format is the following:
{
    "path_to_file": {
        "dependencies": [
            "path_to_file_1",
            "path_to_file_2",
            ...
        ],
        "dependents": [
            "path_to_file_3",
            "path_to_file_4",
            ...
        ]
    },
    ...
}
"""

import json
import sys

def transpile_json(input_json):
    print("Transpiling dependency tree...")
    transpiled_tree = {}

    for node in input_json["modules"]:
        file_name = node["source"]
        dependencies = []
        dependents = []

        for dep in node["dependencies"]:
            dependencies.append(dep["resolved"])

        for dep in node["dependents"]:
            dependents.append(dep)

        transpiled_tree[file_name] = {
            "dependencies": dependencies,
            "dependents": dependents
        }

    return transpiled_tree

if __name__ == "__main__":
    file_name = sys.argv[1]
    output_file_name = sys.argv[2]

    with open(file_name, "r") as f:
        input_json = json.load(f)

    transpiled_tree = transpile_json(input_json)

    with open(output_file_name, "w") as f:
        json.dump(transpiled_tree, f, indent=2)

    print(f"Transpiled dependency tree written to {output_file_name}")
