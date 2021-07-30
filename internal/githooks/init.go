package githooks

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func resolveGitRoot(path string) string {
	f, err := os.Lstat(filepath.Join(path, ".git"))
	if err != nil {
		if os.IsNotExist(err) {
			return resolveGitRoot(filepath.Join(path, ".."))
		}
		panic(err)
	}
	if !f.IsDir() {
		panic(fmt.Errorf(".git must be a directory"))
	}
	return path
}

var hooks = map[string][]byte{
	"pre-commit": []byte(`#!/bin/sh

git add $(tools fmt)
git config --get user.email | tools hook check-email
`),
	"commit-msg": []byte(`#!/bin/sh

tools hook run $(basename "$0")
`),
}

func getGithooks(root string) ([]string, error) {
	githooks := make([]string, 0)

	files, err := ioutil.ReadDir(path.Join(root, ".git/hooks"))
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		if filepath.Ext(f.Name()) == ".sample" {
			githooks = append(githooks, strings.Split(f.Name(), ".")[0])
		}
	}

	return githooks, nil
}

func Init() {
	cwd, _ := os.Getwd()
	root := resolveGitRoot(cwd)
	githooks, _ := getGithooks(root)

	for _, githook := range githooks {
		ioutil.WriteFile(path.Join(root, ".git/hooks", githook), []byte(`#!/bin/sh
tools hook run $(basename "$0")
`), os.ModePerm)
	}
}
