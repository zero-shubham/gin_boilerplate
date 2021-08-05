package cmd

import (
	"basic/controllers"
	"basic/core/schemas"
	"fmt"

	"github.com/spf13/cobra"
)

func InitSuperadminCommands(rootCmd *cobra.Command) error {
	rootCmd.AddCommand(createSuperadminCmd())
	return nil
}

func createSuperadminCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "createsuperadmin",
		Short: "Creates superadmin user to app.",
		Long:  "Creates a superadmin user to app's casbin setup.",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) < 2 {
				fmt.Printf("not enough arguments, username and password required")
			}
			user, err := controllers.CreateUser(
				&schemas.CreateUser{
					Username: args[0],
					Password: args[1],
				},
				[]string{"root"},
			)
			if err != nil {
				fmt.Printf("something went wrong creating user: %v", err.Err)
			}

			fmt.Printf("user added with details: %v", user)
		},
	}
}
