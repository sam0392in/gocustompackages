# File Reader
This package provides reading and writing to the file.

## Usage

### Read the file
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
}

```