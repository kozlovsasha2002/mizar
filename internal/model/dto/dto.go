package dto

import (
	"errors"
)

type CreationEventDto struct {
	Link         string `json:"link"`
	DeadlineDate string `json:"deadlineDate"`
	Description  string `json:"description"`
}

func (c *CreationEventDto) Validate() error {
	if c.Link == "" || c.DeadlineDate == "" || c.Description == "" {
		return errors.New("empty fields")
	}
	return nil
}

type UpdateEventDto struct {
	Link         *string `json:"link"`
	DeadlineDate *string `json:"deadlineDate"`
	IsCompleted  *bool   `json:"isCompleted,omitempty"`
	Description  *string `json:"description"`
}

func (u *UpdateEventDto) Validate() error {
	if u.Link == nil && u.DeadlineDate == nil && u.IsCompleted == nil && u.Description == nil {
		return errors.New("update structure is empty")
	}
	return nil
}
