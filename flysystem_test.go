package flysystem

import (
	"fmt"
	"github.com/edwin-luijten/go_flysystem/adapter"
	"io/ioutil"
	"os"
	"testing"
)

func setup(t *testing.T) {
	if _, err := os.Stat("./_testdata/"); os.IsNotExist(err) {
		err := os.Mkdir("./_testdata/", os.ModePerm)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
	}
}

func teardown(t *testing.T) {
	err := os.RemoveAll("./_testdata/sub1")
	if err != nil {
		t.Log("Nothing to clean up")
	}

	err = os.RemoveAll("./_testdata/sub2")
	if err != nil {
		t.Log("Nothing to clean up")
	}
}

func TestFlysystem_Write(t *testing.T) {
	setup(t)
	defer teardown(t)

	a, err := adapter.NewLocal("./_testdata/sub1")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	b, err := adapter.NewLocal("./_testdata/sub2")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	fs := New(b, a)

	err = fs.Write("test.txt", []byte("hello"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if _, err := os.Stat("./_testdata/sub1/test.txt"); os.IsNotExist(err) {
		t.Log(err)
		t.Fail()
	}

	if _, err := os.Stat("./_testdata/sub2/test.txt"); os.IsNotExist(err) {
		t.Log(err)
		t.Fail()
	}

	os.Mkdir("./_testdata/sub1/err", 0644)
	os.Mkdir("./_testdata/sub2/err", 0644)

	err = fs.Write("err/error.txt", []byte("hello"))
	if err == nil {
		t.Log("expected an error")
		t.Fail()
	}
}

func TestFlysystem_Update(t *testing.T) {
	setup(t)
	defer teardown(t)

	a, err := adapter.NewLocal("./_testdata/sub1")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	b, err := adapter.NewLocal("./_testdata/sub2")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	fs := New(b, a)

	err = fs.Write("test.txt", []byte("hello world"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = fs.Update("test.txt", []byte("hello"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	bytes, err := ioutil.ReadFile("./_testdata/sub1/test.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if string(bytes) != "hello" {
		t.Log("files does not contain: hello")
		t.Fail()
	}

	bytes, err = ioutil.ReadFile("./_testdata/sub2/test.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if string(bytes) != "hello" {
		t.Log("files does not contain: hello")
		t.Fail()
	}
}

func TestFlysystem_Read(t *testing.T) {
	setup(t)
	defer teardown(t)

	a, err := adapter.NewLocal("./_testdata/sub1")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	b, err := adapter.NewLocal("./_testdata/sub2")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	fs := New(b, a)

	err = fs.Write("test.txt", []byte("hello"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	bytes, err := fs.Read("test.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if string(bytes) != "hello" {
		t.Log("files does not contain: hello")
		t.Fail()
	}

	os.Mkdir("./_testdata/sub1/err", os.ModePerm)
	os.Mkdir("./_testdata/sub2/err", os.ModePerm)

	err = fs.Write("err/test.txt", []byte("hello"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	os.Chmod("./_testdata/sub1/err/test.txt", 027)
	os.Chmod("./_testdata/sub2/err/test.txt", 027)

	bytes, err = fs.Read("err/test.txt")
	if err == nil {
		t.Log(fmt.Sprintf("expected an error but got '%s'", string(bytes)))
		t.Fail()
	}
}

func TestFlysystem_Rename(t *testing.T) {
	setup(t)
	defer teardown(t)

	a, err := adapter.NewLocal("./_testdata/sub1")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	b, err := adapter.NewLocal("./_testdata/sub2")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	fs := New(b, a)

	err = fs.Write("test.txt", []byte("hello"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = fs.Rename("test.txt", "test_updated.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if _, err := os.Stat("./_testdata/sub1/test_updated.txt"); os.IsNotExist(err) {
		t.Log(err)
		t.Fail()
	}

	if _, err := os.Stat("./_testdata/sub2/test_updated.txt"); os.IsNotExist(err) {
		t.Log(err)
		t.Fail()
	}
}

func TestFlysystem_Copy(t *testing.T) {
	setup(t)
	defer teardown(t)

	a, err := adapter.NewLocal("./_testdata/sub1")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	b, err := adapter.NewLocal("./_testdata/sub2")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	fs := New(b, a)

	err = fs.Write("test.txt", []byte("hello"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = fs.Copy("test.txt", "test_copied.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	contents, err := fs.Read("test.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	contents2, err := fs.Read("test_copied.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if string(contents) != string(contents2) {
		t.Log("file contents are not equal")
		t.Fail()
	}
}

func TestFlysystem_Delete(t *testing.T) {
	setup(t)
	defer teardown(t)

	a, err := adapter.NewLocal("./_testdata/sub1")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	b, err := adapter.NewLocal("./_testdata/sub2")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	fs := New(b, a)

	err = fs.Write("test.txt", []byte("hello world"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = fs.Delete("test.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if _, err := os.Stat("./_testdata/sub1/test.txt"); err == nil {
		t.Log(err)
		t.Fail()
	}

	if _, err := os.Stat("./_testdata/sub2/test.txt"); err == nil {
		t.Log(err)
		t.Fail()
	}
}

func TestFlysystem_CreateDir(t *testing.T) {
	setup(t)
	defer teardown(t)

	a, err := adapter.NewLocal("./_testdata/sub1")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	b, err := adapter.NewLocal("./_testdata/sub2")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	fs := New(b, a)

	err = fs.CreateDir("subdir")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if _, err := os.Stat("./_testdata/sub1/subdir"); os.IsNotExist(err) {
		t.Log("directory does not exist")
		t.Fail()
	}

	if _, err := os.Stat("./_testdata/sub2/subdir"); os.IsNotExist(err) {
		t.Log("directory does not exist")
		t.Fail()
	}
}

func TestFlysystem_DeleteDir(t *testing.T) {
	setup(t)
	defer teardown(t)

	a, err := adapter.NewLocal("./_testdata/sub1")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	b, err := adapter.NewLocal("./_testdata/sub2")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	fs := New(b, a)

	err = fs.CreateDir("subdir")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = fs.DeleteDir("subdir")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if _, err := os.Stat("./_testdata/sub1/subdir"); os.IsExist(err) {
		t.Log("directory still exists")
		t.Fail()
	}

	if _, err := os.Stat("./_testdata/sub2/subdir"); os.IsExist(err) {
		t.Log("directory still exists")
		t.Fail()
	}
}

func TestFlysystem_SetVisibility(t *testing.T) {
	setup(t)
	defer teardown(t)

	a, err := adapter.NewLocal("./_testdata/sub1")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	b, err := adapter.NewLocal("./_testdata/sub2")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	fs := New(b, a)

	err = fs.Write("test.txt", []byte("hello world"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = fs.SetVisibility("test.txt", "private")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	info, err := os.Stat("./_testdata/sub1/test.txt")
	if err != nil {
		panic(err)
	}

	if info.Mode() != adapter.FilePrivate {
		t.Log(fmt.Println("wrong permissions: expected %i, got %i", adapter.FilePrivate, info.Mode()))
	}

	info, err = os.Stat("./_testdata/sub2/test.txt")
	if err != nil {
		panic(err)
	}

	if info.Mode() != adapter.FilePrivate {
		t.Log(fmt.Println("wrong permissions: expected %i, got %i", adapter.FilePrivate, info.Mode()))
	}
}
