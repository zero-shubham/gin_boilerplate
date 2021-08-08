package cmd

import (
	"basic/services"
	"fmt"

	"github.com/spf13/cobra"
)

func InitPolicyCommands(rootCmd *cobra.Command) error {
	rootCmd.AddCommand(createPolicyCmd())
	return nil
}

func createPolicyCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "createpolicy",
		Short: "Adds policy to app.",
		Long:  "Adds a policy to app's casbin setup.",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) < 3 {
				fmt.Printf("not enough arguments, sub, obj and act required")
				return
			}

			enfrcr, err := services.GetEnforcer()
			if err != nil {
				fmt.Printf("something went wrong creating user: %v", err)
				return
			}

			success, err := enfrcr.AddPolicy(args[0], args[1], args[2])
			if err != nil {
				fmt.Printf("something went wrong creating user: %v", err)
				return
			}
			if success {
				fmt.Println("Added policy successfully.")
			}
		},
	}
}
