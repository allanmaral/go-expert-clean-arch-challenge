package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      []Handler
	WebServerPort string
}

type HttpMethod string

const (
	MethodGet  = http.MethodGet
	MethodPost = http.MethodPost
)

type Handler struct {
	Path   string
	Method HttpMethod
	Func   http.HandlerFunc
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make([]Handler, 0),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method HttpMethod, path string, handler http.HandlerFunc) {
	s.Handlers = append(s.Handlers, Handler{
		Path:   path,
		Method: method,
		Func:   handler,
	})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, handler := range s.Handlers {
		switch handler.Method {
		case MethodGet:
			s.Router.Get(handler.Path, handler.Func)
		case MethodPost:
			s.Router.Post(handler.Path, handler.Func)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
