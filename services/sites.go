package services

import (
	"context"
	"errors"

	"atomiccommits.io/sitebuilder/db"
	"cloud.google.com/go/bigtable"
)

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

type Page struct {
	Url  string `json:"url"  form:"url"  binding:"required"`
	Html string `json:"html" form:"html" binding:"required"`
	Id   string `json:"id"   form:"id"`
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

func ReadSite(c context.Context, id string, versions int) ([]Page, error) {
	var pages []Page = make([]Page, 0)

	if versions < 0 {
		return pages, errors.New("versions must be a positive integer")
	}

	siteId, err := getIdByUUID(c, id)
	if err != nil {
		return pages, err
	}

	sites := db.Client.Open("sites")
	row, _ := sites.ReadRow(c, siteId, bigtable.RowFilter(bigtable.ColumnFilter("html")))

	if row == nil {
		return pages, errors.New("No entry for site with id " + id)
	}

	var items []bigtable.ReadItem = row["content"]

	if versions == 0 || versions == 1 {
		// TODO: Return an un-inverted url
		pages = append(
			pages,
			Page{
				Url:  siteId,
				Html: string(items[0].Value),
				Id:   id,
			})

		return pages, nil
	}

	count := min(versions, len(items))
	for index := 0; index < count; index++ {
		pages = append(
			pages,
			Page{
				Url:  siteId,
				Html: string(items[index].Value),
				Id:   id,
			})
	}

	return pages, nil
}

func ReadSites(c context.Context, prefix string) ([]Page, error) {
	sites := db.Client.Open("sites")

	var pages []Page = make([]Page, 0)

	rr := bigtable.PrefixRange(prefix)
	sites.ReadRows(c, rr, func(r bigtable.Row) bool {
		pages = append(
			pages,
			Page{
				Url:  r.Key(),
				Html: string(r["content"][0].Value),
				Id:   string(r["meta"][0].Value),
			})
		return true
	})

	return pages, nil
}
