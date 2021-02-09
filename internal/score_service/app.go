package score_service

import (
	"github.com/lonmarsDev/ER-rest-service-go/internal/score_service/handler"
	"github.com/lonmarsDev/ER-rest-service-go/internal/score_service/routes"
	"github.com/lonmarsDev/ER-rest-service-go/pkg/log"
	"github.com/lonmarsDev/ER-rest-service-go/pkg/server"
)

const Tag = "score_service"
func Init(){
	//Call server init
	log.Info(Tag, "starting the service")
	srv := server.NewRestAPI()
	//Set rest tag, for logging tag
	srv.SetTag(Tag)
	//Enable logger
	srv.EnableLogger()
	//Call the handler
	handler := handler.Init()
	route := routes.Init()
	route.SetHandler(handler).SetServer(&srv).Build()

	if err := srv.Serve(); err!= nil {
		log.Error(Tag, err.Error())
	}

}
