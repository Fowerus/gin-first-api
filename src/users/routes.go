package users

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	router.GET("/users", Users)
	router.POST("/users", Create)
	router.PATCH("/users/:id", Update)
	router.PUT("/users/:id", Update)
	router.DELETE("users/:id", Delete)
}
