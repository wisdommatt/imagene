package img

import "image"

// GrayToolkit is the object.
type GrayToolkit struct {
	image image.Image
}

// NewGrayToolkit returns a new gray scale manipulation toolkit object.
func NewGrayToolkit(image image.Image) GrayToolkit {
	return GrayToolkit{
		image: image,
	}
}

// AddEffect adds a grayscale effect to the image and returns the new image
// that has the grayscale effect.
func (gt *GrayToolkit) AddEffect() image.Image {
	grayImage := image.NewGray(gt.image.Bounds())
	for y := gt.image.Bounds().Min.Y; y < gt.image.Bounds().Max.Y; y++ {
		for x := gt.image.Bounds().Min.X; x < gt.image.Bounds().Max.X; x++ {
			grayImage.Set(x, y, gt.image.At(x, y))
		}
	}
	return grayImage
}
