package model

type Slack struct {
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Content  string `json:"text"`
}
