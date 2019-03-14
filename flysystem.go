package flysystem

import (
	"github.com/edwin-luijten/go_flysystem/adapter"
	"golang.org/x/sync/errgroup"
	"sync"
)

// Flysystem ...
type Flysystem struct {
	sync.Mutex
	wg       *sync.WaitGroup
	adapters []adapter.Adapter
}

// New creates a new instance with given adapters
func New(adapters ...adapter.Adapter) adapter.Adapter {
	return &Flysystem{
		adapters: adapters,
		wg:       &sync.WaitGroup{},
	}
}

// Write a new file
func (f *Flysystem) Write(path string, contents []byte) error {
	return f.runSync(func(a adapter.Adapter) error {
		return a.Write(path, contents)
	})
}

// Update a file
func (f *Flysystem) Update(path string, contents []byte) error {
	return f.runSync(func(a adapter.Adapter) error {
		return a.Update(path, contents)
	})
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
	return f.runSync(func(a adapter.Adapter) error {
		return a.Rename(path, newPath)
	})
}

// Copy a file
func (f *Flysystem) Copy(path string, newPath string) error {
	return f.runSync(func(a adapter.Adapter) error {
		return a.Copy(path, newPath)
	})
}

// Delete a file
func (f *Flysystem) Delete(path string) error {
	return f.runSync(func(a adapter.Adapter) error {
		return a.Delete(path)
	})
}

// CreateDir creates a directory
func (f *Flysystem) CreateDir(dir string) error {
	return f.runSync(func(a adapter.Adapter) error {
		return a.CreateDir(dir)
	})
}

// DeleteDir deletes a directory
func (f *Flysystem) DeleteDir(dir string) error {
	return f.runSync(func(a adapter.Adapter) error {
		return a.DeleteDir(dir)
	})
}

// SetVisibility sets a file or directory to public or private
func (f *Flysystem) SetVisibility(path string, visibility string) error {
	return f.runSync(func(a adapter.Adapter) error {
		return a.SetVisibility(path, visibility)
	})
}

func (f *Flysystem) runSync(action func(a adapter.Adapter) error) error {
	var g errgroup.Group

	for _, a := range f.adapters {
		a := a

		g.Go(func() error {
			err := action(a)

			return err
		})
	}

	return g.Wait()
}
