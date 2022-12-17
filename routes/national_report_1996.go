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
func (srv *Server) national_report_1996(footer string) {
	srv.router.GET("/production/national_report_1996.html", func(c *gin.Context) {
		header, _ := templates.RenderHeader(c)
		data := gin.H{
			"Header": template.HTML(header),
			"Footer": template.HTML(footer),
		}
		c.HTML(http.StatusOK, "national_report_1996.html", data)
	})
}
