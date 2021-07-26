package main

import (
	"basic/api"
	"basic/config"
	accesscontrol "basic/libs/access_control"
	"basic/services"
	"fmt"

	"gorm.io/driver/postgres"
)

func main() {
	// ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("error initiating router: %s", err))
	}

	router, err := api.NewRouter("v1")
	if err != nil {
		fmt.Println(fmt.Errorf("error initiating router: %s", err))
	}

	err = services.InitDB(
		postgres.Config{
			DSN: cfg.DbUri,
		},
	)
	if err != nil {
		fmt.Println(fmt.Errorf("error initiating database: %s", err))
	}

	db, err := services.GetDB()
	if err != nil {
		fmt.Println(err)
	} else {
		err = accesscontrol.SetupCasbin(db)
		if err != nil {
			fmt.Println(fmt.Errorf("error initiating casbin: %s", err))
		}
	}
	// core.Migrate(core.DB)

	router.Run(fmt.Sprintf(":%s", cfg.ServerPort))
}
