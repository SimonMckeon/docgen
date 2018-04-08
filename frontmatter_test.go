package main

import "testing"

func TestHasFrontMatter(t *testing.T) {
	withFrontMatter := []byte(`
    ---
    title: testTitle
    ---

    # h1
    `)

	withoutFrontMatter := []byte("# h1")

	malformedFrontMatter := []byte(`
    ---
    title: testTitle
    --

    # h1
    `)

	result, _ := HasFrontMatter(withFrontMatter)
	if result == false {
		t.Error("File with front matter reporting no front matter")
	}

	result, _ = HasFrontMatter(withoutFrontMatter)
	if result == true {
		t.Error("File with no front matter reporting front matter")
	}

	result, _ = HasFrontMatter(malformedFrontMatter)
	if result == true {
		t.Error("File with malformed frontmatter reporting valid front matter")
	}
}
