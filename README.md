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
    meta:
      id

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

# Demo

### Create a Site

```curl
curl --location --request POST 'http://localhost:3000/sites' \
--header 'Content-Type: application/json' \
--data-raw '{
    "url": "https://atomiccommits.io/foo",
    "html": "<html>\n    <body>\n        <h1>Bem-vindo</h1>\n    </body>\n</html>\n"
}
'
```

Response: 200

```
{
    "url": "atomiccommits.io/foo",
    "html": "<html>\n    <body>\n        <h1>Bem-vindo</h1>\n    </body>\n</html>\n",
    "id": "c7f18a41-1c2d-44da-9f99-879de5feb5c9"
}
```

### Update a Site

```
curl --location --request PUT 'http://localhost:3000/sites/c7f18a41-1c2d-44da-9f99-879de5feb5c9' \
--header 'Content-Type: application/json' \
--data-raw '{
    "html": "<html>\n    <body>\n        <h1>Welcome 🎉</h1>\n    </body>\n</html>\n"
}'
```

Response: 200

```
{
    "url": "atomiccommits.io/foo",
    "html": "<html>\n    <body>\n        <h1>Welcome 🎉</h1>\n    </body>\n</html>\n",
    "id": "c7f18a41-1c2d-44da-9f99-879de5feb5c9"
}
```

I'm going to run this about 10 more times with different emojis, to illustrate the version history feature of this API.

### Get a Site

```
curl --location --request GET 'http://localhost:3000/sites/c7f18a41-1c2d-44da-9f99-879de5feb5c9?versions=100'
```

(As you can see, `versions` is a max of how many versions to get, with a default of 1.)

Response: 200

```
{
    "data": [
        {
            "url": "atomiccommits.io/foo",
            "html": "<html>\n    <body>\n        <h1>Welcome 👟 ️</h1>\n    </body>\n</html>\n",
            "id": "c7f18a41-1c2d-44da-9f99-879de5feb5c9"
        },
        {
            "url": "atomiccommits.io/foo",
            "html": "<html>\n    <body>\n        <h1>Welcome 🎃 ️</h1>\n    </body>\n</html>\n",
            "id": "c7f18a41-1c2d-44da-9f99-879de5feb5c9"
        },
        {
            "url": "atomiccommits.io/foo",
            "html": "<html>\n    <body>\n        <h1>Welcome 🕵️</h1>\n    </body>\n</html>\n",
            "id": "c7f18a41-1c2d-44da-9f99-879de5feb5c9"
        },
        {
            "url": "atomiccommits.io/foo",
            "html": "<html>\n    <body>\n        <h1>Welcome 👀</h1>\n    </body>\n</html>\n",
            "id": "c7f18a41-1c2d-44da-9f99-879de5feb5c9"
        },
        {
            "url": "atomiccommits.io/foo",
            "html": "<html>\n    <body>\n        <h1>Welcome 🤠</h1>\n    </body>\n</html>\n",
            "id": "c7f18a41-1c2d-44da-9f99-879de5feb5c9"
        },
        {
            "url": "atomiccommits.io/foo",
            "html": "<html>\n    <body>\n        <h1>Welcome 😎</h1>\n    </body>\n</html>\n",
            "id": "c7f18a41-1c2d-44da-9f99-879de5feb5c9"
        },
        {
            "url": "atomiccommits.io/foo",
            "html": "<html>\n    <body>\n        <h1>Welcome 😂</h1>\n    </body>\n</html>\n",
            "id": "c7f18a41-1c2d-44da-9f99-879de5feb5c9"
        },
        {
            "url": "atomiccommits.io/foo",
            "html": "<html>\n    <body>\n        <h1>Welcome 🎉</h1>\n    </body>\n</html>\n",
            "id": "c7f18a41-1c2d-44da-9f99-879de5feb5c9"
        },
        {
            "url": "atomiccommits.io/foo",
            "html": "<html>\n    <body>\n        <h1>Bem-vindo</h1>\n    </body>\n</html>\n",
            "id": "c7f18a41-1c2d-44da-9f99-879de5feb5c9"
        }
    ]
}
```

### Get Many Sites

The call to get all the sites is simple.

```
curl http://localhost:3000/sites
```

It has 1 cool feature, in that you can filter the sites you get back using a prefix, which corresponds inverted url syntax.

```
curl http://localhost:3000/sites?prefix=io.atomiccommits/foo
```

Response: 200

```
{
    "data": [
        {
            "url": "atomiccommits.io/foo",
            "html": "<html>\n    <body>\n        <h1>Welcome 👟 ️</h1>\n    </body>\n</html>\n",
            "id": "c7f18a41-1c2d-44da-9f99-879de5feb5c9"
        },
        {
            "url": "atomiccommits.io/foo/bar",
            "html": "<html>\\n    <body>\\n        <h1>Hello</h1>\\n    </body>\\n</html>\\n\\n",
            "id": "49ab0085-a0fb-427a-8ae8-457935276482"
        }
    ]
}
```
