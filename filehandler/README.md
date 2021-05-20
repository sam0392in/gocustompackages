# File Handler
This package provides reading and writing to the file.

## Usage

### Read the file
```
package main

import (
        "fmt"

        "github.com/sam0392in/gocustompackages/filehandler"
)

func main() {

    file := "test.txt"
    lines := filehandler.Readlines(file)

	for _, line := range lines {
		fmt.Println("\n-> " + line)
    }
}

```

### Write to file
```
package main

import (
	"github.com/sam0392in/gocustompackages/filehandler"
)

func main() {

	// file path
	file := "helper/priority-class-template.yaml"

	// Read the file contents
	lines := filehandler.Readlines(file)

	// Update the lines
	lines[3] = "  name: l1"
	lines[4] = "value: 13"

	// Use Write functionality to write updated lines to file
	writeMode := "overwrite" // options: ["overwrite", "append"]
	filehandler.Writelines(file, lines, writeMode)

}

```