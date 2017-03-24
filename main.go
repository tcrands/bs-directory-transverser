package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"io/ioutil"
	"regexp"
)

func main() {
	flag.Parse()
	root := flag.Arg(0) // 1st argument is the directory location
	filepath.Walk(root, walkRootPath)
}

func walkRootPath(path string, f os.FileInfo, err error) error {
	// MODIFY THIS LINE TO PRINT OUT WHAT YOU WANT!
	if strings.Contains(path, ".xml") {
		file, _ := ioutil.ReadFile(path)

		subPath := getParams(`<!-- SCRIPT-INCLUDE uri="(?P<path>.*?)" -->`, string(file))

		buffer := []string{}
		filepath.Walk("./" + subPath["path"], walkSubPath(path, &buffer))
		
		writeToFile(buffer, file)

	}
	return nil
}


// REMEMBER! Good example of wrapping function for scoping
func walkSubPath(origPath string, buffer *[]string) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, ".brs") {
			tag := "<script type=\"text/brightscript\" uri=\"pkg:/" + path + "\" />"
			*buffer = append(*buffer, tag)
		}
		return nil
	}
}

func writeToFile(buffer []string, file []byte) {
	fmt.Println("WRITE BUFFER TO FILE")
}

func getParams(regEx, url string) (paramsMap map[string]string) {

    compRegEx := regexp.MustCompile(regEx)
    match := compRegEx.FindStringSubmatch(url)

    paramsMap = make(map[string]string)
    for i, name := range compRegEx.SubexpNames() {
        if i > 0 && i <= len(match) {
            paramsMap[name] = match[i]
        }
    }
    return nil
}
