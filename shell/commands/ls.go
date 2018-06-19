package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

// LS is
func LS(root string) int {
	if root == "" {
		fmt.Println("Directory can't be null")
		return 0
	}

	if len(root) == 0 || root == "." {
		curr, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
			return 0
		} else {
			root = curr
		}
	}

	info, err := os.Lstat(root)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	if !info.IsDir() {
		fmt.Println(root)
		return 1
	}

	names, dirOpenError := readDirNames(root)
	if dirOpenError != nil {
		fmt.Println(dirOpenError)
		return 0
	}

	var count uint64
	fmt.Println(".")

	for _, name := range names {
		fileInfo, err := os.Lstat(fmt.Sprintf("%s\\%s", root, name))
		if err != nil {
			fmt.Println(err)
			continue
		}

		if fileInfo.IsDir() {
			fmt.Println(name + "/")
		} else {
			fmt.Println(name)
		}

		count++
	}

	fmt.Printf("--> total: %d\n", count)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return 0
	}

	return 1
}

// readDirNames reads the directory named by dirname and returns
// a sorted list of directory entries.
func readDirNames(dirname string) ([]string, error) {
	folder, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}

	names, err := folder.Readdirnames(-1)
	folder.Close()

	if err != nil {
		return nil, err
	}

	sort.Strings(names)

	return names, nil
}
