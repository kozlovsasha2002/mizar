package model

type Event struct {
	EventId      int    `json:"eventId"`
	Link         string `json:"link"`
	DeadlineDate string `json:"deadlineDate"`
	IsCompleted  bool   `json:"isCompleted"`
	Description  string `json:"description"`
}
