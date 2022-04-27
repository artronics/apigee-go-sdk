package cmd

import (
	"github.com/artronics/apigee/pkg/api"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Manage API resource",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.Organization = cmd.Flags().Lookup("organization").Value.String()
	},
	Run: func(cmd *cobra.Command, args []string) {
		switch cmd.Parent() {
		case getCmd:
			a, err := api.List(config, "api")
			if err != nil {
				log.Fatal(err.Error())
			}
			body, err := ioutil.ReadAll(a)
			if err != nil {
				log.Fatal(err.Error())
			}

			log.Printf(string(body))
		case createCmd:
			log.Printf("creating api")

		default:
			panic("unreachable code: command does not exist")
		}
	},
}

func init() {
	apiCmd.Flags().StringP("name", "n", "", "API name")
	_ = apiCmd.MarkFlagRequired("name")
	apiCmd.Flags().StringP("organization", "o", "", "Apigee account organization")
	_ = apiCmd.MarkFlagRequired("organization")

	getApiCmd := *apiCmd
	createApiCmd := *apiCmd
	createCmd.AddCommand(&createApiCmd)
	getCmd.AddCommand(&getApiCmd)
}
