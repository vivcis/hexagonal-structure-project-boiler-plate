package models

type Notification struct {
	Message string `json:"message"`
	Year    int    `json:"year"`
	Month   int    `json:"month"`
	Day     int    `json:"day"`
}

type Message struct {
	Data    string `json:"data"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
}
