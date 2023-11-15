package routes

import "github.com/gin-gonic/gin"

var router = gin.Default()

func Run() {
	v1 := router.Group("/")
	AddGroupRoute(v1)

	router.Run(":2000")
}
