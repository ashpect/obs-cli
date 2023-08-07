package cmd

import (
	"fmt"

	obsapi "github.com/ashpect/obs-cli/obs-api/endpoints"
	"github.com/spf13/cobra"
)

// infoObsCmd represents the infoObs command
var buildObsCmd = &cobra.Command{
	Use:   "buildObs",
	Short: "Gets a simple directory listing of all repos in the home:Ashpect project",
	Long:  `This command is used to test the obs-api`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gets a simple directory listing of all repos for the specified project")

		project_name, _ := cmd.Flags().GetString("project")
		//fmt.Println("Project name: ", project_name)

		obsapi.Build(project_name)
	},
}

func init() {
	rootCmd.AddCommand(buildObsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoObsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	buildObsCmd.Flags().StringP("project", "p", "", "project to search repos for")
}
