/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	_ "image/jpeg"

	img "github.com/wisdommatt/imagene/internal/image"

	imagene "github.com/wisdommatt/imagene/internal/grayscale"

	"github.com/spf13/cobra"
)

// grayscaleCmd represents the grayscale command
var grayscaleCmd = &cobra.Command{
	Use:   "grayscale",
	Short: "Converts an image to grayscale",
	Run: func(cmd *cobra.Command, args []string) {
		local, _ := cmd.Flags().GetString("local")
		url, _ := cmd.Flags().GetString("url")
		output, _ := cmd.Flags().GetString("output")
		if local == "" && url == "" {
			fmt.Println("Please provide the image path or url using --local or --url respectively")
			return
		}
		if output == "" {
			fmt.Println("Provide provide an output file using the -o flag")
			return
		}
		imageReadWriter := img.NewReadWriter()
		// Getting the image file from either the --local or --url.
		imageFile, err := imageReadWriter.GetImageFromURLorLocalPath(url, local)
		if err != nil {
			fmt.Println("An error occured: ", err.Error())
			return
		}
		// Initializing the gray toolkit with the imageFile.
		grayToolkit := imagene.NewToolkit(imageFile)
		imageFile = grayToolkit.AddEffect()
		// Saving the image file to a file to disk.
		err = imageReadWriter.WriteToFile(imageFile, output)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Image converted to grayscale successfully !")
	},
}

func init() {
	grayscaleCmd.Flags().String("local", "", "Used to specify the local path of the image")
	grayscaleCmd.Flags().String("url", "", "Used to specify the URL of the image")
	grayscaleCmd.Flags().StringP("output", "o", "", "Used to specify the output of the new grayscale image")
	rootCmd.AddCommand(grayscaleCmd)
}
