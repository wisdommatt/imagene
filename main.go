package imagene

import "io"

// GrayToolkit is the interface that describes a gray manipulation
// toolkit object.
type GrayToolkit interface{}

// grayToolkit is the object that satisfies the GrayToolki interface.
type grayToolkit struct {
	image io.Reader
}

// NewGrayToolkit returns a new gray manipulation toolkit object that
// implements the GrayToolkit interface.
func NewGrayToolkit(image io.Reader) GrayToolkit {
	return &grayToolkit{
		image: image,
	}
}
