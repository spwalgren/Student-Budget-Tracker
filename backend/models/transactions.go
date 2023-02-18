package models

type Transaction struct {
	UserID					uint		`json:"userID"`
	Amount	    float32	`json:"amount"`
	Name		    string `json:"name"`
	Date        string `json:"date"`
	Category    string `json:"category"`
	Description string `json:"description"`
}
