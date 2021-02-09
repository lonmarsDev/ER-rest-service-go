package handler

import (
	"encoding/json"
	"fmt"
	"github.com/lonmarsDev/ER-rest-service-go/internal/score_service/models"
	"net/http"
)

type ResponseData struct {
	Content models.Resp
}

type ResponseHandlerAction interface {
	SetSuccess(s bool) *ResponseHandler
	GetStatus() bool
	SetWriter(w http.ResponseWriter )
	SetData(managers, team, other float64) *ResponseHandler
	SetError(err string) *ResponseHandler
	SetErrors(err []string) *ResponseHandler
	SetRespCode(c int) *ResponseHandler
	Render()
}

type ResponseHandler struct {
	Data ResponseData
	RespCode int
	Writer http.ResponseWriter
	ResponseHandlerAction
}

func NewResponseHandler() *ResponseHandler{
	resp := &ResponseHandler{}
	resp.RespCode = http.StatusOK
	resp.Data.Content.Success = false
	resp.Data.Content.Errors = []string{}
	resp.Data.Content.Data = models.RespScore{}
	return resp
}

func (r ResponseHandler) SetSuccess(s bool) *ResponseHandler{
	r.Data.Content.Success = s
	return &r
}
func (r ResponseHandler) GetStatus() bool{
	return r.Data.Content.Success
}

func (r ResponseHandler) SetWriter(w http.ResponseWriter ){
	r.Writer = w
}

func (r ResponseHandler) SetError(err string) *ResponseHandler{
	r.Data.Content.Errors = append(r.Data.Content.Errors, err)
	return &r
}

func (r ResponseHandler) SetErrors(err []string) *ResponseHandler{
	r.Data.Content.Errors = err
	return &r
}

func (r ResponseHandler) SetData(managers, team, other float64) *ResponseHandler{
	r.Data.Content.Data.Manager = managers
	r.Data.Content.Data.Team = team
	r.Data.Content.Data.Others = other
	return &r
}

func (r ResponseHandler) SetRespCode(c int) *ResponseHandler{
	r.RespCode = c
	return &r
}

func (r ResponseHandler) Render(){
	r.Writer.Header().Set("Content-Type", "application/json")
	r.Writer.WriteHeader(r.RespCode)

	output, err := json.Marshal(r.Data.Content)
	if err != nil {
		http.Error(r.Writer, err.Error(), 500)
		return
	}
	fmt.Fprint(r.Writer, string(output))
}