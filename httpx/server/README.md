# server
## Dependency
[mux](https://github.com/gorilla/mux)
[concurrent map](github.com/orcaman/concurrent-map)
[zap](go.uber.org/zap)

## Documentation
### Quick start
```go=
type Response struct {
	Msg string `json:msg`
}

func handleFunc(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if data, err := req.Form["name"]; err {
		jsonObj := &Response{
			Msg: "Hello " + data[0],
		}
		resp, err := json.Marshal(jsonObj)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
		} else {
			w.Header().Set("Content-Type", "application/json;charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			w.Write(resp)
		}
	} else {
		jsonObj := &Response{
			Msg: "Lack of necessary variables name",
		}
		resp, err := json.Marshal(jsonObj)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
		} else {
			w.Header().Set("Content-Type", "application/json;charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			w.Write(resp)
		}
	}
}

func main() {
    // register os signal
    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

    server := server.New(8080)
    server.RegisterAPI("test", "POST", "/api/test", handleFunc)
    server.Start()

    for {
        select {
        case sig := <-interrupt:

            fmt.Println(sig)

            // shutdown HTTP server
            server.Stop()
            server = nil

            os.Exit(0)
        }
    }
}
```

### Function
#### New
Use this function to get a server
```go=
type Server struct {
	httpServer *http.Server
	exitDone   *sync.WaitGroup
	Port       int
	Middleware []mux.MiddlewareFunc
	Routes     cmap.ConcurrentMap
}

func New(port int) *Server
```
| Parameter |      Description      |
| --------- |:---------------------:|
| port      | port number of server |

| Response |    Description     |
| -------- |:------------------:|
| server   | instance of server |

##### Usage
```go=
server := server.New(8080)
```
---
#### Get
Use this function to get a server instance after New()
```go=
func Get() *Server
```
| Response |    Description     |
| -------- |:------------------:|
| server   | instance of server |

##### Usage
```go=
server := server.get()
```
---
#### RegisterAPI
Use this function to Register API for server
```go=
func (server *Server) RegisterAPI(id string, method string, pattern string, handler http.HandlerFunc, roles ...model.IRole)
```

| Parameter |                Description                |
| --------- |:-----------------------------------------:|
| id        |                 id of API                 |
| method    |                API method                 |
| pattern   |                URL of API                 |
| handler   |             handler function              |
| roles     | Roles which allow to access Api if needed |

| Response |    Description     |
| -------- |:------------------:|
| server   | instance of server |
##### Usage
```go=
type Response struct {
	Msg string `json:msg`
}

func handleFunc(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if data, err := req.Form["name"]; err {
		jsonObj := &Response{
			Msg: "Hello " + data[0],
		}
		resp, err := json.Marshal(jsonObj)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
		} else {
			w.Header().Set("Content-Type", "application/json;charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			w.Write(resp)
		}
	} else {
		jsonObj := &Response{
			Msg: "Lack of necessary variables name",
		}
		resp, err := json.Marshal(jsonObj)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
		} else {
			w.Header().Set("Content-Type", "application/json;charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			w.Write(resp)
		}
	}
}


server := server.New(8080)
server.RegisterAPI("test", "POST", "/api/test", handleFunc)
```
---
#### RegisterMiddleware
Use this function to Register Middleware for server
```go=
func (server *Server) RegisterMiddleware(mwf mux.MiddlewareFunc)
```
| Parameter |        Description         |
| --------- |:--------------------------:|
| mwf       | middleware function of mux |
##### Usage
```go=
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Do stuff here
        log.Println(r.RequestURI)
        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)
    })
}


server := server.New(8080)
server.RegisterMiddleware(loggingMiddleware)
```
---
#### Start
Use this function to start server
```go=
func (server *Server) Start()
```
##### Usage
```go=
server := server.New(8080)
server.Start()
```
---
#### Close
Use this function to close server
```go=
func (server *Server) Close()
```
##### Usage
```go=
server := server.New(8080)
server.Start()
defer server.Close()
```