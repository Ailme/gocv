package gocv

import "image"

// Return crop image
func (img *Mat) Crop(rect image.Rectangle) Mat {
	croppedMat := img.Region(rect)
	return croppedMat.Clone()
}