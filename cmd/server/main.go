package main

import (
	"lothardp/godep/server"
)

func main() {
	r := server.SetupRouter()
	r.Run(":8080")
}
