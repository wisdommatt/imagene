package image

import (
	"bytes"
	"errors"
	"image"
	"io/ioutil"
	"net/http"
)

// Reader is the interface that describes an image
// reader object.
type Reader interface {
	ReadFromURL(url string) (img image.Image, err error)
	ReadFromLocalPath(localPath string) (img image.Image, err error)
}

// Writer is the interface that describes an image
// writer object.
type Writer interface{}

// ReadWriter is the interface that groups the
// image Reader and Writer interfaces.
type ReadWriter interface {
	Reader
	Writer
}

// reader is the object that satisfies the Reader
// interface.
type reader struct {
}

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
		return nil, errors.New("The provided local image path in invalid")
	}
	img, _, err = image.Decode(bytes.NewBuffer(file))
	return
}
