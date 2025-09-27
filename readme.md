#go-api-template

##Microservice

Simple microservice created to be a sample.
Now its using Mongo as database. (it can be changed)

##How to run?

Create an .env file with MONGO_URI and PORT (default 8080)
```
make run
```

##CURL Examples


```
//Create
curl -X POST -d '{"ID": "1", "Title": "New task", "Done": false}' -H "Content-Type: application/json" localhost:8080/tasks

//Get
curl localthost:8080/tasks

//Delete
curl -X DELETE "localhost:8080/tasks?id=1"  
```