package flysystem

import (
	"github.com/edwin-luijten/go_flysystem/adapter"
	"golang.org/x/sync/errgroup"
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
	var g errgroup.Group

	for _, a := range f.adapters {
		a := a
		g.Go(func() error {
			err := a.Write(path, contents)

			return err
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

// Update a file
func (f *Flysystem) Update(path string, contents []byte) error {
	var g errgroup.Group

	for _, a := range f.adapters {
		a := a
		g.Go(func() error {
			err := a.Update(path, contents)

			return err
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

// Read a file
func (f *Flysystem) Read(path string) ([]byte, error) {
	var g errgroup.Group
	contents := make([][]byte, len(f.adapters))

	for i, a := range f.adapters {
		i, a := i, a
		g.Go(func() error {
			bytes, err := a.Read(path)

			if err == nil {
				contents[i] = bytes
			}

			return err
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return contents[0], nil
}

// Rename a file
func (f *Flysystem) Rename(path string, newPath string) error {
	var g errgroup.Group

	for _, a := range f.adapters {
		a := a
		g.Go(func() error {
			err := a.Rename(path, newPath)

			return err
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

// Copy a file
func (f *Flysystem) Copy(path string, newPath string) error {
	var g errgroup.Group

	for _, a := range f.adapters {
		a := a
		g.Go(func() error {
			err := a.Copy(path, newPath)

			return err
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

// Delete a file
func (f *Flysystem) Delete(path string) error {
	var g errgroup.Group

	for _, a := range f.adapters {
		a := a
		g.Go(func() error {
			err := a.Delete(path)

			return err
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

// CreateDir creates a directory
func (f *Flysystem) CreateDir(dir string) error {
	var g errgroup.Group

	for _, a := range f.adapters {
		a := a
		g.Go(func() error {
			err := a.CreateDir(dir)

			return err
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

// DeleteDir deletes a directory
func (f *Flysystem) DeleteDir(dir string) error {
	var g errgroup.Group

	for _, a := range f.adapters {
		a := a
		g.Go(func() error {
			err := a.DeleteDir(dir)

			return err
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

// SetVisibility sets a file or directory to public or private
func (f *Flysystem) SetVisibility(path string, visibility string) error {
	var g errgroup.Group

	for _, a := range f.adapters {
		a := a
		g.Go(func() error {
			err := a.SetVisibility(path, visibility)

			return err
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
