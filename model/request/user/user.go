package user

// FindByID struct
type FindByID struct {
	UserID string `json:"userID,omitempty"`
}

// Find struct
type Find struct {
	Fullname string `json:"fullname"`
}

// Create struct
type Create struct {
	Fullname string `json:"fullname"`
}
