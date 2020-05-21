package controllers

import (
	"strconv"

	"atomiccommits.io/sitebuilder/services"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	sites := router.Group("/sites")
	{
		sites.POST("", CreateSite)
		sites.PUT("/:id", UpdateSite)
		sites.GET("/:id", GetSite)
		sites.GET("", GetSites)
	}
}

func CreateSite(c *gin.Context) {
	body := new(services.Page)
	err := c.Bind(body)

	if err != nil {
		c.String(400, "Request body not formed correctly.")
	}

	responseBody, e := services.CreateSite(c.Request.Context(), body)
	if e != nil {
		c.PureJSON(400, gin.H{
			"error": e.Error(),
		})
	} else {
		c.PureJSON(200, responseBody)
	}
}

func UpdateSite(c *gin.Context) {
	id := c.Param("id")
	body := new(services.UpdatePage)
	err := c.Bind(body)

	if err != nil {
		c.String(400, "Request body not formed correctly.")
	}

	responseBody, e := services.UpdateSite(c.Request.Context(), id, body)
	if e != nil {
		c.PureJSON(400, gin.H{
			"error": e.Error(),
		})
	} else {
		c.PureJSON(200, responseBody)
	}
}

func GetSite(c *gin.Context) {
	id := c.Param("id")

	versions, _ := strconv.Atoi(c.Query("versions"))
	values, err := services.ReadSite(c.Request.Context(), id, versions)

	if err != nil {
		c.PureJSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		c.PureJSON(200, gin.H{
			"data": values,
		})
	}
}

func GetSites(c *gin.Context) {
	prefix := c.Query("prefix")
	values, _ := services.ReadSites(c.Request.Context(), prefix)

	c.PureJSON(200, gin.H{
		"data": values,
	})
}
