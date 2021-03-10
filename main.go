package imagene

import (
	"image"
)

// GrayToolkit is the interface that describes a gray manipulation
// toolkit object.
type GrayToolkit interface {
	AddEffect() image.Image
}

// grayToolkit is the object that satisfies the GrayToolki interface.
type grayToolkit struct {
	image image.Image
}

// NewGrayToolkit returns a new gray manipulation toolkit object that
// implements the GrayToolkit interface.
func NewGrayToolkit(image image.Image) GrayToolkit {
	return &grayToolkit{
		image: image,
	}
}

// AddEffect adds a gray effect to the image and returns the new image
// that has the gray effect.
func (gt *grayToolkit) AddEffect() image.Image {
	grayImage := image.NewGray(gt.image.Bounds())
	for y := gt.image.Bounds().Min.Y; y < gt.image.Bounds().Max.Y; y++ {
		for x := gt.image.Bounds().Min.X; x < gt.image.Bounds().Max.X; x++ {
			grayImage.Set(x, y, gt.image.At(x, y))
		}
	}
	return grayImage
}
