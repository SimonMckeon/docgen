package main

import (
	"gopkg.in/russross/blackfriday.v2"
)

func CompileMarkdown(source []byte) ([]byte, error) {
	// If a source file is empty, don't continue processing
	// TODO: How to handle early returns?
	if len(source) == 0 {
		return source, nil
	}
	output := blackfriday.Run(source)
	return output, nil
}
