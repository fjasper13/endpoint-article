# EndPoint API (Create Read)
### I'm Using Golang Language with Redis Cache

Golang Documentation : https://go.dev/doc/
Golang Instalation : https://go.dev/doc/install
Redis Instalation : https://redis.io/download
<br/>

### How To Run The Code
> go run main.go

## EndPoint Documentation ( Using Postman)
### Get All Articles
1. standart response
http://localhost:8000/articles -> standard response

2. pagination response:
Example to Add in Params:
key: page, value: 1
key: per_page , value:3
key: paginate , value:1 (true or false for using pagination response)

3. sort by {collumn} ( default sort by last created_at )
Example to Add in Params:
key: sort , value author|asc -> for asc
key: sort , value author|desc -> for desc

4. search any word in title or body
Example to Add in Params:
key: search , value:Golang Coding

5. filter by {column} (can add more than 1 filter)
Example to Add in Params:
key:filter[] , value:{"option":"author","operator":"=","value":"jasper"}
key:filter[] , value:{"option":"title","operator":"=","value":"Golang Coding"}

### Get Single Article
http://localhost:8000/articles/{article_id}

### Post Single Article
http://localhost:8000/articles
body:
{
    "author": "Jasper",
    "title": "Golang Code",
    "body": "This is a Body of Golang Code Article"
}



