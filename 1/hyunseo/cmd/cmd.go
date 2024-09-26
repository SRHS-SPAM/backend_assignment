package cmd

import (
	"fmt"
	"github.com/3boku/backend1/repositories"
	"github.com/3boku/backend1/router"
)

type Cmd struct {
	repository *repositories.Repository
	router     *router.Router
}

func NewCmd() *Cmd {
	dsn := fmt.Sprintf("host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Seoul")

	db, err := repositories.NewPostgres(dsn)
	if err != nil {
		panic(err)
	}

	c := &Cmd{
		repository: repositories.NewRepository(db),
		router:     router.NewRouter(),
	}

	err = c.router.ServerStart()
	if err != nil {
		panic(err)
	}

	return c
}
