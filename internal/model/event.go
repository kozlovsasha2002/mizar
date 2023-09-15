package model

// Event TODO: нужно сделать ввод даты строго определенного формата
type Event struct {
	EventId      int    `json:"eventId"`
	Link         string `json:"link"`
	DeadlineDate string `json:"deadlineDate"`
	DeadlineTime string `json:"deadlineTime"`
	IsCompleted  bool   `json:"isCompleted"`
	Description  string `json:"description"`
}
