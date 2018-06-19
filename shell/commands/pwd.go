package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// PWD ...
func PWD(root string) {
	if len(root) == 0 {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(dir)
	} else {
		fmt.Println(root)
	}
}
