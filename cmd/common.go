package cmd

import "github.com/spf13/cobra"

func commonFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("organization", "o", "", "Apigee account organization")
	_ = cmd.MarkFlagRequired("organization")
}
