package sorcery

type StatusMsg struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func NewStatusMsg(statusCode int, msg string) StatusMsg {
	return StatusMsg{StatusCode: statusCode, Message: msg}
}
