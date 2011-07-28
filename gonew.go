package main

import "fmt"
import "flag"
import "bytes"
import "io/ioutil"
import "os"

func Usage() {
	fmt.Printf("Usage:\n\tgonew <type> <name>\n<type>\tpkg, clib or cmd for a package, library or a command\n<name>\tThe project name\n")
}

func main() {
	flag.Parse()
	
	if flag.NArg() != 2 {
		Usage()
		os.Exit(1)
	}
	
	if flag.Arg(0) != "cmd" && flag.Arg(0) != "clib" && flag.Arg(0) != "pkg" {
		Usage()
		os.Exit(1)
	}

	prjtype := flag.Arg(0)
	prjname := flag.Arg(1)
	
	dir := fmt.Sprintf("./%s", prjname)
		
	makefile := fmt.Sprintf(`include $(GOROOT)/src/Make.inc

TARG = %s
GOFILES = \
				%s.go

include $(GOROOT)/src/Make.%s
`, prjname, prjname, prjtype)

	gofile := fmt.Sprintf(`package main

import "fmt"

func main() {
	fmt.Printf("Hello World!")
}
`)

	readme := fmt.Sprintf(`%s
======

This is a go project.

## License

(The MIT License)

Copyright (c) 2011 Author &lt;e@mail.com&gt;

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
`, prjname)

	wintype := "exe"
	if prjtype != "cmd" {
		wintype = "lib"
	}
	
	makewin := fmt.Sprintf(`#!/bin/sh

export GOOS=windows
export GOARCH=386

8g -o _go_.8 %s.go
8l -o %s.%s _go_.8
`, prjname, prjname, wintype)

	gitignore := fmt.Sprintf("*.6\n*.8\n%s\n", prjname)
	
	gofilename := fmt.Sprintf("%s.go", prjname)
	os.Mkdir(dir, 0777)
	WriteFile(dir, "Makefile", makefile, 0666)
	WriteFile(dir, gofilename, gofile, 0666)
	WriteFile(dir, "README.md", readme, 0666)
	WriteFile(dir, "make_w32.sh", makewin, 0777)
	WriteFile(dir, ".gitignore", gitignore, 0666)

	fmt.Printf("Created a new project: %s\n", prjname)
}

func WriteFile(dir, name, content string, perm uint32) {
	file := bytes.NewBuffer([]byte{})
	
	fmt.Fprintf(file, content)
	
	outname := fmt.Sprintf("%s/%s", dir, name)
	fmt.Printf("Creating file %s\n", outname)
	
	ioutil.WriteFile(outname, file.Bytes(), perm)
}
