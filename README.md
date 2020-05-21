# Sitebuilder

An API written in Go, built on top of Google BigTable.

It functions for a site editor.

It creates and updates versions of sites, which are just strings — presumed to be valid HTML.

## Setting up BigTable

1. Create a BigTable Instance.
2. Create tables, column families, and seed the database.
3. Create a service account with the proper permissions and set the environment variable `GOOGLE_APPLICATION_CREDENTIALS`.

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

Here are the `cbt` commands that I used to create the tables, column families, and put some seed data in the db.

```
cbt createtable sites
cbt createfamily sites content
cbt createfamily sites meta

cbt createtable site-ids
cbt createfamily site-ids a


cbt set sites io.atomiccommits/welcome meta:id="671221eb-654a-434b-8363-b9bead78c68b" content:html="<html>\n    <body>\n        <h1>Welcome</h1>\n    </body>\n</html>\n\n"
cbt set sites io.atomiccommits/welcome meta:id="671221eb-654a-434b-8363-b9bead78c68b" content:html="<html>\n    <body>\n        <h1>Bienvenidos</h1>\n    </body>\n</html>\n\n"
cbt set sites io.atomiccommits/welcome meta:id="671221eb-654a-434b-8363-b9bead78c68b" content:html="<html>\n    <body>\n        <h1>Bem-vindo</h1>\n    </body>\n</html>\n\n"

cbt set site-ids 671221eb-654a-434b-8363-b9bead78c68b a:a="io.atomiccommits/welcome"
```
