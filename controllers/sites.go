package controllers

import (
	"io/ioutil"

	"atomiccommits.io/sitebuilder/internal"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	sites := router.Group("/sites")
	{
		sites.POST("", CreateSite)
	}
}

func CreateSite(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	html := string(body)

	c.String(200, internal.CreateSite(html))
}
