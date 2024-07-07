package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title	string
	Description	string
	Body	string
	Tags	[]string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title: readMetaLine(titleSeparator), 
		Description: readMetaLine(descriptionSeparator), 
		Tags: strings.Split(readMetaLine(tagsSeparator), ", "),
		Body: readBody(scanner)}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()	// skip the body separator

	buf := bytes.Buffer{}
	for scanner.Scan() {	// scans until EOF
		fmt.Fprintln(&buf, scanner.Text())	// we use println() because Scan() removes the \n
	}
	return strings.TrimSuffix(buf.String(), "\n")	// then trim the final \n because thats what we added
}