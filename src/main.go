package main

import (
	user "tidy/users"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	user.Routes(router)

	router.Run("localhost:8080")
}
