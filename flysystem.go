package flysystem

import (
	"github.com/edwin-luijten/go_flysystem/adapter"
	"sync"
)

// Flysystem ...
type Flysystem struct {
	adapters []adapter.Adapter
}

// New creates a new instance with given adapters
func New(adapters ...adapter.Adapter) adapter.Adapter {
	return &Flysystem{
		adapters: adapters,
	}
}

// Write a new file
func (f *Flysystem) Write(path string, contents []byte) error {
	var wg sync.WaitGroup
	wg.Add(len(f.adapters))

	errors := make(chan error, 1)

	for _, a := range f.adapters {
		go func(a adapter.Adapter) {
			defer wg.Done()

			err := a.Write(path, contents)
			if err != nil {
				errors <- err
			}
		}(a)
	}

	wg.Wait()
	close(errors)

	select {
	case err := <-errors:
		return err
	default:
		return nil
	}
}

// Update a file
func (f *Flysystem) Update(path string, contents []byte) error {
	var wg sync.WaitGroup
	wg.Add(len(f.adapters))

	errors := make(chan error, 1)

	for _, a := range f.adapters {
		go func(a adapter.Adapter) {
			defer wg.Done()

			err := a.Update(path, contents)
			if err != nil {
				errors <- err
			}
		}(a)
	}

	wg.Wait()
	close(errors)

	select {
	case err := <-errors:
		return err
	default:
		return nil
	}
}

// Read a file
func (f *Flysystem) Read(path string) ([]byte, error) {
	var wg sync.WaitGroup
	wg.Add(len(f.adapters))

	errors := make(chan error, 1)
	contents := make(chan []byte)

	for _, a := range f.adapters {
		go func(a adapter.Adapter) {
			bytes, err := a.Read(path)
			if err != nil {
				errors <- err
			}

			wg.Done()
			contents <- bytes
		}(a)
	}

	wg.Wait()

	select {
	case err := <-errors:
		close(errors)
		return nil, err
	case content := <-contents:
		close(contents)
		return content, nil
	}
}

// Rename a file
func (f *Flysystem) Rename(path string, newPath string) error {
	panic("to implement")
}

// Copy a file
func (f *Flysystem) Copy(path string, newPath string) error {
	panic("to implement")
}

// Delete a file
func (f *Flysystem) Delete(path string) error {
	panic("to implement")
}

// CreateDir creates a directory
func (f *Flysystem) CreateDir(dir string) error {
	panic("to implement")
}

// DeleteDir deletes a directory
func (f *Flysystem) DeleteDir(dir string) error {
	panic("to implement")
}

// SetVisibility sets a file or directory to public or private
func (f *Flysystem) SetVisibility(path string, visibility string) error {
	panic("to implement")
}
