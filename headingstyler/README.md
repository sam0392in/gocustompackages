# Heading Styler
This package provide stylish colored fonts for texts and headings.

## Usage

```
package main

import (
        "fmt"

        headingstyler "github.com/sam0392in/gocustompackages/headingstyler"
)

func main() {
        data := "samarth"
        color := "Blue"
        // Colors: Blue, Yellow, Green, Red, Cyan ; Default: White

        font := "3d"
        // Font: for simple font do font:= ""

        // Call the function
        renderdata := headingstyler.Styler(&data, &color, &font)
        fmt.Println(renderdata)
}


```