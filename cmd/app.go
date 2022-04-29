package cmd

import (
	"fmt"
	"github.com/artronics/apigee/resource"
	"io"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var appData resource.AppData
		appData.Token = cmd.Flags().Lookup("token").Value.String()
		appData.BaseUrl = cmd.Flags().Lookup("base-url").Value.String()
		appData.Organization.Name = cmd.Flags().Lookup("organization").Value.String()

		var body io.ReadCloser
		var err error

		switch cmd.Parent() {
		case getCmd:
			appData.Id = cmd.Flags().Lookup("id").Value.String()

			body, err = resource.Get(resource.App, appData)

		case listCmd:
			body, err = resource.List(resource.App, appData)
		default:
			panic("unreachable code: command does not exist")
		}

		if err != nil {
			log.Fatal(err.Error())
		}
		bodyBuf, err := ioutil.ReadAll(body)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(string(bodyBuf))
	},
}

func init() {
	addIdFlag := func(cmd *cobra.Command, required bool) {
		cmd.Flags().String("id", "", "App ID (UUID)")
		if required {
			_ = cmd.MarkFlagRequired("id")
		}
	}
	// get
	getAppCmd := *appCmd
	commonFlags(&getAppCmd)
	addIdFlag(&getAppCmd, true)

	getCmd.AddCommand(&getAppCmd)

	// list
	listAppCmd := *appCmd
	commonFlags(&listAppCmd)

	listCmd.AddCommand(&listAppCmd)
}
