package request

import "time"

type FindByID struct {
	UserID string `json:"userID,omitempty"`
}

type TestAddTranslateQuery struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Fullname  string     `json:"fullname"`
	Available *bool      `json:"available"`
}
