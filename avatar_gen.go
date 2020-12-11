// Package avatargen generats an 300x300 px avatar based on a input string.
// The avatar is aways unique
package avatargen

import (
	"crypto/md5"
	"fmt"
	"io"
)

// Generate generates a svg avatar from an input, with width and height
func Generate(input string) string {
	h := md5.New()
	io.WriteString(h, input)
	hash := h.Sum(nil)

	// first 3 bytes make the background color
	backgroundColorBytes := hash[:4]
	backgroundColor := makeRGBColor(backgroundColorBytes)
	invertedBackgroundColor := makeRGBColor([]byte{backgroundColor[2], backgroundColor[1], backgroundColor[0]})

	// last 3 bytes make the element color
	elementColorBytes := hash[len(hash)-3 : len(hash)]
	elementColor := makeRGBColor(elementColorBytes)
	invertedElementColor := makeRGBColor([]byte{elementColorBytes[2], elementColorBytes[1], elementColorBytes[0]})

	paths := []string{
		generatePath(int(hash[4]), int(hash[5]), 2),
		generatePath(int(hash[5]), int(hash[7]), 1),
		generatePath(int(hash[8]), int(hash[8]), 0),
	}

	return hydrateSVG(paths[0], paths[1], paths[2], backgroundColor, elementColor, invertedElementColor, invertedBackgroundColor)
}

func generatePath(curveVal, posVal, index int) string {
	bigC := 300 - curveVal

	return fmt.Sprintf("m 150 %d Q %d %d %d 150 Q %d %d 150 %d Q %d %d %d 150 Q %d %d 150 %d z",
		(100 + posVal + (200 * index)),
		bigC,
		curveVal,
		(200 - posVal - (200 * index)),
		bigC,
		bigC,
		(200 - posVal - (200 * index)),
		curveVal,
		bigC,
		(100 + posVal + (200 * index)),
		curveVal, curveVal, (100 + posVal + (200 * index)))
}

func hydrateSVG(p0, p1, p2 string, backgroundColor string, c0, c1, c2 string) string {
	return fmt.Sprintf("<svg width=\"300\" height=\"300\" viewBox=\"0 0 300 300\" "+
		"xmlns=\"http://www.w3.org/2000/svg\"><rect id=\"bg\" width=\"300\" height=\"300\" "+
		"fill=\"rgb(%s)\" /><path d=\"%s\" "+
		"fill=\"rgb(%s)\" /><path d=\"%s\" "+
		"fill=\"rgb(%s)\" /><path d=\"%s\" "+
		"fill=\"rgb(%s)\" /></svg>", backgroundColor, p0, c0, p1, c1, p2, c2)
}

func makeRGBColor(bytes []byte) string {
	return fmt.Sprintf("%d,%d,%d", bytes[0], bytes[1], bytes[2])
}
