package entity

import "github.com/nfnt/resize"

type Scaling struct {
	Path      string
	Width     uint
	Height    uint
	InterFunc resize.InterpolationFunction
	OutName   string
	Ext       string
}

func (s *Scaling) MimeType() string {
	switch s.Ext {
	case "jpg":
		return "image/jpeg"
	case "jpeg":
		return "image/jpeg"
	case "png":
		return "image/png"
	case "gif":
		return "image/gif"
	}

	return "text/plain"
}
