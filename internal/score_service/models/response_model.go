package models

type RespScore struct {
	Manager float64 `json:"managers,omitempty"`
	Team float64 `json:"team,omitempty"`
	Others float64 `json:"others,omitempty"`
}

type Resp struct {
	Success bool `json:"success"`
	Data RespScore `json:"data"`
	Errors []string `json:"errors"`
}

