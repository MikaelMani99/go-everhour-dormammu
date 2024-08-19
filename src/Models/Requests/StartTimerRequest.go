package Requests

type StartTimerRequest struct {
	Task     string `json:"task"`
	UserDate string `json:"userDate"`
	Comment  string `json:"comment"`
}
