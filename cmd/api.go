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
		var apiData api.ApiData
		apiData.Token = cmd.Flags().Lookup("token").Value.String()
		apiData.BaseUrl = cmd.Flags().Lookup("base-url").Value.String()
		apiData.Organization.Name = cmd.Flags().Lookup("organization").Value.String()

		getApi := func() string {
			resBody, err := api.Get(api.Api, apiData)
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
			apiData.Name = cmd.Flags().Lookup("name").Value.String()
			apiData.Action = cmd.Flags().Lookup("action").Value.String()
			apiData.ZipBundle = cmd.Flags().Lookup("bundle").Value.String()
			resBody, err := api.Create(api.Api, apiData)
			if err != nil {
				log.Fatal(err.Error())
			}
			body, err := ioutil.ReadAll(resBody)
			if err != nil {
				log.Fatal(err.Error())
			}

			fmt.Println(string(body))

		default:
			panic("unreachable code: command does not exist")
		}
	},
}

func init() {
	addNameFlag := func(cmd *cobra.Command, required bool) {
		cmd.Flags().StringP("name", "n", "", "Name of the API proxy. Restrict the characters used to: A-Za-z0-9._-")
		if required {
			_ = cmd.MarkFlagRequired("name")
		}
	}

	// get
	getApiCmd := *apiCmd
	commonFlags(&getApiCmd)
	addNameFlag(&getApiCmd, true)

	// list
	listApiCmd := *apiCmd
	commonFlags(&listApiCmd)
	listApiCmd.Flags().Bool("includeMetaData", false, "include metadata")
	listApiCmd.Flags().Bool("includeRevisions", false, "include revisions")

	// create
	createApiCmd := *apiCmd
	commonFlags(&createApiCmd)
	addNameFlag(&createApiCmd, true)
	createApiCmd.Flags().String("action", "", `Action to perform when importing an API proxy configuration bundle. Set this parameter to one of "import" or "validate"`)
	_ = createApiCmd.MarkFlagRequired("action")
	createApiCmd.Flags().String("bundle", "", "path to the zip file of proxy bundle")
	_ = createApiCmd.MarkFlagRequired("bundle")

	// add commands
	getCmd.AddCommand(&getApiCmd)
	listCmd.AddCommand(&listApiCmd)
	createCmd.AddCommand(&createApiCmd)
}

func commonFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("organization", "o", "", "Apigee account organization")
	_ = cmd.MarkFlagRequired("organization")
}
