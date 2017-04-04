package main

import "flag"

func main() {
	flag.Parse()
	root := flag.Arg(0)      // 1st argument is the directory location
	extention := flag.Arg(1) // 2nd argument is the file extention
	walker := NewWalker(root, walkRootPath(extention))
	walker.Walk()
}
