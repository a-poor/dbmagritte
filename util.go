package main

import (
	"errors"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

var (
	ErrNotAtProjRoot = errors.New("not at git root")
)

// shortHash returns the first 7 characters of a git hash.
func shortHash(hash plumbing.Hash) string {
	return hash.String()[:7]
}

// isAtGitRoot returns true if `projRootPath` is the
// root of a git repository.
func isAtGitRoot(projRootPath string) bool {
	p := path.Join(projRootPath, ".git")
	return doesDirExist(p)
}

// getGitHash returns the current git hash.
// If `dirPath` is not the root of a git repository,
// ErrNotAtProjRoot is returned.
func getGitHash(dirPath string) (string, error) {
	if !isAtGitRoot(dirPath) {
		return "", ErrNotAtProjRoot
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

type FilePathState int

const (
	PathDoesNotExist FilePathState = iota
	PathIsFile
	PathIsDir
)

// whatsThatPath simplifies the logic of checking if a
// path exists and is a file or directory by converting
// the error to a FilePathState. An error is returned if
// a `*PathError` is encountered other than `ErrNotExist`.
func whatsThatPath(path string) (FilePathState, error) {
	// Get the path info.
	info, err := os.Stat(path)

	// Check if an unknown error occurred.
	if err != nil && !os.IsNotExist(err) {
		return PathDoesNotExist, err
	}
	// Otherwise, error means the path does not exist.
	if err != nil {
		return PathDoesNotExist, nil
	}
	// Otherwise, no error...the path exists.

	// Is the path a directory?
	if info.IsDir() {
		return PathIsDir, nil
	}
	// Otherwise, the path is a file.
	return PathIsFile, nil
}

// doesPathExist returns true if `path` exists and
// is an (accessable) file or directory.
func doesPathExist(path string) bool {
	state, err := whatsThatPath(path)
	return err == nil && state != PathDoesNotExist
}

// doesFileExist returns true if `path` exists and
// is an (accessable) file.
func doesFileExist(path string) bool {
	state, err := whatsThatPath(path)
	return err == nil && state == PathIsFile
}

// doesDirExist returns true if `path` exists and
// is an (accessable) directory.
func doesDirExist(path string) bool {
	state, err := whatsThatPath(path)
	return err == nil && state == PathIsDir
}
