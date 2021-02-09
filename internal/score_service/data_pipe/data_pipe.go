package data_pipe

import (
	"fmt"
	"github.com/lonmarsDev/ER-rest-service-go/internal/score_service/models"
)

type Score struct {
	Manager, Team, Other float64
}

type DAtaPipeAction interface {
	SetData(pd models.Scores) *DataPipe
	GetSuccess() bool
	GetScores() *Score
	computeManagersScore() float64
	computeTeamsScore() float64
	computeOthersScore() float64
	Validate() []string
	Analize()
	SetError(format string, a ...interface{}) *DataPipe
}

type DataPipe struct {
	Data models.Scores
	Error []string
	Score Score
	IsSuccess bool
	DAtaPipeAction
}

func Init()*DataPipe{
	return &DataPipe{
		Data: models.Scores{
			AllScore: models.AllScore{
				Managers:[]models.UserScore{},
				Team: []models.UserScore{},
				Others: []models.UserScore{},
			},
		},
		Error: []string{},
		Score: Score{Manager: 0, Other: 0, Team: 0},
		IsSuccess: false,
	}
}

func (dp DataPipe)SetData(pd models.Scores)*DataPipe{
	dp.Data = pd
	return &dp
}

func (dp DataPipe)setSuccess(s bool)*DataPipe{
	dp.IsSuccess = s
	return &dp
}

func (dp DataPipe)GetSuccess()bool{
	return dp.IsSuccess
}

func (dp DataPipe)GetScores()*Score{
	return &dp.Score
}

func (dp DataPipe)Analize() *DataPipe{

	dp.Score.Manager = dp.computeManagersScore()
	dp.Score.Team = dp.computeTeamsScore()
	dp.Score.Other = dp.computeOthersScore()

	if len(dp.Error)  > 0{
	 	dp.setSuccess(false)
	}else{
		dp.setSuccess(true)
	}
	return &dp
}

func (dp DataPipe)computeManagersScore()float64 {
	for _, manager := range dp.Data.AllScore.Managers {
		dp.Score.Manager  += float64(manager.Score)
	}
	return dp.Score.Manager/ float64(len(dp.Data.AllScore.Managers))
}

func (dp DataPipe)computeTeamsScore() float64{
	for _, team := range dp.Data.AllScore.Team {
		dp.Score.Team  += float64(team.Score)
	}
	return dp.Score.Team/ float64(len(dp.Data.AllScore.Team))

}

func (dp DataPipe)computeOthersScore()float64 {
	for _, other := range dp.Data.AllScore.Others {
		dp.Score.Other  += float64(other.Score)
	}
	return dp.Score.Other / float64(len(dp.Data.AllScore.Others))

}


func (dp DataPipe)Validate() []string {
	if len(dp.Data.AllScore.Managers) == 0 {
		dp.Error = append(dp.Error, "Score for managers is empty")
	}
	if len(dp.Data.AllScore.Team) == 0 {

		dp.Error = append(dp.Error, "Score for teams is empty")
	}
	if len(dp.Data.AllScore.Others) == 0 {
		dp.Error = append(dp.Error, "Score for other is empty")
	}
	users := make(map[uint8]uint8)

	for _, manager := range dp.Data.AllScore.Managers {
		if users[manager.UserID] == manager.UserID {
			dp.Error = append(dp.Error, fmt.Sprintf("duplicate user id %d", manager.UserID))
		}
		if manager.Score > 5 {
			dp.Error = append(dp.Error, fmt.Sprintf("invalid score on user id %d", manager.UserID))
		}
		users[manager.UserID] = manager.UserID
	}
	for _, team := range dp.Data.AllScore.Team {
		if users[team.UserID] == team.UserID {
			dp.Error = append(dp.Error, fmt.Sprintf("duplicate user id %d", team.UserID))
		}
		if team.Score >5 {
			dp.Error = append(dp.Error, fmt.Sprintf("invalid score on user id %d", team.UserID))
		}
		users[team.UserID] = team.UserID
	}
	for _, other := range dp.Data.AllScore.Others {
		if users[other.UserID] == other.UserID {
			dp.Error = append(dp.Error, fmt.Sprintf("duplicate user id %d", other.UserID))
		}
		if other.Score >5 {
			dp.Error = append(dp.Error, fmt.Sprintf("invalid score on user id %d", other.UserID))
		}
		users[other.UserID] = other.UserID
	}
	return dp.Error
}

func (dp DataPipe)SetError(format string, a ...interface{}) *DataPipe{
	err:= fmt.Sprintf(format, a)
	dp.Error = append(dp.Error, err)
	return &dp
}
