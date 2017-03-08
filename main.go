package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func walkpath(path string, f os.FileInfo, err error) error {
	// MODIFY THIS LINE TO PRINT OUT WHAT YOU WANT!
	if strings.Contains(path, ".") {
		fmt.Printf("<script type=\"text/brightscript\" uri=\"pkg:/source/sky/roku/client-lib-bs-ott/" + path + "\" />\n")
	}
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0) // 1st argument is the directory location
	filepath.Walk(root, walkpath)
}
