package web

import "time"

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message,omitempty"`
}

// Response represent the response default
type Response struct {
	Message string `json:"message,omitempty"`
}

// ResponseSuccess represent the response success login & register
type ResponseSuccess struct {
	Token string `json:"token,omitempty"`
}

type ResponseQuestions struct {
	ID        int       `json:"id,omitempty"`
	Body      string    `json:"body,omitempty"`
	OptionA   string    `json:"option_a,omitempty"`
	OptionB   string    `json:"option_b,omitempty"`
	OptionC   string    `json:"option_c,omitempty"`
	OptionD   string    `json:"option_d,omitempty"`
	Score     int       `json:"score,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseExercise struct {
	ID          int                 `json:"id,omitempty"`
	Title       string              `json:"title,omitempty"`
	Description string              `json:"description,omitempty"`
	Questions   []ResponseQuestions `json:"questions,omitempty"`
}
