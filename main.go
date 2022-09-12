package main

import (
	"os"

	"github.com/Mayurhole95/Library-Management-System/app"
	"github.com/Mayurhole95/Library-Management-System/config"
	"github.com/Mayurhole95/Library-Management-System/db"
	"github.com/Mayurhole95/Library-Management-System/server"
	"github.com/urfave/cli"
)

func main() {
	config.Load()
	app.Init()
	defer app.Close()

	cliApp := cli.NewApp()
	cliApp.Name = "Library Management System"
	cliApp.Version = "1.0.0"
	cliApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start server",
			Action: func(c *cli.Context) error {
				server.StartAPIServer()
				return nil
			},
		},
		{
			Name:  "create_migration",
			Usage: "create_migration file",
			Action: func(c *cli.Context) error {
				return db.CreateMigrationFile(c.Args().Get(0))
			},
		},
		{
			Name:  "migrate",
			Usage: "run db migrations",
			Action: func(c *cli.Context) error {
				err := db.RunMigrations()
				return err
			},
		},
		{
			Name:  "rollback",
			Usage: "Rollback migrations",
			Action: func(c *cli.Context) error {
				return db.RollbackMigrations(c.Args().Get(0))

			},
		},
	}
	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}
