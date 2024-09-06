package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine)  {
	api := r.Group("/api")
	{
		v1:=api.Group("/v1")
		{
			v1.GET("/hello",hello)
		}
	}
}

func hello(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{"message":"hi"})
	
}