package entity

import (
	"context"
	"sync"
	"time"
)

// ExerciseRepository represent the exercise's repository contract
type ExerciseRepository interface {
	FetchByID(ctx context.Context, id int) (*Exercise, error)
	StoreExercise(ctx context.Context, exercise *Exercise) error
	StoreQuestion(ctx context.Context, question *Question) error
	StoreAnswer(ctx context.Context, answer *Answer) error
	FetchQuestionByExerciseID(ctx context.Context, idExercise int, idQuestion int) (*Question, error)
	FetchAnswerByExerciseID(ctx context.Context, idExercise int, idUser int) ([]Answer, error)
}

// ExerciseUsecase represent the exercise's usecase contract
type ExerciseUsecase interface {
	FetchByID(ctx context.Context, id int) (*Exercise, error)
	StoreExercise(ctx context.Context, exercise *CreateExercise) (*Exercise, error)
	StoreQuestion(ctx context.Context, question *CreateQuestion, creatorID int) error
	StoreAnswer(ctx context.Context, answer *CreateAnswer, userID int) error
	GetScore(ctx context.Context, userID int, exerciseID int) (string, error)
}

type Exercise struct {
	ID          int        `json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Questions   []Question `json:"questions,omitempty"`
}

type Question struct {
	ID            int
	ExerciseID    int
	Body          string
	OptionA       string
	OptionB       string
	OptionC       string
	OptionD       string
	CorrectAnswer string
	Score         int
	CreatorID     int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Answer struct {
	ID         int
	ExerciseID int
	QuestionID int
	UserID     int
	Answer     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CreateExercise struct {
	Title       string `json:"title,omitempty" binding:"required"`
	Description string `json:"description,omitempty" binding:"required"`
}

type CreateQuestion struct {
	ExerciseID    int    `json:"exercise_id,omitempty" binding:"required"`
	Body          string `json:"body,omitempty" binding:"required"`
	OptionA       string `json:"option_a,omitempty" binding:"required"`
	OptionB       string `json:"option_b,omitempty" binding:"required"`
	OptionC       string `json:"option_c,omitempty" binding:"required"`
	OptionD       string `json:"option_d,omitempty" binding:"required"`
	CorrectAnswer string `json:"correct_answer,omitempty" binding:"required"`
}

type CreateAnswer struct {
	ExerciseID int    `json:"exercise_id,omitempty" binding:"required"`
	QuestionID int    `json:"question_id,omitempty" binding:"required"`
	Answer     string `json:"answer,omitempty" binding:"required"`
}

type Score struct {
	Total int
	mu    sync.Mutex
}

func (s *Score) Inc(value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Total += value
}
