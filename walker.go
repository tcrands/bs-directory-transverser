package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Walker struct {
	p  string
	fn filepath.WalkFunc
}

func (w *Walker) Walk() {
	filepath.Walk(w.p, w.fn)
}

func (w *Walker) WriteToFile(buffer []string, file []byte) {
	fmt.Println(buffer)
}

/////////////////////
// Walking Functions
/////////////////////

func NewWalker(p string, fn filepath.WalkFunc) *Walker {
	return &Walker{
		p:  p,
		fn: fn,
	}
}

func walkRootPath(extention string) filepath.WalkFunc {
	fmt.Println(extention)
	return func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, extention) {
			file, _ := ioutil.ReadFile(path)

			subPaths := getParams(`<!-- SCRIPT-INCLUDE uri="(?P<path>.*?)" -->`, string(file))

			var wg sync.WaitGroup
			wg.Add(len(subPaths["path"]))

			buffer := []string{}

			for i := range subPaths["path"] {
				go func(i *string) {
					dereferencedI := *i
					walker := NewWalker("./"+dereferencedI, walkSubPath(path, &buffer))
					walker.Walk()
					wg.Done()
				}(&subPaths["path"][i])
			}

			wg.Wait()

			splitPoint := getSplitPoint(`<!-- SCRIPT-INCLUDE uri=".*?" -->`, string(file))
			processedFile := processUpdatedFile(buffer, file, splitPoint)

			ioutil.WriteFile(path, []byte(processedFile), 0644)

		}
		return nil
	}
}

func walkSubPath(origPath string, buffer *[]string) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, ".brs") {
			tag := "<script type=\"text/brightscript\" uri=\"pkg:/" + path + "\" />"
			*buffer = append(*buffer, tag)
		}
		return nil
	}
}

func processUpdatedFile(buffer []string, file []byte, splitPoint string) string {
	newString := "\n"
	for i := range buffer {
		newString = newString + buffer[i] + "\n"
	}

	splitSlice := strings.SplitAfter(string(file), splitPoint)
	newFile := splitSlice[0] + newString + splitSlice[1]

	return newFile
}
