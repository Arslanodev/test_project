### Endpoints
**Blog operations**  
`GET`: `api/v1/blog` - returns all blogs  

`GET`: `api/v1/blog/{id}` - returns blog at id  

`POST`: `api/v1/blog` - creates a new blog. (Admin restricted). Request body:  
```json
{
    "title" : "This is a Title",
    "text": "This a new text"
}
```
`DELETE`: `api/v1/blog/{id}` - deletes blog at id. (Admin restricted)

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
All requests to the `api/v1/blog` endpoint are available only to authorized users. To be authorized, user should request a token by logging in with *username* and *password*.  

Each request to the `api/v1/blog` should contain `"Authorization"` header with token as value.

### dependencies:
- Gorm
- Gin
- bcrypt
- jwt-go
- GoDotEnv

