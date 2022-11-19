package usecase

import (
	"context"
	"exercise/domain/entity"
	"exercise/helper"
	"strconv"
	"strings"
	"sync"
)

type exerciseUsecase struct {
	ExerciseRepository entity.ExerciseRepository
}

func NewExerciseUsecase(exerciseRepository entity.ExerciseRepository) *exerciseUsecase {
	return &exerciseUsecase{ExerciseRepository: exerciseRepository}
}

func (e *exerciseUsecase) FetchByID(ctx context.Context, id int) (*entity.Exercise, error) {
	exercise, err := e.ExerciseRepository.FetchByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return exercise, err
}

func (e *exerciseUsecase) StoreExercise(ctx context.Context, exercise *entity.CreateExercise) (*entity.Exercise, error) {
	var result entity.Exercise

	helper.CopyStruct(&result, &exercise)
	err := e.ExerciseRepository.StoreExercise(ctx, &result)

	return &result, err
}

func (e *exerciseUsecase) StoreQuestion(ctx context.Context, question *entity.CreateQuestion, creatorID int) error {
	_, err := e.ExerciseRepository.FetchByID(ctx, question.ExerciseID)
	if err != nil {
		return err
	}

	var result entity.Question

	helper.CopyStruct(&result, &question)
	result.CreatorID = creatorID
	err = e.ExerciseRepository.StoreQuestion(ctx, &result)

	return err
}

func (e *exerciseUsecase) StoreAnswer(ctx context.Context, answer *entity.CreateAnswer, userID int) error {
	_, err := e.ExerciseRepository.FetchQuestionByExerciseID(ctx, answer.ExerciseID, answer.QuestionID)
	if err != nil {
		return err
	}

	var result entity.Answer

	helper.CopyStruct(&result, &answer)
	result.UserID = userID
	err = e.ExerciseRepository.StoreAnswer(ctx, &result)

	return err
}

func (e *exerciseUsecase) GetScore(ctx context.Context, userID int, exerciseID int) (string, error) {
	exercise, err := e.ExerciseRepository.FetchByID(ctx, exerciseID)
	if err != nil {
		return "", err
	}

	answers, err := e.ExerciseRepository.FetchAnswerByExerciseID(ctx, exerciseID, userID)
	if err != nil {
		return "", err
	}

	mapQA := make(map[int]entity.Answer)
	for _, answer := range answers {
		mapQA[answer.QuestionID] = answer
	}

	var score entity.Score
	wg := new(sync.WaitGroup)
	for _, question := range exercise.Questions {
		wg.Add(1)
		go func(question entity.Question) {
			defer wg.Done()
			if strings.EqualFold(question.CorrectAnswer, mapQA[question.ID].Answer) {
				score.Inc(question.Score)
			}
		}(question)
	}

	wg.Wait()

	return strconv.Itoa(score.Total), nil
}
