/*
Copyright © 2025 Khorkov Dmitriy <dkhorkov.dev@gmail.com>
*/
package cmd

import (
	"fmt"
	"hmtm-cli/createservice"

	"github.com/spf13/cobra"
)

// createServiceCmd represents the create command.
var createServiceCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new service structure",
	Long: `
		Creates a new service boilerplate. All necessary base directories and files will be created.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic("Please provide a service name")
		}

		serviceName := args[0]
		err := createservice.Init(serviceName)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Successfully created \"%s\" service\n\n", serviceName)
	},
}

func init() {
	rootCmd.AddCommand(createServiceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
