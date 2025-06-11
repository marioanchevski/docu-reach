# 📄 Docu-Reach

**docu-reach** is a simple REST API written in **Golang**, designed to manage and search "documents". It provides functionality for creating, retrieving, deleting, and performing advanced text search over documents.

---

## 🚀 Features

- ✅ Create new documents
- 🔍 Retrieve documents by ID or get all documents
- ❌ Delete documents
- 🔎 Perform advanced **text search** in both **title** and **description** fields
  - Supports **multiple search terms**
  - Logical operators: **AND**, **OR**, **NOT** (via intuitive syntax)
- 💾 In-memory storage (no database required)

---


## Getting started

To get started clone this repo, after that you have a few options to run the applicaiton.
1. To use this app you need Go verison 1.24.3 If you fulfill this requirement you can run the app with
```
go build -o bin/docu-reach cmd/main.go && ./bin/docu-reach
```

2. You can also use the make build tool
```
make run
```

3. If you dont have go and dont want to install it, you can use docker
```bash
docker build -t docu-reach .
docker run -p 8080:8080 docu-reach
```
---

## 📦 Document Structure

Each document has the following structure:

```json
{
  "id": 1,
  "title": "My Title",
  "description": "A longer description"
}
```

## API usage

The app provides the following APIs
```
POST	/api/v1/documents	Create a new document
GET	/api/v1/documents	List all documents
GET	/api/v1/documents/{id}	Get document by ID
DELETE	/api/v1/documents/{id}	Delete document by ID
GET	/api/v1/documents/search?...	Search documents by fields
```

### Search functionality

While most of the API are intuitive and easy to use, the search functionality requires some aditional explanation.

The search functionality checks if the terms provided by the user are contained in the document. For this the engine uses loose equality.

example document dabatase 
```json
[
    {
        "id": 1,
        "title": "title1",
        "description": "pot"
    },
    {
        "id": 2,
        "title": "title2",
        "description": "desc containing potato"
    },
    {
        "id": 3,
        "title": "title3",
        "description": "something else"
    },
]
```
Searching for the term `pot` in the description will return documents that contain `pot` but also `potato` in their titles. (documents with id 1 and 2)


In order to use the search functionality, the user has to provide one or more things to search by. The possible query parameters are
- title
- description
- op

#### providing search terms for title, descripton
When providing a query parameter it has to follow some rules:
The user provides search terms separated by a comma character `,`
Prefixing the term with `-` means exclude the term, otherwise include

example
```
/api/v1/documents/search?title=bar,-foo,baz
```
in this example we want to find all titles that contain `bar`, and contain `baz` and dont contain `foo`

example
```
/api/v1/documents/search?title=bar,-foo,baz&description=-potato
```
in this example we want to find all titles that contain `bar`, and contain `baz` and dont contain `foo` and dont contain `potato` in the description

`In order for the search to work title or description need to be provided.`

#### the op query parameter
The op query parameter is optional, it is used when providing both `title` and `description`.
The possible values for op are `and` `or`
If you provide title and description but not op, or provide an unsported value it defaults to `and`


example
```
/api/v1/documents/search?title=bar,-foo,baz&description=-potato
```
op defaults to `and`

example
```
/api/v1/documents/search?title=bar,-foo,baz&description=-potato&op=asjkdfhas
```
op defaults to `and`

The op param is used in order to determine the relationship between `title` and `description`
If the value is `and` it has to satisfy all the terms in title `AND` all the terms in description, but with a value of `or`, it tells the engine to satisfy the search terms in title `OR` those in the description


example
```
/api/v1/documents/search?title=bar,-foo,baz&description=-potato&op=or
```
in this example we want to find all titles that contain `bar`, and contain `baz` and dont contain `foo` `OR` dont contain `potato` in the description
