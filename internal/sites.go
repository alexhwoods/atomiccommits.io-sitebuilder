package internal

import "github.com/gin-gonic/gin"

// TODO: Define the struct of a site

func GetAllSites(c *gin.Context) {
	var sites []string = []string{"a", "b", "c", "d"}
	c.JSON(200, gin.H{"data": sites})
}
