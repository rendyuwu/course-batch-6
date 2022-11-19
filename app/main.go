package main

import (
	"exercise/config"
	_exerciseHandler "exercise/exercise/delivery/http"
	_exerciseRepo "exercise/exercise/repository/mysql"
	_exerciseUsecase "exercise/exercise/usecase"
	_authorHandler "exercise/user/delivery/http"
	_authorRepo "exercise/user/repository/mysql"
	_authorUsecase "exercise/user/usecase"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	env := config.NewConfig()
	db := config.NewDB(env)

	r := gin.Default()
	ur := _authorRepo.NewMysqlUserRepository(db)
	uu := _authorUsecase.NewUserUsecase(ur)
	er := _exerciseRepo.NewMysqlExerciseRepository(db)
	eu := _exerciseUsecase.NewExerciseUsecase(er)
	_authorHandler.NewUserHandler(r, uu)
	_exerciseHandler.NewExerciseHandler(r, eu)

	log.Fatal(r.Run(":1234"))
}
