package server

import (
	"lothardp/godep/dependencytree"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

type Server struct {
	dt *dependencytree.DependencyTree
}

func SetupServer(dt *dependencytree.DependencyTree) *gin.Engine {
	s := Server{dt: dt}

	r := gin.Default()

	r.GET("/summary", func(c *gin.Context) {
		c.JSON(http.StatusOK, s.GetSummary())
	})

	return r
}
