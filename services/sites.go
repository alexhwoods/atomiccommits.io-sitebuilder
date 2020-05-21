package services

import (
	"context"
	"errors"

	"atomiccommits.io/sitebuilder/db"
	"cloud.google.com/go/bigtable"
)

func wrapInArray(str string) []string {
	return []string{str}
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

func getIdByUUID(c context.Context, uuid string) (string, error) {
	siteIds := db.Client.Open("site-ids")
	row, _ := siteIds.ReadRow(c, uuid, bigtable.RowFilter(bigtable.ColumnFilter("a")))

	if row == nil {
		return "", errors.New("No site for this id")
	}

	return string(row["a"][0].Value), nil
}

func CreateSite(c context.Context, html string) string {
	sites := db.Client.Open("sites")
	row, _ := sites.ReadRow(c, "io.atomiccommits/welcome", bigtable.RowFilter(bigtable.ColumnFilter("html")))
	return string(row["content"][0].Value)
}

func ReadSite(c context.Context, id string, versions int) ([]string, error) {
	if versions < 0 {
		return wrapInArray(""), errors.New("versions must be a positive integer")
	}

	siteId, err := getIdByUUID(c, id)
	if err != nil {
		return wrapInArray(""), err
	}

	sites := db.Client.Open("sites")
	row, _ := sites.ReadRow(c, siteId, bigtable.RowFilter(bigtable.ColumnFilter("html")))

	if row == nil {
		return wrapInArray(""), errors.New("No entry for site with id " + id)
	}

	var items []bigtable.ReadItem = row["content"]

	if versions == 0 || versions == 1 {
		page := string(items[0].Value)
		return wrapInArray(page), nil
	}

	count := min(versions, len(items))
	pages := make([]string, count)
	for index := 0; index < count; index++ {
		value := string(items[index].Value)
		pages[index] = value
	}

	return pages, nil
}
