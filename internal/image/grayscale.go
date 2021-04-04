package img

import "image"

// toolkit is the object that satisfies the Toolkit interface.
type Toolkit struct {
	image image.Image
}

// NewToolkit returns a new gray manipulation toolkit object that
// implements the Toolkit interface.
func NewToolkit(image image.Image) Toolkit {
	return Toolkit{
		image: image,
	}
}

// AddEffect adds a grayscale effect to the image and returns the new image
// that has the grayscale effect.
func (gt *Toolkit) AddEffect() image.Image {
	grayImage := NewGray(gt.Bounds())
	for y := gt.Bounds().Min.Y; y < gt.Bounds().Max.Y; y++ {
		for x := gt.Bounds().Min.X; x < gt.Bounds().Max.X; x++ {
			graySet(x, y, gt.At(x, y))
		}
	}
	return grayImage
}
