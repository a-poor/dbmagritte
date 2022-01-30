package main

import (
	"errors"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

var (
	ErrNotAtGitRoot = errors.New("not at git root")
)

// shortHash returns the first 7 characters of a git hash.
func shortHash(hash plumbing.Hash) string {
	return hash.String()[:7]
}

// isAtGitRoot returns true if `dirPath` is the
// root of a git repository.
func isAtGitRoot(dirPath string) bool {
	p := path.Join(dirPath, ".git")
	_, err := os.Stat(p)
	return err == nil
}

// getGitHash returns the current git hash.
// If `dirPath` is not the root of a git repository,
// ErrNotAtGitRoot is returned.
func getGitHash(dirPath string) (string, error) {
	if !isAtGitRoot(dirPath) {
		return "", ErrNotAtGitRoot
	}

	repo, err := git.PlainOpen(dirPath)
	if err != nil {
		return "", err
	}

	ref, err := repo.Head()
	if err != nil {
		return "", err
	}

	return shortHash(ref.Hash()), nil
}
