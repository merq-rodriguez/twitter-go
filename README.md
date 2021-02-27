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

```
sonar-scanner \
  -Dsonar.projectKey=Clon-Twitter \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://localhost:9000 \
  -Dsonar.login=4579e90050f27193bb757a5c0c6c55bc42cb6f93
```



docker run \
    --rm \
    -e SONAR_HOST_URL="http://127.0.0.1:9000" \
    -e SONAR_LOGIN="4579e90050f27193bb757a5c0c6c55bc42cb6f93" \
    -v "$HOME/Documentos/Projects/GoProjects/src/github.com/merq-rodriguez/twitter-go:/usr/src" \
    sonarsource/sonar-scanner-cli