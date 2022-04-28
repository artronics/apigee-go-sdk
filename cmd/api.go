package cmd

import (
	"fmt"
	"github.com/artronics/apigee/api"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"strconv"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Manage API resource",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var apiData api.Api

		getApi := func() string {
			apiData.Organization.Name = cmd.Flags().Lookup("organization").Value.String()

			resBody, err := api.Get(config, api.ApigeeApi, apiData)
			if err != nil {
				log.Fatal(err.Error())
			}
			body, err := ioutil.ReadAll(resBody)
			if err != nil {
				log.Fatal(err.Error())
			}

			return string(body)
		}

		switch cmd.Parent() {
		case getCmd:
			apiData.Name = cmd.Flags().Lookup("name").Value.String()
			fmt.Printf(getApi())

		case listCmd:
			iMetadata, _ := strconv.ParseBool(cmd.Flags().Lookup("includeMetaData").Value.String())
			iRevisions, _ := strconv.ParseBool(cmd.Flags().Lookup("includeRevisions").Value.String())
			apiData.IncludeMetaData = iMetadata
			apiData.IncludeRevisions = iRevisions

			fmt.Printf(getApi())
		case createCmd:
			log.Printf("creating api")

		default:
			panic("unreachable code: command does not exist")
		}
	},
}

func init() {
	getApiCmd := *apiCmd
	commonFlags(&getApiCmd)
	getApiCmd.Flags().StringP("name", "n", "", "API name")
	_ = getApiCmd.MarkFlagRequired("name")

	listApiCmd := *apiCmd
	commonFlags(&listApiCmd)
	listApiCmd.Flags().Bool("includeMetaData", false, "include metadata")
	listApiCmd.Flags().Bool("includeRevisions", false, "include revisions")

	createApiCmd := *apiCmd
	commonFlags(&createApiCmd)
	createApiCmd.Flags().StringP("name", "n", "", "API name")
	_ = createApiCmd.MarkFlagRequired("name")

	getCmd.AddCommand(&getApiCmd)
	listCmd.AddCommand(&listApiCmd)
	createCmd.AddCommand(&createApiCmd)
}

func commonFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("organization", "o", "", "Apigee account organization")
	_ = cmd.MarkFlagRequired("organization")
}
