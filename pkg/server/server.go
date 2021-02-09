package server

import (
	"fmt"
	"net/http"

	"github.com/lonmarsDev/ER-rest-service-go/pkg/server/constants"
	"github.com/lonmarsDev/ER-rest-service-go/pkg/config"
	"github.com/lonmarsDev/ER-rest-service-go/pkg/log"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	defaultPort    = "8080"
	defaultAddress = "127.0.0.1"
)

// RestConfig hold the config for the rest service
type RestConfig struct {
	Address string
	Port    string
	Tag 	string
}

// RestSvc struct hold rest svc object
type RestSvc struct {
	Rest   *chi.Mux
	Config *RestConfig
}

// NewRestAPI a constructor of rest svc, it returns RestSvc with default parameter.
func NewRestAPI() RestSvc {
	envConfig := config.GetEnvironment()
	config := new(RestConfig)
	config.Port = envConfig.GetStr(constants.Port, defaultPort)
	config.Address = envConfig.GetStr(constants.Address, defaultAddress)
	return RestSvc{
		Rest:   chi.NewMux(),
		Config: config,
	}
}

// EnableLogger a rest svc logger switch, default is false
func (rs RestSvc) EnableLogger() {
	rs.Rest.Use(middleware.Logger)
}

// SetAddress set custom address to serve rest API
func (rs RestSvc) SetAddress(addr string) RestSvc {
	rs.Config.Address = addr
	return rs
}

// SetPort set custom port to serve rest API
func (rs RestSvc) SetPort(port string) RestSvc {
	rs.Config.Port = port
	return rs
}

func (rs RestSvc) getAddress() string {
	return fmt.Sprintf("%s:%s", rs.Config.Address, rs.Config.Port)
}

func (rs RestSvc) SetTag( t string) RestSvc {
	rs.Config.Tag = t
	return rs
}

// Serve to run the rest api server
func (rs RestSvc) Serve() error {

	log.Info(rs.Config.Tag, rs.getAddress())
	return http.ListenAndServe(rs.getAddress(), rs.Rest)
}
