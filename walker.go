package main

import (
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

/*
	@function Walk - runs the inbuild file walker
	@return {error}
*/
func (w *Walker) Walk() error {
	return filepath.Walk(w.p, w.fn)
}

/////////////////////
// Walking Functions
/////////////////////

/*
	@function NewWalker - creates a new walker struct
	@param {String} p - the directory path to transverce
	@param {filepath.WalkFunc} fn - the function to run on each directory pass
	@return {Walker} - a new instance of a Walker
*/
func NewWalker(p string, fn filepath.WalkFunc) *Walker {
	return &Walker{
		p:  p,
		fn: fn,
	}
}

/*
	@function walkRootPath - transverses the root path
	@param {String} extention - the file extention to be isolated
	@return {filepath.WalkFunc} - a wrapper for custom logic on the directory transversal
*/
func walkRootPath(extention string) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, extention) {
			file, _ := ioutil.ReadFile(path)
			subPaths := getParams(`<!-- SCRIPT-INCLUDE uri="(?P<path>.*?)" -->`, string(file))

			if len(subPaths) != 0 {
				buffer := processSubPaths(subPaths)
				splitPoint := getSplitPoint(`<!-- SCRIPT-INCLUDE uri=".*?" -->`, string(file))
				processedFile := processUpdatedFile(buffer, file, splitPoint)
				ioutil.WriteFile(path, []byte(processedFile), 0644)
			}
		}
		return nil
	}
}

/*
	@function walkSubPath - transverses the sub path
	@param {Write Channel of Strings} queue - to hold each generatered script line
	@return {filepath.WalkFunc} - a wrapper for custom logic on the directory transversal
*/
func walkSubPath(queue chan<- string) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, ".brs") {
			tag := "<script type=\"text/brightscript\" uri=\"pkg:/" + path + "\" />"
			queue <- tag
		}
		return nil
	}
}

/*
	@function processSubPaths - transerves all sub paths asynchrously and stores the result
	@param {map[string][]string} subPaths - a map of all sub path to be transversed
	@return {[]string} - a slice containing everything from the queue channel
*/
func processSubPaths(subPaths map[string][]string) []string {
	var wg sync.WaitGroup
	buffer := []string{}
	queue := make(chan string, 1)

	for _, subPath := range subPaths["path"] {
		wg.Add(1)

		/*
			@function anon - transerves all sub paths asynchrously
			@param {String} subPath - the sub path to be transversed
			@param {chan<- string} queue - the channel to write too
		*/
		go func(subPath string, queue chan<- string) {
			walker := NewWalker("./"+subPath, walkSubPath(queue))
			if err := walker.Walk(); err == nil {
				wg.Done()
			}
		}(subPath, queue)
	}

	/*
		@function anon - monitors the channel for new items
		@param {*[]string} buffer - the slice to hold the new script tags
	*/
	go func(buffer *[]string) {
		for t := range queue {
			*buffer = append(*buffer, t)
		}
	}(&buffer)

	wg.Wait()

	return buffer

}

/*
	@function processUpdatedFile - writes the new script tags to the original file
	@param {[]string} buffer - a map of all sub path to be transversed
	@param {[]byte} file - the file to be written too
	@param {String} splitPoint - the position in whcih to insert the new script tags
	@return {String} - the new file containing all the script tags
*/
func processUpdatedFile(buffer []string, file []byte, splitPoint string) string {
	newString := "\n"
	for i := range buffer {
		newString = newString + buffer[i] + "\n"
	}

	splitSlice := strings.SplitAfter(string(file), splitPoint)
	newFile := splitSlice[0] + newString + splitSlice[1]

	return newFile
}
