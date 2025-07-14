package dir

import (
	"fmt"
	"os"
)

func GenerateDirTree() {
	path := os.Args[2]
	graphOutput := os.Args[3]

	_, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Error checking output path: %v\n", err)
		return
	}

	WalkPath(path, graphOutput)
}
