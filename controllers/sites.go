package controllers

import (
	"atomiccommits.io/sitebuilder/internal"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	sites := router.Group("/sites")
	{
		sites.GET("", internal.GetAllSites)
	}
}
