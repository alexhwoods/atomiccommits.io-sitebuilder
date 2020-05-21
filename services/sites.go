package services

import (
	"context"
	"fmt"

	"atomiccommits.io/sitebuilder/db"
	"cloud.google.com/go/bigtable"
)

func CreateSite(c context.Context, html string) string {
	sites := db.Client.Open("sites")
	row, _ := sites.ReadRow(c, "io.atomiccommits/welcome", bigtable.RowFilter(bigtable.ColumnFilter("html")))
	return string(row["content"][0].Value)
}

func ReadSite(c context.Context) string {
	sites := db.Client.Open("sites")
	fmt.Println("18")
	row, _ := sites.ReadRow(c, "io.atomiccommits/welcome", bigtable.RowFilter(bigtable.ColumnFilter("html")))
	fmt.Println("20")
	return string(row["content"][0].Value)
}
