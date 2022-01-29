package main

import (
	"errors"
	"os"

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

// isAtGitRoot returns true if the current directory is the
// root of a git repository.
func isAtGitRoot() bool {
	_, err := os.Stat(".git")
	return err == nil
}

// getGitHash returns the current git hash.
// If the current directory is not the root of a git repository,
// ErrNotAtGitRoot is returned.
func getGitHash() (string, error) {
	if !isAtGitRoot() {
		return "", ErrNotAtGitRoot
	}

	repo, err := git.PlainOpen(".")
	if err != nil {
		return "", err
	}

	ref, err := repo.Head()
	if err != nil {
		return "", err
	}

	return shortHash(ref.Hash()), nil
}
