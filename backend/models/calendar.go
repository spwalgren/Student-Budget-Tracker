package models

type EventContent struct {
	Frequency   Period  `json:"frequency"`
	StartDate   string  `json:"startDate"`
	EndDate     string  `json:"endDate"`
	Category    string  `json:"category"`
	TotalSpent  float32 `json:"totalSpent"`
	AmountLimit float32 `json:"amountLimit"`
}

type Event struct {
	Data    EventContent `json:"data" gorm:"embedded"`
	UserID  uint         `json:"userId"`
	EventID uint         `json:"eventId" gorm:"unique;primaryKey"`
}

type EventsResponse struct {
	Events []Event `json:"events"`
}