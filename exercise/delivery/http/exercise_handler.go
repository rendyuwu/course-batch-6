package http

import (
	"exercise/domain/entity"
	"exercise/domain/web"
	"exercise/exercise/delivery/http/middleware"
	"exercise/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ExerciseHandler struct {
	ExerciseUsecase entity.ExerciseUsecase
}

func NewExerciseHandler(r *gin.Engine, exerciseUsecase entity.ExerciseUsecase) {
	handler := &ExerciseHandler{ExerciseUsecase: exerciseUsecase}

	r.GET("/exercises/:id", middleware.WithAuth(), handler.GetExerciseByID)
	r.GET("/exercise/:exercisesId/score", middleware.WithAuth(), handler.GetScore)
	r.POST("/exercises", middleware.WithAuth(), handler.CreateExercise)
	r.POST("/exercises/:exercisesId/questions", middleware.WithAuth(), handler.CreateQuestion)
	r.POST("/exercises/:exercisesId/questions/:questionsId/answer", middleware.WithAuth(), handler.CreateAnswer)
}

func (e *ExerciseHandler) GetExerciseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{Message: "BAD REQUEST"})
		return
	}

	exercise, err := e.ExerciseUsecase.FetchByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(helper.GetStatusCode(err), web.ResponseError{Message: err.Error()})
		return
	}

	// convert exercise to response exercise(filtered response)
	var response web.ResponseExercise
	helper.CopyStruct(&response, &exercise)

	c.JSON(http.StatusOK, response)
}

func (e *ExerciseHandler) CreateExercise(c *gin.Context) {
	var exercise entity.CreateExercise
	if err := c.ShouldBind(&exercise); err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{Message: "BAD REQUEST"})
		return
	}

	result, err := e.ExerciseUsecase.StoreExercise(c.Request.Context(), &exercise)
	if err != nil {
		c.JSON(helper.GetStatusCode(err), web.ResponseError{Message: err.Error()})
	}

	// convert exercise to response exercise(filtered response)
	var response web.ResponseExercise
	helper.CopyStruct(&response, &result)

	c.JSON(http.StatusCreated, response)
}

func (e *ExerciseHandler) CreateQuestion(c *gin.Context) {
	var question entity.CreateQuestion

	exerciseIdString := c.Param("exercisesId")
	exerciseID, err := strconv.Atoi(exerciseIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{Message: "BAD REQUEST"})
		return
	}

	question.ExerciseID = exerciseID
	if err := c.ShouldBind(&question); err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{Message: "BAD REQUEST"})
		return
	}

	userID := c.Request.Context().Value("user_id").(int)
	if err := e.ExerciseUsecase.StoreQuestion(c.Request.Context(), &question, userID); err != nil {
		c.JSON(helper.GetStatusCode(err), err.Error())
		return
	}

	c.JSON(http.StatusCreated, web.Response{Message: "Question created"})
}

func (e *ExerciseHandler) CreateAnswer(c *gin.Context) {
	var answer entity.CreateAnswer

	exerciseIdString := c.Param("exercisesId")
	exerciseID, err := strconv.Atoi(exerciseIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{Message: "BAD REQUEST"})
		return
	}

	questionIdString := c.Param("questionsId")
	questionID, err := strconv.Atoi(questionIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{Message: "BAD REQUEST"})
		return
	}

	answer.ExerciseID = exerciseID
	answer.QuestionID = questionID
	if err := c.ShouldBind(&answer); err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{Message: "BAD REQUEST"})
		return
	}

	userID := c.Request.Context().Value("user_id").(int)
	if err := e.ExerciseUsecase.StoreAnswer(c.Request.Context(), &answer, userID); err != nil {
		c.JSON(helper.GetStatusCode(err), web.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, web.Response{Message: "success"})
}

func (e *ExerciseHandler) GetScore(c *gin.Context) {
	exerciseId, err := strconv.Atoi(c.Param("exercisesId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{Message: "BAD REQUEST"})
		return
	}

	userId := c.Request.Context().Value("user_id").(int)
	score, err := e.ExerciseUsecase.GetScore(c.Request.Context(), userId, exerciseId)
	if err != nil {
		c.JSON(helper.GetStatusCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, web.Response{Message: score})
}
