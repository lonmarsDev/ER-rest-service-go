package handler

import (
	"encoding/json"
	"github.com/lonmarsDev/ER-rest-service-go/internal/score_service/data_pipe"
	"github.com/lonmarsDev/ER-rest-service-go/internal/score_service/models"
	"net/http"
)

type HandlerAction interface {
	CalculateScore(w http.ResponseWriter, r *http.Request)
}

// Handler handler struct
type Handler struct {
	HandlerAction
}

// Init is the handler constructor
func Init() *Handler {
	return &Handler{}
}


// CalculateScore is to Calculate score
func(h Handler) CalculateScore(w http.ResponseWriter, r *http.Request){

	var scores models.Scores
	respHandler := NewResponseHandler()
	respHandler.Writer = w
	err := json.NewDecoder(r.Body).Decode(&scores)
	if err != nil {
		respHandler.RespCode = http.StatusBadRequest
		respHandler.Data.Content.Errors = []string{err.Error()}
		respHandler.Render()
		return
	}
	dp  := data_pipe.Init().SetData(scores)
	errs := dp.Validate()
	if len(errs) > 0 {
		respHandler.RespCode = http.StatusBadRequest
		respHandler.SetSuccess(false)
		respHandler.Data.Content.Data = models.RespScore{}
		respHandler.Data.Content.Errors = errs


	}else {
		dp.Analize()
		scores := dp.GetScores()
		respHandler.SetData(scores.Manager, scores.Team, scores.Other)
		respHandler.SetRespCode(http.StatusOK)
		respHandler.SetSuccess(true)

	}
	respHandler.Render()
}


