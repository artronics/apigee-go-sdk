package cmd

import (
	"github.com/artronics/apigee/pkg"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

var config pkg.ApigeeConfig

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "apigee",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Token = cmd.Flags().Lookup("token").Value.String()
		config.BaseUrl = cmd.Flags().Lookup("base-url").Value.String()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("token", "t", "", "Apigee Access Token")
	_ = rootCmd.MarkPersistentFlagRequired("token")

	rootCmd.PersistentFlags().String("base-url", "https://api.enterprise.apigee.com/v1", "Apigee api base url")

	// FIXME: viper part doesn't work. It should map APIGEE_TOKEN to token
	err := viper.BindPFlag("token", rootCmd.Flags().Lookup("token"))
	if err != nil {
		//log.Printf(err.Error())
	}
}

func initConfig() {
	viper.SetEnvPrefix("apigee")
	viper.AutomaticEnv() // read in environment variables that match
}
