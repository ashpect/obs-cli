/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	obsapi "github.com/ashpect/obs-cli/obs-api/endpoints"
	"github.com/spf13/cobra"
)

// architecturesObsCmd represents the architecturesObs command
var architecturesObsCmd = &cobra.Command{
	Use:   "architecturesObs",
	Short: "List of known architectures in the obs-api",
	Long:  `This command is used to list the known architectures in the obs-api.`,
	Run: func(cmd *cobra.Command, args []string) {
		obsapi.Architectures()
	},
}

func init() {
	rootCmd.AddCommand(architecturesObsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// architecturesObsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// architecturesObsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
