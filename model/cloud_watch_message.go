package model

type CloudWatchMessage struct {
	Message   string  `json:"message"`
	Context   context `json:"context"`
	Level     int     `json:"level"`
	LevelName string  `json:"level_name"`
	Channel   string  `json:"channel"`
	Datetime  string  `json:"datetime"`
}

type context struct {
	UserId    int `json:"userId"`
	Exception exception
}

type exception struct {
	Class   string `json:"class"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	File    string `json:"file"`
}
