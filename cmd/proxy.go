package cmd

import (
	"fmt"
	"github.com/artronics/apigee/resource"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"log"
	"strconv"
)

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Manage API proxy resource",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var proxyData resource.ProxyData
		proxyData.Token = cmd.Flags().Lookup("token").Value.String()
		proxyData.BaseUrl = cmd.Flags().Lookup("base-url").Value.String()
		proxyData.Organization.Name = cmd.Flags().Lookup("organization").Value.String()

		var body io.ReadCloser
		var err error

		switch cmd.Parent() {
		case getCmd:
			proxyData.Name = cmd.Flags().Lookup("name").Value.String()

			body, err = resource.Get(resource.Proxy, proxyData)

		case listCmd:
			iMetadata, _ := strconv.ParseBool(cmd.Flags().Lookup("includeMetaData").Value.String())
			iRevisions, _ := strconv.ParseBool(cmd.Flags().Lookup("includeRevisions").Value.String())
			proxyData.IncludeMetaData = iMetadata
			proxyData.IncludeRevisions = iRevisions

			body, err = resource.List(resource.Proxy, proxyData)

		case createCmd:
			proxyData.Name = cmd.Flags().Lookup("name").Value.String()
			proxyData.Action = cmd.Flags().Lookup("action").Value.String()
			proxyData.ZipBundle = cmd.Flags().Lookup("bundle").Value.String()

			body, err = resource.Create(resource.Proxy, proxyData)

		case deleteCmd:
			proxyData.Name = cmd.Flags().Lookup("name").Value.String()

			body, err = resource.Delete(resource.Proxy, proxyData)
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
	addNameFlag := func(cmd *cobra.Command, required bool) {
		cmd.Flags().StringP("name", "n", "", "Name of the API proxy. Restrict the characters used to: A-Za-z0-9._-")
		if required {
			_ = cmd.MarkFlagRequired("name")
		}
	}

	// get
	getProxyCmd := *proxyCmd
	commonFlags(&getProxyCmd)
	addNameFlag(&getProxyCmd, true)

	getCmd.AddCommand(&getProxyCmd)

	// list
	listProxyCmd := *proxyCmd
	listProxyCmd.Aliases = append(listProxyCmd.Aliases, "proxies")
	commonFlags(&listProxyCmd)
	listProxyCmd.Flags().Bool("includeMetaData", false, "include metadata")
	listProxyCmd.Flags().Bool("includeRevisions", false, "include revisions")

	listCmd.AddCommand(&listProxyCmd)

	// create
	createProxyCmd := *proxyCmd
	commonFlags(&createProxyCmd)
	addNameFlag(&createProxyCmd, true)
	createProxyCmd.Flags().String("action", "", `Action to perform when importing an API proxy configuration bundle. Set this parameter to one of "import" or "validate"`)
	_ = createProxyCmd.MarkFlagRequired("action")
	createProxyCmd.Flags().String("bundle", "", "path to the zip file of proxy bundle")
	_ = createProxyCmd.MarkFlagRequired("bundle")

	createCmd.AddCommand(&createProxyCmd)

	// TODO: update

	// delete
	deleteProxyCmd := *proxyCmd
	commonFlags(&deleteProxyCmd)
	addNameFlag(&deleteProxyCmd, true)

	deleteCmd.AddCommand(&deleteProxyCmd)
}
