# Sitebuilder

An API written in Go, built on top of Google BigTable.

It functions for a site editor.

It creates and updates versions of sites, which are just strings — presumed to be valid HTML.

## To Run

```
git clone https://github.com/alexhwoods/atomiccommits.io-sitebuilder.git
go run .
```

## Format of Database

This tutorial uses a BigTable database, with two tables:

```
sites
  (key) <inverted-url>:
    content:
      html
    editor:
      user

site-ids
  (key) <uuid>:
    a:
      a:
        <inverted-url>  // key of sites
```
