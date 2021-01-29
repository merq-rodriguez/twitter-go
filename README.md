# Clone Twitter Backend Go + MongoDB + JWT

# Install Gin
```
 go get github.com/codegangsta/gin
```

# Run live reload
```bash
gin --port 4000 run main.go
```
ó
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