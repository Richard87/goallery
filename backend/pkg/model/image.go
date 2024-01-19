package model

import "image"

type Image struct {
	Id               string
	OriginalFilename string
	Path             string
	Mime             string
	Placeholder      string
	Rect             image.Point
	Size             int64
}
