package controllers

import (
	"io/ioutil"

	"atomiccommits.io/sitebuilder/services"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	sites := router.Group("/sites")
	{
		sites.POST("", CreateSite)
		sites.POST("/:id", UpdateSite)
		sites.GET("/:id", GetSite)
	}
}

func CreateSite(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	html := string(body)

	c.String(200, services.CreateSite(html))
}

func UpdateSite(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	html := string(body)

	c.String(200, services.CreateSite(html))
}

func GetSite(c *gin.Context) {
	id := c.Param("id")

	c.String(200, "getting site "+id)
}
