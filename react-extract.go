package main

import (
	"fmt"
	"os"
  "io/ioutil"
	"regexp"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

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

	// read the entry file
	dat, err := ioutil.ReadFile(os.Args[1])
	check(err)
	fmt.Println("Data loaded: ")
	fmt.Print(string(dat))

	// regex to grab file paths from imports
	match, _ := regexp.MatchString(`^import\s+(.* from\s+)?['|"]([.\/|~\/].*)[['|"]|[';|";]]$`, string(dat))
	fmt.Println("Regex Match: ")
  fmt.Println(match)

}
