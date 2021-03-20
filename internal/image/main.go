package image

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Reader is the interface that describes an image
// reader object.
type Reader interface {
	ReadFromURL(url string) (img image.Image, err error)
	ReadFromLocalPath(localPath string) (img image.Image, err error)
	GetImageFromURLorLocalPath(url, localPath string) (img image.Image, err error)
}

// Writer is the interface that describes an image
// writer object.
type Writer interface {
	WriteToFile(image image.Image, output string) error
}

// ReadWriter is the interface that groups the
// image Reader and Writer interfaces.
type ReadWriter interface {
	Reader
	Writer
}

// reader is the object that satisfies the Reader
// interface.
type reader struct{}

// writer is the object that satisfies the Writer interface.
type writer struct{}

// readWriter is the object that satisfies the Reader and Writer
// interfaces.
type readWriter struct {
	reader
	writer
}

// NewReadWriter returns an image reader object that implements the
// Reader interface.
func NewReader() Reader {
	return reader{}
}

// NewWriter returns an image writer object that implements the
// Writer interface.
func NewWriter() Writer {
	return writer{}
}

// NewReadWriter returns an image object that implements the
// Reader and Writer interface.
func NewReadWriter() ReadWriter {
	return readWriter{}
}

// ReadFromURL reads and return an image from the url.
func (reader) ReadFromURL(url string) (img image.Image, err error) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, errors.New("The image url provided is invalid")
	}
	defer resp.Body.Close()
	img, _, err = image.Decode(resp.Body)
	return
}

// ReadFromLocalPath reads and return an image from a local
// directory file path.
func (reader) ReadFromLocalPath(localPath string) (img image.Image, err error) {
	file, err := ioutil.ReadFile(localPath)
	if err != nil {
		return nil, errors.New("The provided local image path is invalid")
	}
	img, _, err = image.Decode(bytes.NewBuffer(file))
	return
}

// WriteToFile writes an image contents to a file on disk.
func (writer) WriteToFile(image image.Image, output string) error {
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	fileExtension := getFileExtensionFromFileName(output)
	switch {
	// saving the file using png.Encode if fileExtension is png.
	case fileExtension == "png":
		return png.Encode(file, image)

		// saving the file using jpeg.Encode if fileExtension is jpg or jpeg.
	case fileExtension == "jpg" || fileExtension == "jpeg":
		return jpeg.Encode(file, image, nil)

		// returning an unsupported file type error if file type is not (png, jpg, jpeg).
	default:
		return errors.New(fileExtension + " output file type is not supported")
	}
}

// GetImageFromURLorLocalPath returns an image from the url if url is not
// empty or returns an image from the local path if localPath is not empty
// and url is empty.
func (reader reader) GetImageFromURLorLocalPath(url, localPath string) (img image.Image, err error) {
	switch {
	case url != "":
		return reader.ReadFromURL(url)

	case localPath != "":
		return reader.ReadFromLocalPath(localPath)
	}
	return nil, errors.New("Both url and local path was not provided")
}

// getFileExtensionFromFileName returns a file extension from the file name.
func getFileExtensionFromFileName(filename string) string {
	splittedStr := strings.Split(filename, ".")
	return splittedStr[len(splittedStr)-1]
}
