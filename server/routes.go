// This file is autogenerated by Golem
// Do NOT make modifications, they will be lost
package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Routes() {
	s.Router.GET("/", homeHandler)

}

func homeHandler(c *gin.Context) {
	Home(c)
}

func Home(c *gin.Context) {
	c.String(http.StatusOK, "home")
}
