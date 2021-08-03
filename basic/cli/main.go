package main

import (
	"basic/cli/cmd"
	"basic/config"
	"basic/services"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
)

func main() {
	cobra.OnInitialize(initConfig)

	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "A cli for the current app.",
		Long:  `This cli helps managing the app.`,
	}

	cfg, err := config.LoadCliConfig()
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
		err = services.SetupCasbin(db)
		if err != nil {
			fmt.Println(fmt.Errorf("error initiating casbin: %s", err))
		}
	}

	cmd.InitSuperadminCommands(rootCmd)

	cobra.CheckErr(rootCmd.Execute())

}

var cfgFile string

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName(".basic")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
