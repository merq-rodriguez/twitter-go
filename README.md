# Clone Twitter Backend Go + MongoDB + JWT


# Run
```bash
go run main.go
```


### Signup

> POST http://host:port/signup

```json
{
    "name":"John",
    "lastname":"Snow",
    "birthDate":"1994-06-30T00:00:00Z",
    "email":"snow@gmail.com",
    "password":"12345678",
    "avatar":"/",
    "banner":"/",
    "location": "Lugar",
    "website": "www.website.com"
}
```