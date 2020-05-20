package main

import (
	"context"
	"fmt"

	"atomiccommits.io/sitebuilder/controllers"
	"atomiccommits.io/sitebuilder/db"
	"cloud.google.com/go/bigtable"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db.InitDB("atomic-commits", "sitebuilder")

	sites := db.Client.Open("sites")
	row, _ := sites.ReadRow(context.Background(), "io.atomiccommits/welcome", bigtable.RowFilter(bigtable.ColumnFilter("html")))
	fmt.Println(string(row["content"][0].Value))

	controllers.Routes(router)

	router.Run(":3000")
}
