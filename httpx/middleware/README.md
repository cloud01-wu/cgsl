# middleware
##  Dependency
[mux](https://github.com/gorilla/mux)
[context](github.com/gorilla/context)
[jwt-go](github.com/dgrijalva/jwt-go)
[server](http://gitlab.ninthbase.com/commons/cgsl/-/blob/main/httpx/server/README.md)
[redisx](http://gitlab.ninthbase.com/commons/cgsl/-/tree/main/redisx)
## Documentation
We provide two default middleware to use if needed
### CorsMiddleware
If need CORS you can register this middleware
#### Usage
```go=
server := server.New(8080)
server.RegisterMiddleware(middleware.CorsMiddleware)
```
---
### RoleMiddleware
If need role check you can register this middleware
#### *You Must init a redis using [redisx](http://gitlab.ninthbase.com/commons/cgsl/-/tree/main/redisx)
#### Usage
```go=
server := server.New(8080)
server.RegisterMiddleware(middleware.RoleMiddleware)
```
---
### TenancyMiddleware
If need tenancy you can register this middleware
#### *You Must init a redis using [redisx](http://gitlab.ninthbase.com/commons/cgsl/-/tree/main/redisx)
#### Usage
```go=
server := server.New(8080)
server.RegisterMiddleware(middleware.TenancyMiddleware)
```