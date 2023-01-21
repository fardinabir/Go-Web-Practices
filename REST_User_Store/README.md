# REST_User_Store
This project is for implementing API based CRUD operations. It covers the basic concepts of REST API, Gorrila MUX and GORM. 

The main requests and responses sample are as like below:

### POST /users ###
```
Request:
{
    "firstName": <firstname>,
    "lastName": <lastname>,
    "password": <password>,
    "phone": <phone>
}
Response:
{
    "id": <id>
}
```
<a href="https://ibb.co/1zQGntL"><img src="https://i.ibb.co/kmB2ysx/POST-req.png" alt="POST-req" border="0"></a>

### GET /users/{id} ###
```
Response:
{
    "id": <id>,
    "name": <full name [firstname and lastname]>,
    "phone": <phone>
}
```
<a href="https://ibb.co/020wRhD"><img src="https://i.ibb.co/xLn4tJC/GET-users.png" alt="GET-users" border="0"></a>

### POST /users/{id}/tags ###
```
Request:
{
    "tags": [<tag 1>, <tag 2>, ...],
    "expiry": <miliseconds>        
}
Response:
{}
```
<a href="https://ibb.co/1Z0tLvQ"><img src="https://i.ibb.co/6Yy7mnw/POST-tags.png" alt="POST-tags" border="0"></a>

### GET /users?tags=tag1,tag2... ###
```
Response:
{
    "users": [
        {
            "id": <id 1>
            "name": "<full name 1>"
            "tags": [<tag 1>, <tag 2>, ...]
        },
    ]
}
```
<a href="https://ibb.co/P5qfkVT"><img src="https://i.ibb.co/cN47qsb/GET-tags.png" alt="GET-tags" border="0"></a>
