package controller

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"golang-project-template/internal/domain"
	"log"
)

type ImageCLIController struct {
	UseCase domain.ImageUseCase
}

func (i *ImageCLIController) ImageBuild(cmd *cobra.Command, args []string) {
	fmt.Println("imageBuild called")
	request, err := cmd.Flags().GetString("request")
	if err != nil {
		log.Fatal(err)
	}

	var imageRequest domain.ImageRequest
	err = json.Unmarshal([]byte(request), &imageRequest)
	if err != nil {
		log.Fatal(err)
	}
	response, err := i.UseCase.Create(&imageRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", response)
}
