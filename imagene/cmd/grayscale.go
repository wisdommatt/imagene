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
	"errors"
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
		var err error
		local, _ := cmd.Flags().GetString("local")
		url, _ := cmd.Flags().GetString("url")
		if local == "" && url == "" {
			fmt.Println("Please provide the image path or url using --local or --url respectively")
			return
		}
		// Getting the image file from either the --local or --url.
		imageFile, err = getImageFile(local, url)
		if err != nil {
			fmt.Println("An error occured: ", err.Error())
			return
		}
		// Initializing the gray toolkit with the imageFile.
		grayToolkit = imagene.NewGrayToolkit(imageFile)
		imageFile = grayToolkit.AddEffect()
		// Getting the output path of the processed image from the --output
		// flag.
		output, err := cmd.Flags().GetString("output")
		if err != nil || output == "" {
			fmt.Println("Please try again and specify the output using -o flag")
			return
		}
		// Getting the output file path file extension which will be used
		// later to determing the image type.
		outputFileExtension := getFileExtensionFileName(output)
		if outputFileExtension == "" {
			fmt.Println("Please add an extension to your output filename e.g img.jpeg")
			return
		}
		// outputFile is the file when the grayscale image will be saved.
		outputFile, err := os.Create(output)
		if err != nil {
			fmt.Println("An error occured: ", err.Error())
			return
		}
		defer outputFile.Close()
		// Saving the image file to a file on disk.
		err = saveImageToFile(outputFile, imageFile, outputFileExtension)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Image converted to grayscale successfully !")
	},
}

// getFileExtensionFileName returns a file extension from the file name.
func getFileExtensionFileName(filename string) string {
	splittedStr := strings.Split(filename, ".")
	return splittedStr[len(splittedStr)-1]
}

// getImageFile returns the image file of a local image or a url image
// depending on the flag specified.
func getImageFile(local, url string) (img image.Image, err error) {
	switch {
	// Fetching the image file of a local image file if the --local
	// flag is not empty.
	case local != "":
		img, err = getLocalImageFile(local)
		if err != nil {
			return nil, err
		}
	// Fetching the image file of a url image file if the --url is
	// not empty and --local flag is empty.
	case url != "" && local == "":
		img, err = getURLImageFile(url)
		if err != nil {
			return nil, err
		}
	}
	return
}

// saveImageToFile saves the image with type image.Image to file.
func saveImageToFile(file *os.File, image image.Image, fileExtension string) error {
	switch {
	// saving the file using png.Encode if fileExtension is png.
	case fileExtension == "png":
		err := png.Encode(file, image)
		if err != nil {
			return err
		}
	// saving the file using jpeg.Encode if fileExtension is jpg or jpeg.
	case fileExtension == "jpg" || fileExtension == "jpeg":
		err := jpeg.Encode(file, image, nil)
		if err != nil {
			return err
		}
	// returning an unsupported file type error if file type is not (png, jpg, jpeg).
	default:
		return errors.New(fileExtension + " file type is not supported")
	}
	return nil
}

// getLocalImageFile returns the image file contents of the image specified in the
// --local path after reading the file contents from disk.
func getLocalImageFile(imageLocalPath string) (img image.Image, err error) {
	file, err := ioutil.ReadFile(imageLocalPath)
	if err != nil {
		return nil, errors.New("The provided local image path in invalid")
	}
	img, _, err = image.Decode(bytes.NewBuffer(file))
	return
}

// getURLImageFile returns the image file contents of the image specified  in the
// --url image url after making an API request to get it's contents.
func getURLImageFile(imageURL string) (img image.Image, err error) {
	resp, err := http.Get(imageURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, errors.New("The image url provided is invalid")
	}
	defer resp.Body.Close()
	img, _, err = image.Decode(resp.Body)
	return
}

func init() {
	grayscaleCmd.Flags().String("local", "", "Used to specify the local path of the image")
	grayscaleCmd.Flags().String("url", "", "Used to specify the URL of the image")
	grayscaleCmd.Flags().StringP("output", "o", "", "Used to specify the output of the new grayscale image")
	rootCmd.AddCommand(grayscaleCmd)
}
