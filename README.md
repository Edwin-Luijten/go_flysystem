# Go Flysystem

Go Flysystem is a filesystem abstraction which allows you to easily swap out a local filesystem for a remote one.  

Inspired by: https://github.com/thephpleague/flysystem  

## Installation
``` go get github.com/edwin-luijten/go_flysystem ```  

## Usage

```go
import (
	"github.com/edwin-luijten/go_flysystem/adapter"
    flysystem "github.com/edwin-luijten/go_flysystem"
)

func main() {
    a, err := adapter.NewLocal("./_testdata")
    if err != nil {
    	panic(err)
    }
    
    fs := flysystem.New(a)
    
    // Write
    err = fs.Write("test.txt", []byte("hello"))
    if err != nil {
        t.Log(err)
        t.Fail()
    }
}
```
