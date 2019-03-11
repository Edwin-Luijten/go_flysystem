package adapter

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// Local ...
type Local struct {
	BaseAdapter
	root    string
	lock    *sync.Mutex
	permMap map[string]map[string]os.FileMode
}

// FilePrivate represents 0600 file permissions
const FilePrivate = 0600

// FilePublic represents 0644 file permissions
const FilePublic = 0644

// DirPrivate represents 0700 directory permissions
const DirPrivate = 0700

// DirPublic represents 0755 directory permissions
const DirPublic = 0755

// NewLocal creates a new instance of Local
func NewLocal(root string) (Adapter, error) {
	var permMap = map[string]map[string]os.FileMode{
		"file": {},
		"dir":  {},
	}

	permMap["file"]["public"] = FilePublic
	permMap["file"]["private"] = FilePrivate
	permMap["dir"]["public"] = DirPublic
	permMap["dir"]["private"] = DirPrivate

	a := &Local{
		permMap: permMap,
	}

	err := a.ensureDirectory(root)
	if err != nil {
		return nil, err
	}

	a.SetPathPrefix(root)
	a.lock = &sync.Mutex{}

	return a, nil
}

// Write a new file
func (a *Local) Write(path string, contents []byte) error {
	a.lock.Lock()

	defer a.lock.Unlock()

	location := a.ApplyPathPrefix(path)

	dir, err := filepath.Abs(filepath.Dir(location))
	if err != nil {
		return err
	}

	err = a.ensureDirectory(dir)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(location, contents, FilePublic)
	if err != nil {
		return err
	}

	return nil
}

// Update a file
func (a *Local) Update(path string, contents []byte) error {
	a.lock.Lock()

	defer a.lock.Unlock()

	location := a.ApplyPathPrefix(path)

	err := ioutil.WriteFile(location, contents, FilePublic)
	if err != nil {
		return err
	}

	return nil
}

// Read a file
func (a *Local) Read(path string) ([]byte, error) {
	location := a.ApplyPathPrefix(path)

	return ioutil.ReadFile(location)
}

// Rename a file
func (a *Local) Rename(path string, newPath string) error {
	a.lock.Lock()

	defer a.lock.Unlock()

	location := a.ApplyPathPrefix(path)
	destination := a.ApplyPathPrefix(newPath)

	return os.Rename(location, destination)
}

// Copy a file
func (a *Local) Copy(path string, newPath string) error {
	a.lock.Lock()

	defer a.lock.Unlock()

	location := a.ApplyPathPrefix(path)
	destination := a.ApplyPathPrefix(newPath)

	// Get file permissions
	info, err := os.Stat(location)
	if err != nil {
		return err
	}

	input, err := ioutil.ReadFile(location)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(destination, input, info.Mode())
}

// Delete a file
func (a *Local) Delete(path string) error {
	a.lock.Lock()

	defer a.lock.Unlock()

	location := a.ApplyPathPrefix(path)

	return os.Remove(location)
}

// CreateDir creates a directory
func (a *Local) CreateDir(dir string) error {
	a.lock.Lock()

	defer a.lock.Unlock()

	location := a.ApplyPathPrefix(dir)

	return os.Mkdir(location, DirPublic)
}

// DeleteDir deletes a directory
func (a *Local) DeleteDir(dir string) error {
	a.lock.Lock()

	defer a.lock.Unlock()

	location := a.ApplyPathPrefix(dir)

	return os.RemoveAll(location)
}

// SetVisibility sets a file or directory to public or private
func (a *Local) SetVisibility(path string, visibility string) error {
	a.lock.Lock()

	defer a.lock.Unlock()

	location := a.ApplyPathPrefix(path)

	info, err := os.Stat(location)
	if err != nil {
		return err
	}

	var perm os.FileMode

	if info.IsDir() {
		perm = a.permMap["dir"][visibility]
	} else {
		perm = a.permMap["file"][visibility]
	}

	return os.Chmod(location, perm)
}

func (a *Local) ensureDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, DirPublic)
		if err != nil {
			return fmt.Errorf("impossible to create the root directory %s", dir)
		}
	}

	return nil
}
