package controllers

import (
	"io/ioutil"
	"strconv"

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

	c.Request.Context()
	c.String(200, services.CreateSite(c.Request.Context(), html))
}

func UpdateSite(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	html := string(body)

	c.String(200, services.CreateSite(c.Request.Context(), html))
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
