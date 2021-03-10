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
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wisdommatt/imagene"
)

// grayscaleCmd represents the grayscale command
var grayscaleCmd = &cobra.Command{
	Use:   "grayscale",
	Short: "Converts an image to grayscale",
	Run: func(cmd *cobra.Command, args []string) {
		var grayToolkit imagene.GrayToolkit
		var imageFile image.Image
		local, err := cmd.Flags().GetString("local")
		if err != nil {
			fmt.Println(err)
			return
		}
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			fmt.Println(err)
			return
		}
		if local == "" && url == "" {
			fmt.Println("Please provide the image path or url using --local or --url respectively")
			return
		}
		switch {
		case local != "":
			file, err := ioutil.ReadFile(local)
			if err != nil {
				fmt.Println("An error occured: ", err.Error())
				return
			}
			imageFile, _, err = image.Decode(bytes.NewBuffer(file))
			if err != nil {
				fmt.Println("An error occured: ", err.Error())
				return
			}

		case url != "" && local == "":
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("An error occured: ", err)
				return
			}
			defer resp.Body.Close()
			imageFile, _, err = image.Decode(resp.Body)
			if err != nil {
				fmt.Println("An error occured: ", err.Error())
				return
			}
		}

		grayToolkit = imagene.NewGrayToolkit(imageFile)
		imageFile = grayToolkit.AddEffect()

		output, err := cmd.Flags().GetString("output")
		if err != nil || output == "" {
			fmt.Println("Please try again and specify the output using -o flag")
			return
		}
		outputFileExtension := getFileExtension(output)
		if outputFileExtension == "" {
			fmt.Println("Please add an extension to your output filename e.g img.jpeg")
			return
		}
		outputFile, err := os.Create(output)
		if err != nil {
			fmt.Println("An error occured: ", err.Error())
			return
		}
		defer outputFile.Close()

		switch {
		case outputFileExtension == "png":
			err = png.Encode(outputFile, imageFile)
			if err != nil {
				fmt.Println("An error occured: ", err.Error())
				return
			}

		case outputFileExtension == "jpg" || outputFileExtension == "jpeg":
			err = jpeg.Encode(outputFile, imageFile, nil)
			if err != nil {
				fmt.Println("An error occured: ", err.Error())
				return
			}
		}
		fmt.Println("Image converted to grayscale successfully !")
	},
}

// getFileExtension returns a file extension from the file name.
func getFileExtension(filename string) string {
	splittedStr := strings.Split(filename, ".")
	return splittedStr[len(splittedStr)-1]
}

func init() {
	grayscaleCmd.Flags().String("local", "", "Used to specify the local path of the image")
	grayscaleCmd.Flags().String("url", "", "Used to specify the URL of the image")
	grayscaleCmd.Flags().StringP("output", "o", "", "Used to specify the output of the new grayscale image")
	rootCmd.AddCommand(grayscaleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grayscaleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grayscaleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
