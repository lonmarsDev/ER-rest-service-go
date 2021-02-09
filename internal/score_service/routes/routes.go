package routes

import (
	"github.com/lonmarsDev/ER-rest-service-go/internal/score_service/handler"
	"github.com/lonmarsDev/ER-rest-service-go/pkg/server"
)

type Route struct {
	server server.RestSvc
	handler handler.Handler
}

const rootRoute = "/"

func Init() *Route{
	return &Route{}
}

func (r Route) SetServer(s *server.RestSvc) *Route{
	if s != nil{
		r.server = *s
	}
	return &r
}

func (r Route) SetHandler(h *handler.Handler) *Route{
	if h != nil{
		r.handler = *h
	}
	return &r
}

func (r Route)Build()*Route{
	r.server.Rest.Post(rootRoute, r.handler.CalculateScore)
	return &r
}