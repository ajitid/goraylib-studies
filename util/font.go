package util

import (
	"log"

	"github.com/adrg/sysfont"
)

func GetSystemFontPath() string {
	finder := sysfont.NewFinder(&sysfont.FinderOpts{
		Extensions: []string{".ttf"},
	})
	font := finder.Match("Segoe UI")
	if font == nil {
		log.Fatal("system font not found")
	}
	return font.Filename
}
