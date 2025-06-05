package localFonts

import (
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

type Fonts struct {
	Face  font.Face
	Face2 font.Face
}

var AllFonts Fonts

func LoadFonts() error {
	ttf, err := opentype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
		return err
	}
	face2, err := opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    60,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return err
	}
	fontData, err := os.ReadFile("localFonts/assets/font.ttf")
	if err != nil {
		log.Fatal(err)
	}
	ttfFont, err := truetype.Parse(fontData)
	if err != nil {
		log.Fatal(err)
		return err
	}
	face := truetype.NewFace(ttfFont, &truetype.Options{
		Size:    130,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	AllFonts = Fonts{
		Face:  face,
		Face2: face2,
	}
	return nil
}
