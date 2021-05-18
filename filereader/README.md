# File Reader
This package provides reading the lines of a given file and returns an array.

## Usage

```
package main

import (
        "fmt"

        "github.com/sam0392in/gocustompackages/filereader"
)

func main() {
    
    file := "test.txt"
    lines := filereader.Readlines(file)

	for _, line := range lines {
		fmt.Println("\n-> " + line)

}

```