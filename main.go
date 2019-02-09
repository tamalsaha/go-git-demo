package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"k8s.io/client-go/util/homedir"
	"log"
	"os"
	"path/filepath"
)


func main() {
	dir := filepath.Join(homedir.HomeDir(), ".rolodex", "tamalsaha", "nats-demo")

	r, err := git.PlainClone(dir, true, &git.CloneOptions{
		URL:      "https://github.com/tamalsaha/nats-demo.git",
		Progress: os.Stdout,
	})
	if err != nil && err == git.ErrRepositoryAlreadyExists {
		r, err = git.PlainOpen(dir)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// ... retrieves the branch pointed by HEAD
	ref, err := r.Head()
	if err != nil {
		log.Fatalln(err)
	}

	// ... retrieves the commit history
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		log.Fatalln(err)
	}

	// ... just iterates over the commits, printing it
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c.Author.Name, c.Author.Email)
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
}
