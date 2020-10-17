package meme

import (
	"log"

	"github.com/fogleman/gg"
)

// CreateMeme Takes the format and creates an image, and adds the user-inputted caption
func CreateMeme(format, caption1, caption2 string, height, width int) {
	im, err := gg.LoadImage(format)
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("./fonts/Impact.ttf", 96); err != nil {
		panic(err)
	}

	h := float64(height)
	w := float64(width)
	margin := float64(50)
	
	dc.DrawRoundedRectangle(0, 0, h, w, 0)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringAnchored(caption1, w/2, margin, 0.5, 0.5)
	dc.DrawStringAnchored(caption2, w/2, h-margin, 0.5, 0.5)
	dc.Clip()
	dc.SavePNG("meme.png")
}
