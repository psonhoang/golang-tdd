package blogposts

import (
	"errors"
	"io/fs"
)

type StubFailingFS struct {}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i alwatys fail :(")
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err	//todo: needs clarification, should we totally fail if one file fails, or just ignore?
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}