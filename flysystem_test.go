package flysystem

import (
	"github.com/edwin-luijten/go_flysystem/adapter"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	a, err := adapter.NewLocal("./_testdata")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if _, err := os.Stat("./_testdata"); os.IsNotExist(err) {
		t.Log(err)
		t.Fail()
	}

	fs := New(a)

	err = fs.Write("test.txt", []byte("hello"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}


	if _, err := os.Stat("./_testdata/test.txt"); os.IsNotExist(err) {
		t.Log(err)
		t.Fail()
	}

	err = os.Remove("./_testdata/test.txt")
	if err != nil {
		panic(err)
	}
}