# EndPoint API (Create Read)
### I'm Using Golang Language with Redis Cache

Golang Documentation : https://go.dev/doc/ <br/>
Golang Instalation : https://go.dev/doc/install <br/>
Redis Instalation : https://redis.io/download
<br/>

### How To Run The Code
> go run main.go

## EndPoint Documentation ( Using Postman)
### Get All Articles
1. standart response <br/>
> http://localhost:8000/articles

2. pagination response:
Example to Add in Params: <br/>
key: page, value: 1 <br/>
key: per_page , value:3 <br/>
key: paginate , value:1 (true or false for using pagination response) <br/>
> http://localhost:8000/articles?page=1&per_page=3&paginate=1

3. sort by {collumn} ( default sort by last created_at )
Example to Add in Params: <br/>
key: sort , value author|asc -> for asc <br/>
key: sort , value author|desc -> for desc <br/>
> http://localhost:8000/articles?sort=author|asc

4. search any word in title or body
Example to Add in Params: <br/>
key: search , value:Golang Coding <br/>
> http://localhost:8000/articles?search=Golang Coding

5. filter by {column} (can add more than 1 filter)
Example to Add in Params: <br/>
key:filter[] , value:{"option":"author","operator":"=","value":"jasper"} <br/>
key:filter[] , value:{"option":"title","operator":"=","value":"Golang Coding"} <br/>
> http://localhost:8000/articles?filter[]={"option":"author","operator":"=","value":"jasper"}&filter[]={"option":"title","operator":"=","value":"Golang Coding"}

### Get Single Article
1. Run redis server first 
> redis-server
2. Check redis server
> redis-cli ping

> http://localhost:8000/articles/{article_id} 

### Post Single Article
> http://localhost:8000/articles 
<br/>
body:
{
    "author": "Jasper",
    "title": "Golang Code",
    "body": "This is a Body of Golang Code Article"
}



