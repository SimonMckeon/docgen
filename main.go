package main

import (
	"fmt"
)

func main() {
	source, err := ReadSourceFile("testmarkdown")
	if err != nil {
		panic(err)
	}

	fm, content, err := SeparateFrontMatter(source)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fm[:]))
	fmt.Println(string(content[:]))

	compiled, err := CompileMarkdown(content)
	if err != nil {
		panic(err)
	}

	err = WriteCompiledFile("testcompiled", compiled)
	if err != nil {
		panic(err)
	}
}
