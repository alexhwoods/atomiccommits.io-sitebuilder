package main

import (
	"atomiccommits.io/sitebuilder/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	controllers.Routes(router)

	router.Run(":3000")
}
