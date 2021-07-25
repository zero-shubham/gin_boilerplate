package main

import (
	"basic/api"
	"basic/config"
	"basic/core"
	accesscontrol "basic/libs/access_control"
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

	err = core.InitDB(
		postgres.Config{
			DSN: cfg.DbUri,
		},
	)
	if err != nil {
		fmt.Println(fmt.Errorf("error initiating database: %s", err))
	}

	err = accesscontrol.SetupCasbin(core.DB)
	if err != nil {
		fmt.Println(fmt.Errorf("error initiating casbin: %s", err))
	}
	// core.Migrate(core.DB)

	router.Run(fmt.Sprintf(":%s", cfg.ServerPort))
}
