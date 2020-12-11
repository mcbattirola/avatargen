package main

import (
	"fmt"
	"os"

	"github.com/mcbattirola/avatargen"
)

func main() {
	str := os.Args[1]
	svg := avatargen.Generate(str)

	fileName := fmt.Sprintf("%s.svg", str)

	f, _ := os.Create(fileName)
	defer f.Close()

	f.WriteString(svg)
}
