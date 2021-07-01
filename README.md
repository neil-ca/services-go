## This repo is a collection of services written in go, various are in experimental
## One example inside of jsonstore 

| Endpoint                 | Method  | Description |
| :----------------------: | :----: | :-----: |
| /v1/user/id              | Get  | Get a user using ID |
| /v1/user                 | Post | Create a new user |
| /v1/user?first_name=NAME | Get  | Get all users by the given first name|
| /v1/order/id             | Get  | Get an order with the given ID |
| /v1/order                | Post | Create a new order |

## Create resource
```
curl -X POST \
http://localhost:8000/v1/user \
-H 'cache-control: no-cache' \
-H 'content-type: application/json' \
-d '{
"username": "neil",
"email_address": "ulicode4@gmail.com",
"first_name": "Neil",
"last_name": "Ulises"
}'
```

## Obtain resource
```
curl -X GET http://localhost:8000/v1/user/1
```
## It returns all the details about the user
``` 
{"user":{"ID":1,"CreatedAt":"2020-09-16T13:05:37.01959-05:00",
"UpdatedAt":"2020-09-16T13:05:37.01959-05:00","DeletedAt":null,"Orders":null},
"data":{"email_address":"ulicode4@gmail.com","first_name":"Christopher Marshal","last_name":"ulises","username":"neil"}}
```
```
curl -X GET 'http://localhost:8000/v1/user?first_name=Christopher Marshal'
```

