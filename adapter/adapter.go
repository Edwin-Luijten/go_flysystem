package adapter

import (
	"fmt"
	"os"
	"strings"
)

// Adapter ...
type Adapter interface {
	Write(path string, contents []byte) error
	Update(path string, contents []byte) error
	Read(path string) ([]byte, error)
	Rename(path string, newPath string) error
	Copy(path string, newPath string) error
	Delete(path string) error
	CreateDir(dir string) error
	DeleteDir(dir string) error
	SetVisibility(path string, visibility string) error
}

// BaseAdapter ...
type BaseAdapter struct {
	pathPrefix *string
}

// SetPathPrefix sets the path prefix
func (a *BaseAdapter) SetPathPrefix(prefix string) {
	if prefix == "" {
		a.pathPrefix = nil
	}

	p := fmt.Sprintf("%s%s", prefix, string(os.PathSeparator))
	a.pathPrefix = &p
}

// ApplyPathPrefix applies the path prefix
func (a *BaseAdapter) ApplyPathPrefix(path string) string {
	return fmt.Sprintf("%s%s", *a.pathPrefix, strings.TrimPrefix(path, string(os.PathSeparator)))
}
