package main

import (
	"gin-mongo-crud/routers"

	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	routers.RegisterRoutes(r)
	r.Run(":8080")
}