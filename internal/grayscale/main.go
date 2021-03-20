package imagene

import (
	"image"
)

// Toolkit is the interface that describes a grayscale image manipulation
// toolkit object.
type Toolkit interface {
	AddEffect() image.Image
}

// toolkit is the object that satisfies the Toolkit interface.
type toolkit struct {
	image image.Image
}

// NewToolkit returns a new gray manipulation toolkit object that
// implements the Toolkit interface.
func NewToolkit(image image.Image) Toolkit {
	return &toolkit{
		image: image,
	}
}

// AddEffect adds a grayscale effect to the image and returns the new image
// that has the grayscale effect.
func (gt *toolkit) AddEffect() image.Image {
	grayImage := image.NewGray(gt.image.Bounds())
	for y := gt.image.Bounds().Min.Y; y < gt.image.Bounds().Max.Y; y++ {
		for x := gt.image.Bounds().Min.X; x < gt.image.Bounds().Max.X; x++ {
			grayImage.Set(x, y, gt.image.At(x, y))
		}
	}
	return grayImage
}
