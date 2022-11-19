package mysql

import (
	"context"
	"errors"
	"exercise/domain"
	"exercise/domain/entity"
	"gorm.io/gorm"
)

type mysqlExerciseRepository struct {
	DB *gorm.DB
}

func NewMysqlExerciseRepository(DB *gorm.DB) *mysqlExerciseRepository {
	return &mysqlExerciseRepository{DB: DB}
}

func (m *mysqlExerciseRepository) FetchByID(ctx context.Context, id int) (*entity.Exercise, error) {
	var exercise entity.Exercise
	err := m.DB.WithContext(ctx).Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = domain.ErrNotFound
	}

	return &exercise, err
}

func (m *mysqlExerciseRepository) StoreExercise(ctx context.Context, exercise *entity.Exercise) error {
	err := m.DB.WithContext(ctx).Create(&exercise).Error
	return err
}

func (m *mysqlExerciseRepository) StoreQuestion(ctx context.Context, question *entity.Question) error {
	question.Score = 10
	err := m.DB.WithContext(ctx).Create(&question).Error
	return err
}

func (m *mysqlExerciseRepository) StoreAnswer(ctx context.Context, answer *entity.Answer) error {
	err := m.DB.WithContext(ctx).Create(&answer).Error
	return err
}

func (m *mysqlExerciseRepository) FetchQuestionByExerciseID(ctx context.Context, idExercise int, idQuestion int) (*entity.Question, error) {
	var question entity.Question
	err := m.DB.WithContext(ctx).Where("id = ? AND exercise_id = ?", idQuestion, idExercise).Take(&question).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = domain.ErrNotFound
	}

	return &question, err
}

func (m *mysqlExerciseRepository) FetchAnswerByExerciseID(ctx context.Context, idExercise int, idUser int) ([]entity.Answer, error) {
	var answers []entity.Answer
	err := m.DB.WithContext(ctx).Where("exercise_id = ? AND user_id = ?", idExercise, idUser).Find(&answers).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = domain.ErrNotFound
	}

	return answers, err
}
