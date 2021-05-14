# Heading Styler
This package provide stylish colored fonts for texts and headings.

## Usage

```
package main

import (
	"fmt"

	"go/custompkgs/headingstyler"
)

func main() {
	data := "samarth"
	color := "Blue"
	// Colors: Blue, Yellow, Green, Red, Cyan ; Default: White
	font := "3d"
	// Font: for simple font do font:= ""
	renderdata := headingstyler.Styler(&data, &color, &font)
	fmt.Println(renderdata)
}

```