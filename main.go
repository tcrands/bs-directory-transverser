package main

import "flag"

func main() {
	flag.Parse()
	root := "./"        // 1st argument is the directory location
	extention := ".xml" // 2nd argument is the file extention
	walker := NewWalker(root, walkRootPath(extention))
	walker.Walk()
}
