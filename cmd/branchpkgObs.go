/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	obsapi "github.com/ashpect/obs-cli/obs-api/endpoints"
	"github.com/spf13/cobra"
)

// branchpkgObsCmd represents the branchpkgObs command
var branchpkgObsCmd = &cobra.Command{
	Use:   "branchpkgObs",
	Short: "Branches pkg to another pkg under a subproject",
	Long:  `This command is used to branch a pkg to another pkg under a subproject.`,
	Run: func(cmd *cobra.Command, args []string) {

		project_name, _ := cmd.Flags().GetString("project")
		package_name, _ := cmd.Flags().GetString("package")
		subproject_name, _ := cmd.Flags().GetString("subproject")

		// fmt.Println("Branching pkg to another pkg under a subproject")
		// fmt.Println("Project:", project_name)
		// fmt.Println("Package:", package_name)
		// fmt.Println("Subproject:", subproject_name)
		obsapi.Branchpkg(project_name, package_name, subproject_name)
	},
}

func init() {
	rootCmd.AddCommand(branchpkgObsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// branchpkgObsCmd.PersistentFlags().String("foo", "", "A help for foo")

	branchpkgObsCmd.Flags().StringP("project", "p", "", "specify project name")
	branchpkgObsCmd.Flags().StringP("package", "k", "", "specify pkg name")
	branchpkgObsCmd.Flags().StringP("subproject", "u", "", "specify target subproject name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// branchpkgObsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
