# Restful-go

| HTTP verb |                 Path                 | Action | Resource |
| :-------- | :----------------------------------: | -----: | -------- |
| POST      |   /v1/train(details as JSON body)    | Create | Train    |
| POST      |  /v1/station(details as JSON body)   | Create | Station  |
| GET       |             /v1/train/id             |   Read | Train    |
| GET       |            /v1/station/id            |   Read | Station  |
| POST      | /v1/schedule(source and destination) | Create | Route    |


| HTTP verb |               Path               | Action | Resource |
| :-------- | :------------------------------: | -----: | -------- |
| POST      | /v1/movies(details as JSON body) | Create | Movie    |
| GET       |          /v1/movies/id           |    Get | Movie    |

## Create resource
```
curl -X POST \
http://localhost:8000/v1/movies \
-H 'cache-control: no-cache' \
-H 'content-type: application/json' \
-H 'postman-token: 6ef9507e-65b3-c3dd-4748-3a2a3e055c9c' \
-d '{ "name" : "The Dark Knight", "year" : "2008", "directors" : [
"Christopher Nolan" ], "writers" : [ "Jonathan Nolan", "Christopher Nolan"
], "boxOffice" : { "budget" : 185000000, "gross" : 533316061 }
}'
```

## Obtain resource
```
curl -X GET \
http://localhost:8000/v1/movies/5958be2a057d926f089a9700 \
-H 'cache-control: no-cache' \
-H 'postman-token: 00282916-e7f8-5977-ea34-d8f89aeb43e2'
```