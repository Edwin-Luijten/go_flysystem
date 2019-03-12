package adapter

import (
	"testing"
)

func TestBaseAdapter_SetPathPrefix(t *testing.T) {
	a := &BaseAdapter{}
	a.SetPathPrefix("data")

	if *a.pathPrefix != "data/" {
		t.Log("unexpected path prefix")
		t.Fail()
	}
}

func TestBaseAdapter_SetPathPrefixEmpty(t *testing.T) {
	a := &BaseAdapter{}
	a.SetPathPrefix("")

	if *a.pathPrefix != "/" {
		t.Log("unexpected path prefix")
		t.Fail()
	}
}

func TestBaseAdapter_ApplyPathPrefix(t *testing.T) {
	a := &BaseAdapter{}
	a.SetPathPrefix("data")
	p := a.ApplyPathPrefix("sub")

	if *a.pathPrefix != "data/" {
		t.Log("unexpected path prefix")
		t.Fail()
	}

	if p != "data/sub" {
		t.Log("unexpected path")
		t.Fail()
	}
}
