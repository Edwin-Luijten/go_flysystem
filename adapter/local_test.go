package adapter

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestLocal_Write(t *testing.T) {
	fs, err := NewLocal("../_testdata")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = fs.Write("test.txt", []byte("hello world"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	bytes, err := ioutil.ReadFile("../_testdata/test.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if string(bytes) != "hello world" {
		t.Log("files does not contain: hello world")
		t.Fail()
	}

	err = os.Remove("../_testdata/test.txt")
	if err != nil {
		panic(err)
	}
}

func TestLocal_Update(t *testing.T) {
	fs, err := NewLocal("../_testdata")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = fs.Update("test.txt", []byte("hello"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	bytes, err := ioutil.ReadFile("../_testdata/test.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if string(bytes) != "hello" {
		t.Log("files does not contain: hello")
		t.Fail()
	}

	err = os.Remove("../_testdata/test.txt")
	if err != nil {
		panic(err)
	}
}

func TestLocal_Read(t *testing.T) {
	fs, err := NewLocal("../_testdata")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = fs.Update("test.txt", []byte("hello"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	contents, err := fs.Read("test.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if string(contents) != "hello" {
		t.Log("files does not contain: hello")
		t.Fail()
	}

	err = os.Remove("../_testdata/test.txt")
	if err != nil {
		panic(err)
	}
}

func TestLocal_Rename(t *testing.T) {
	fs, err := NewLocal("../_testdata")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = fs.Write("test.txt", []byte("hello world"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = fs.Rename("test.txt", "test_updated.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if _, err := os.Stat("../_testdata/test_updated.txt"); os.IsNotExist(err) {
		t.Log(err)
		t.Fail()
	}

	err = os.Remove("../_testdata/test_updated.txt")
	if err != nil {
		panic(err)
	}
}

func TestLocal_Copy(t *testing.T) {
	fs, err := NewLocal("../_testdata")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = fs.Write("test.txt", []byte("hello world"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = fs.Copy("test.txt", "test2.txt")

	if _, err := os.Stat("../_testdata/test2.txt"); os.IsNotExist(err) {
		t.Log(err)
		t.Fail()
	}

	contents, err := fs.Read("test.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
	}


	contents2, err := fs.Read("test2.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if string(contents) != string(contents2) {
		t.Log("file contents are not equal")
		t.Fail()
	}

	err = os.Remove("../_testdata/test.txt")
	if err != nil {
		panic(err)
	}

	err = os.Remove("../_testdata/test2.txt")
	if err != nil {
		panic(err)
	}
}

func TestLocal_Delete(t *testing.T) {
	fs, err := NewLocal("../_testdata")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

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

	if _, err := os.Stat("../_testdata/test.txt"); err == nil {
		t.Log(err)
		t.Fail()
	}
}

func TestLocal_CreateDir(t *testing.T) {
	fs, err := NewLocal("../_testdata")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = fs.CreateDir("subdir")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if _, err := os.Stat("../_testdata/subdir"); os.IsNotExist(err) {
		t.Log("directory does not exist")
		t.Fail()
	}

	err = os.Remove("../_testdata/subdir")
	if err != nil {
		panic(err)
	}
}

func TestLocal_DeleteDir(t *testing.T) {
	fs, err := NewLocal("../_testdata")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

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
}

func TestLocal_SetVisibility(t *testing.T) {
	fs, err := NewLocal("../_testdata")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

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

	info, err := os.Stat("../_testdata/test.txt")
	if err != nil {
		panic(err)
	}

	if info.Mode() != FilePrivate {
		t.Log(fmt.Println("wrong permissions: expected %i, got %i", FilePrivate, info.Mode()))
	}

	err = os.Remove("../_testdata/test.txt")
	if err != nil {
		panic(err)
	}
}