package main


import (
    "os"
    dt "lothardp/godep/dependency_tree"
)

func main() {
    if len(os.Args) < 2 {
        panic("Please provide a file name")
    }
    fileName := os.Args[1]
    err := dt.ValidateJSON(fileName)

    if err != nil {
        panic(err)
    }
}
