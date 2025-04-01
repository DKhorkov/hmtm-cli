/*
Copyright © 2025 Khorkov Dmitriy <dkhorkov.dev@gmail.com>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/DKhorkov/hmtm-cli/internal/createservice"
)

// createServiceCmd represents the create command.
var createServiceCmd = &cobra.Command{
	Use:   "create-service <service-name>",
	Args:  cobra.ExactArgs(1), // Only service name arg is expected
	Short: "Creates a new service structure",
	Long: `Creates a new service boilerplate. 
All necessary base directories and files will be created.
`,
	Run: func(_ *cobra.Command, args []string) {
		serviceName := strings.ToLower(args[0])
		err := createservice.Create(serviceName)
		if err != nil {
			fmt.Printf("Failer to create \"%s\" service\n. Error: %v\n", serviceName, err)
		}

		fmt.Printf("Successfully created \"%s\" service\n\n", serviceName)
	},
}

func init() {
	rootCmd.AddCommand(createServiceCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createServiceCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createServiceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
