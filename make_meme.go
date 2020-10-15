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
	
		//Label is a struct
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
func MakeMeme(imgs []ImageLayer, captions Caption[], bgProperty BgProperty) (*image.RGBA, error) {
	// Create the background layer
	bgImg := image.NewRGBA(image.Rect(0, 0, bgProperty.Width, bgProperty.Length))

	// Set the background color
	draw.Draw(bgImg, bgImg.Bounds(), &image.Uniform{bgProperty.BbColor}, image.ZP, draw.Src)

	// Put the image on the layer
	offset := image.Pt(img.XPos, img.YPos)

	// Combine the image
	draw.Draw(bgImg, img.Image.Bounds().Add(offset), img.image, image.ZP, draw.Over)

	// Call another function to handle the captions and combine it
	bgImg, err := addLabel(bgImg, labels)
	if err != nil {
		return nil, err
	}

	return bgImg, nil

