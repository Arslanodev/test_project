### Endpoints
To run this application simply run `make run` in the main project folder. If you want to build this project run 'make build'

**Blog operations**  
`GET`: `api/v1/post` - returns all posts  

`GET`: `api/v1/post/{id}` - returns post at id  

`POST`: `api/v1/post` - creates a new post. (Admin restricted). Request body:  
```json
{
    "title" : "This is a Title",
    "text": "This a new text"
}
```
`DELETE`: `api/v1/post/{id}` - deletes post at id. (Admin restricted)

**User operations**  
`POST`: `api/v1/user/register` - Creates a new user. Request body:  
*if user is admin add*:  `"admin": true`
```json
{
    "name": "Arslan",
    "username": "Arslan123",
    "password": "arslan123123"
}
```
`POST`: `api/v1/user/login` - Generates token for a logged in user. Request body:  
```json
{
    "username": "Arslan123",
    "password": "arslan123123"
}
```
**Authentication**  
All requests to the `api/v1/post` endpoint are available only to authorized users. To be authorized, user should request a token by logging in with *username* and *password*.  

Each request to the `api/v1/post` should contain `"Authorization"` header with bearer token as value.

### dependencies:
- Gorm
- Gin
- bcrypt
- jwt-go
- GoDotEnv

