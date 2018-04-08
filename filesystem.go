package main

import (
	"io/ioutil"
	"os"
)

const sourceDir = "./content/"
const outputDir = "./public/"

func ReadSourceFile(name string) ([]byte, error) {
	data, err := ioutil.ReadFile(sourceDir + name + ".md")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteCompiledFile(name string, data []byte) error {
	// Create output directory if it doesn't exist
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			return err
		}
	}

	err := ioutil.WriteFile(outputDir+name+".html", data, 0644)
	if err != nil {
		return err
	}
	return nil
}
