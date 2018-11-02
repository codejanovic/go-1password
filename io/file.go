package io

// File interface
type File interface {
	Create() error
	Exists() bool
	Path() string
	IsEqualTo(file File) bool
	AsBytes() ([]byte, error)
	Write(data []byte) error
}
