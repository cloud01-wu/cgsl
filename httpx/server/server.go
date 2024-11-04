package server

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	cmap "github.com/orcaman/concurrent-map"

	"go.uber.org/zap"

	"github.com/cloud01-wu/cgsl/httpx/role"
	"github.com/cloud01-wu/cgsl/logger"
)

type Server struct {
	httpServer *http.Server
	exitDone   *sync.WaitGroup
	Host       string
	Port       int
	Middleware []mux.MiddlewareFunc
	Routes     cmap.ConcurrentMap
}

type RouteItem struct {
	ID      string
	Method  string
	Pattern string
	Handler http.HandlerFunc
	Roles   []role.IRole
}

type WsRouteItem struct {
	ID      string
	Pattern string
	Handler http.HandlerFunc
}

var (
	instance *Server = nil
	once     sync.Once
)

// return a instance of web server
func New(host string, port int) *Server {
	once.Do(func() {
		instance = &Server{
			Routes: cmap.New(),
			Host:   host,
			Port:   port,
		}
	})

	return instance
}

func Get() *Server {
	if instance == nil {
		return nil
	}

	return instance
}

func (server *Server) RegisterAPI(id string, method string, pattern string, handler http.HandlerFunc, roles ...role.IRole) {
	server.Routes.Set(
		id,
		RouteItem{
			ID:      id,
			Method:  method,
			Pattern: pattern,
			Handler: handler,
			Roles:   roles,
		},
	)
}

func (server *Server) RegisterWebSocket(id string, pattern string, handler http.HandlerFunc) {
	server.Routes.Set(
		id,
		WsRouteItem{
			ID:      id,
			Pattern: pattern,
			Handler: handler,
		},
	)
}

func (server *Server) RegisterMiddleware(mwf mux.MiddlewareFunc) {
	server.Middleware = append(server.Middleware, mwf)
}

func (server *Server) Start() {
	// terminal: $ go tool pprof -http=:8081 http://localhost:6060/debug/pprof/heap
	// web:
	// 1、http://localhost:8081/ui
	// 2、http://localhost:6060/debug/charts
	// 3、http://localhost:6060/debug/pprof

	server.exitDone = &sync.WaitGroup{}
	server.exitDone.Add(1)

	router := server.newRouter()
	server.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", server.Host, server.Port),
		Handler: router,
	}

	// start http server
	go func() {
		defer server.exitDone.Done()

		// always returns error
		// ErrServerClosed on graceful close
		if err := server.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			logger.New().Fatal("SERVER CLOSED ON ERROR", zap.Error(err))
		}
	}()
}

func (server *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.httpServer.Shutdown(ctx); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}

	// wait until the routine exits
	server.exitDone.Wait()
}

func (server *Server) newRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	for tuple := range server.Routes.IterBuffered() {
		if route, ok := tuple.Val.(RouteItem); ok {
			r.Name(route.ID).
				Methods(route.Method).
				Path(route.Pattern).
				Handler(route.Handler)
		} else if route, ok := tuple.Val.(WsRouteItem); ok {
			r.Name(route.ID).
				Path(route.Pattern).
				Handler(route.Handler)
		}
	}

	for _, middleware := range server.Middleware {
		if middleware != nil {
			r.Use(middleware)
		}
	}

	return r
}
