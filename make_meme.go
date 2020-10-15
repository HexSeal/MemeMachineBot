package main

import (
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"log"

	"github.com/golang/freetype"
)

// Helped by https://medium.com/@arrafiv/basic-image-processing-with-go-combining-images-and-texts-8510d9214e55

type (
		//ImageLayer is a struct
		ImageLayer struct {
			Image image.Image
			XPos  int
			YPos  int
		}
	
		//caption is a struct
		Caption struct {
			Text     string
			FontPath string
			FontType string
			Size     float64
			Color    image.Image
			DPI      float64
			Spacing  float64
			XPos     int
			YPos     int
		}
	
		//BgProperty is background property struct
		BgProperty struct {
			Width   int
			Length  int
			BgColor color.Color
		}
)

// Combines the meme format and caption
func MakeMeme(img []ImageLayer, captions []Caption, bgProperty BgProperty) (*image.RGBA, error) {
	// Create the background layer
	bgImg := image.NewRGBA(image.Rect(0, 0, bgProperty.Width, bgProperty.Length))

	// Set the background color
	draw.Draw(bgImg, bgImg.Bounds(), &image.Uniform{bgProperty.BbColor}, image.ZP, draw.Src)

	// Put the image on the layer
	offset := image.Pt(img.XPos, img.YPos)

	// Combine the image
	draw.Draw(bgImg, img.Image.Bounds().Add(offset), img.image, image.ZP, draw.Over)

	// Call another function to handle the captions and combine it
	bgImg, err := addCaption(bgImg, captions)
	if err != nil {
		return nil, err
	}

	return bgImg, nil
}

func addCaption(img *image.RGBA, captions []Caption) (*image.RGBA, error) {
	// Initialize the context
	c := freetype.NewContext()
	for _, caption := range captions {
		// Read the font data
		fontBytes, err := ioutil.ReadFile(caption.FontPath + caption.FontType)
		if err != nil {
			return nil, err
		}
		f, err := freetype.ParseFont(fontBytes)
		if err != nil {
			return nil, err
		}

		// Caption Settings
		c.SetDPI(caption.DPI)
		c.SetFont(f)
		c.SetFontSize(caption.Size)
		c.SetClip(img.Bounds())
		c.SetDst(img)
		c.SetSrc(caption.Color)

		// Positioning the caption
		pt := freetype.Pt(caption.XPos, caption.YPos+int(c.PointToFixed(caption.Size)>>6))

		// Put the caption on the image
		_, err = c.DrawString(caption.Text, pt)
		if err != nil {
			log.Println(err)
			return img, nil
		}
		pt.Y += c.PointToFixed(caption.Size * caption.Spacing)
	}
	return img, nil
}

