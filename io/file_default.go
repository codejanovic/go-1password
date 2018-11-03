package io

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	throw "github.com/codejanovic/go-1password/throw"
)

// FileByPath implementing File interface
type FileByPath struct {
	path string
}

// NewFileByAbsolutePath creates a new file by path
func NewFileByAbsolutePath(path string) File {
	file := new(FileByPath)
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		throw.Throw(fmt.Errorf("Unable to create absolute path for %s", path), "Maybe the path is incorrect, or the opvault does not exist?")
	}
	file.path = absolutePath
	return file
}

// NewFileByPath creates a new file by path
func NewFileByPath(path string) File {
	file := new(FileByPath)
	file.path = path
	return file
}

// Path gets the path to the file
func (f *FileByPath) Path() string {
	return f.path
}

// IsEqualTo compares the file
func (f *FileByPath) IsEqualTo(file File) bool {
	return file.Path() == f.path
}

// Exists whether or not the file exists
func (f *FileByPath) Exists() bool {
	_, err := os.Stat(f.path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Create empty file if it does not exist
func (f *FileByPath) Create() error {
	os.MkdirAll(path.Dir(f.path), os.ModePerm)
	file, err := os.Create(f.path)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

// AsBytes from file
func (f *FileByPath) AsBytes() ([]byte, error) {
	return ioutil.ReadFile(f.path)
}

func (f *FileByPath) Write(data []byte) error {
	return ioutil.WriteFile(f.path, data, 0644)
}
