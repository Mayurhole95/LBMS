package server

import (
	"github.com/Mayurhole95/Library-Management-System/app"
	"github.com/Mayurhole95/Library-Management-System/db"
	"github.com/Mayurhole95/Library-Management-System/user"
)

type dependencies struct {
	UserService user.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	userService := user.NewService(dbStore, logger)

	return dependencies{
		UserService: userService,
	}, nil
}
