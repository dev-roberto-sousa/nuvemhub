package main

import (
	"github.com/dev-roberto-sousa/nuvemhub/internals/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", handlers.GetAllUsers)

	r.Run()
}
