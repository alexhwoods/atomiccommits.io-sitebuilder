package main

import (
	"atomiccommits.io/sitebuilder/controllers"
	"atomiccommits.io/sitebuilder/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db.InitDB("atomic-commits", "sitebuilder")

	controllers.Routes(router)

	router.Run(":3000")

	db.Client.Close()
}
