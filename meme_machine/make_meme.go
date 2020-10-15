package meme_machine

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	_ "image/png"

	m "github.com/arrafiv/bannergenerator"
)


func createMeme(format, caption1, caption2 string) {
	// Originally a link instead of format
	tempImg1, err := downloadMainImage(format)
	if err != nil {
		log.Println(err)
		return
	}

	imgs := []m.ImageLayer{
		m.ImageLayer{
			Image: tempImg1,
			XPos:  200,
			YPos:  -100,
		},
	}

	bg := m.BgProperty{
		Width:   500,
		Length:  380,
		BgColor: color.RGBA{227, 221, 221, 1},
	}

	//add label
	labels := []m.Label{
		m.Label{
			FontPath: "../../golang/freetype/testdata/",
			Size:     48,
			FontType: "BebasNeue-Regular.ttf",
			Color:    image.White,
			DPI:      72,
			Spacing:  1.5,
			Text:     caption1,
			XPos:     10,
			YPos:     0,
		},
		m.Label{
			FontPath: "../../golang/freetype/testdata/",
			Size:     48,
			FontType: "BebasNeue-Regular.ttf",
			Color:    image.White,
			DPI:      72,
			Spacing:  1.5,
			Text:     caption2,
			XPos:     10,
			YPos:     50,
		},
	}

	res, err := m.GenerateBanner(imgs, labels, bg)
	if err != nil {
		log.Printf("Error generating banner: %+v\n", err)
	}

	out, err := os.Create("./meme.jpg")
	if err != nil {
		log.Printf("Error creating image file: %+v\n", err)
		return
	}

	var opt jpeg.Options
	opt.Quality = 80

	err = jpeg.Encode(out, res, &opt)
	if err != nil {
		log.Printf("Error creating image file: %+v\n", err)
		return
	}

	log.Println("Banner Generated")
}

func downloadMainImage(url string) (image.Image, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	m, _, err := image.Decode(res.Body)
	if err != nil {
		return nil, err
	}

	return m, err
}
