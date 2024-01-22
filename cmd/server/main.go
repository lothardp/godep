package main

import (
	"flag"
	"lothardp/godep/dependencytree"
	"lothardp/godep/server"
)

func main() {
	depsFile := flag.String("deps-file", "", "Path to the json dependencies file")
	port := flag.String("port", ":8080", "Port to run the server on")
	flag.Parse()

	if *depsFile == "" {
		panic("Missing dependencies file")
	}

	dt := dependencytree.NewDependencyTree(*depsFile)

	r := server.SetupServer(dt)
	r.Run(*port)
}
