package main

import (
	"bytes"
	"errors"
	"fmt"

	"gopkg.in/yaml.v2"
)

var YAMLDelimiter []byte = []byte("---")

func HasFrontMatter(source []byte) (bool, error) {
	source = bytes.TrimSpace(source)
	if !bytes.HasPrefix(source, YAMLDelimiter) {
		return false, nil
	}
	source = bytes.TrimPrefix(source, YAMLDelimiter)
	elements := bytes.SplitN(source, YAMLDelimiter, 2)
	if len(elements) != 2 {
		return false, errors.New("Frontmatter is malformed and cannot be processed")
	}
	out := make(map[string]interface{})
	err := yaml.Unmarshal(elements[0], out)
	if err != nil {
		return false, err
	}
	return true, nil
}

func SeparateFrontMatter(source []byte) ([]byte, []byte, error) {
	// Nothing to process
	if len(source) == 0 {
		return nil, source, nil
	}

	b, err := HasFrontMatter(source)
	if err != nil {
		return nil, nil, err
	}

	// return source if no front matter
	if !b {
		fmt.Println("No frontmatter, returning original source")
		return nil, source, nil
	}

	source = bytes.TrimSpace(source)
	source = bytes.TrimPrefix(source, YAMLDelimiter)

	elements := bytes.SplitN(source, YAMLDelimiter, 2)
	fm := bytes.TrimSpace(elements[0])
	rest := bytes.TrimSpace(elements[1])

	return fm, rest, nil
}
