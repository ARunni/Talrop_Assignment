package domain

type User struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	PhoneNumber string  `json:"phone_number"`
	Country     string  `json:"country"`
	Score       float64 `json:"score"`
}

type Response struct {
	Total       int     `json:"total"`
	Results     []User  `json:"results"`
	TimeTakenMs float64 `json:"time_taken_ms"`
}


