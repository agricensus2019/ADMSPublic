package routes

import (
	"html/template"
	"net/http"

	"ADMSPublic/templates"

	"github.com/gin-gonic/gin"
)

/*
Only the booklet number is unique
*/
func (srv *Server) cropping(footer string) {
	srv.router.GET("/production/cropping.html", func(c *gin.Context) {
		header, _ := templates.RenderHeader(c)
		data := gin.H{
			"Header": template.HTML(header),
			"Footer": template.HTML(footer),
		}
		c.HTML(http.StatusOK, "cropping.html", data)
	})
}
