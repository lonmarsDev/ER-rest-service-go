package models


type UserScore struct {
	UserID uint8 `json:"userId"`
	Score uint8 `json:"score"`
}

type AllScore struct {
	Managers []UserScore `json:"managers"`
	Team []UserScore `json:"team"`
	Others []UserScore `json:"others"`
}

type Scores struct {
	AllScore AllScore `json:"scores"`
}