package cmd

import (
	"fmt"

	obsapi "github.com/ashpect/obs-cli/obs-api/endpoints"
	"github.com/spf13/cobra"
)

// infoObsCmd represents the infoObs command
var infoObsCmd = &cobra.Command{
	Use:   "infoObs",
	Short: "Hits the obs-api for testing purposes",
	Long:  `This command is used to test the obs-api`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hits the obs-api for testing purposes")

		obsapi.About()
	},
}

func init() {
	rootCmd.AddCommand(infoObsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoObsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoObsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
