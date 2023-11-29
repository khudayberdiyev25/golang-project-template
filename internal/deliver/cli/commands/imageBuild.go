/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang-project-template/cmd/app"
	"golang-project-template/internal/bootstrap"
	"golang-project-template/internal/deliver/cli/controller"
	"golang-project-template/internal/usecase"
)

var imageController = controller.ImageCLIController{UseCase: usecase.NewImageUseCase(bootstrap.SetupDB())}

// imageBuildCmd represents the imageBuild command
var imageBuildCmd = &cobra.Command{
	Use:   "imageBuild",
	Short: "A brief description of your command",
	Long:  `Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("command1 called")
	},
}

func init() {
	imageBuildCmd.Flags().StringP("request", "r", "", "request body")
	app.RootCmd.AddCommand(imageBuildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imageBuildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imageBuildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
