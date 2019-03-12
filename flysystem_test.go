package flysystem

import (
	"github.com/edwin-luijten/go_flysystem/adapter"
	"io/ioutil"
	"os"
	"testing"
)

func TestFlysystem_Write(t *testing.T) {
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

	err = os.RemoveAll("./_testdata/sub1")
	if err != nil {
		panic(err)
	}

	err = os.RemoveAll("./_testdata/sub2")
	if err != nil {
		panic(err)
	}
}

func TestFlysystem_Update(t *testing.T) {
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

	err = os.RemoveAll("./_testdata/sub1")
	if err != nil {
		panic(err)
	}

	err = os.RemoveAll("./_testdata/sub2")
	if err != nil {
		panic(err)
	}
}

func TestFlysystem_Read(t *testing.T) {
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
}
