package flysystem

import "github.com/edwin-luijten/go_flysystem/adapter"

// Flysystem ...
type Flysystem struct {
	adapter adapter.Adapter
}

// New creates a new instance with a given adapter
func New(adapter adapter.Adapter) adapter.Adapter {
	return &Flysystem{
		adapter: adapter,
	}
}

// Write a new file
func (f *Flysystem) Write(path string, contents []byte) error {
	return f.adapter.Write(path, contents)
}

// Update a file
func (f *Flysystem) Update(path string, contents []byte) error {
	return f.adapter.Update(path, contents)
}

// Read a file
func (f *Flysystem) Read(path string) ([]byte, error) {
	return f.adapter.Read(path)
}

// Rename a file
func (f *Flysystem) Rename(path string, newPath string) error {
	return f.adapter.Rename(path, newPath)
}

// Copy a file
func (f *Flysystem) Copy(path string, newPath string) error {
	return f.adapter.Copy(path, newPath)
}

// Delete a file
func (f *Flysystem) Delete(path string) error {
	return f.adapter.Delete(path)
}

// CreateDir creates a directory
func (f *Flysystem) CreateDir(dir string) error {
	return f.adapter.CreateDir(dir)
}

// DeleteDir deletes a directory
func (f *Flysystem) DeleteDir(dir string) error {
	return f.adapter.DeleteDir(dir)
}

// SetVisibility sets a file or directory to public or private
func (f *Flysystem) SetVisibility(path string, visibility string) error {
	return f.adapter.SetVisibility(path, visibility)
}
