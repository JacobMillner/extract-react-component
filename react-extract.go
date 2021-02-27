package main

import (
	"fmt"
	"os"
)

func main() {

	// check args
	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "ENTRY_FILE_PATH", "COPY_DIR")
		return
	}

	entry_file := os.Args[1]
	copy_dir := os.Args[2]

	fmt.Println(entry_file)
	fmt.Println(copy_dir)
}
