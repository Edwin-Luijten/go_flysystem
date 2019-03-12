# Go Flysystem

[![Build Status](https://travis-ci.com/Edwin-Luijten/go_flysystem.svg?branch=master)](https://travis-ci.com/Edwin-Luijten/go_flysystem) 
[![Maintainability](https://api.codeclimate.com/v1/badges/6e48f895875537f89b42/maintainability)](https://codeclimate.com/github/Edwin-Luijten/go_flysystem/maintainability) 
[![Test Coverage](https://api.codeclimate.com/v1/badges/6e48f895875537f89b42/test_coverage)](https://codeclimate.com/github/Edwin-Luijten/go_flysystem/test_coverage)  

Go Flysystem is a filesystem abstraction which allows you to easily swap out a local filesystem for a remote one.  

Inspired by: https://github.com/thephpleague/flysystem  

## Installation
``` go get github.com/edwin-luijten/go_flysystem ```  

## Usage

### Adapters

```go
import (
    "github.com/edwin-luijten/go_flysystem/adapter"
)

func main() {
    a, err := adapter.NewLocal("./_testdata")
    if err != nil {
    	panic(err)
    }
    
    // Write
    err = a.Write("test.txt", []byte("hello"))
    if err != nil {
        t.Log(err)
        t.Fail()
    }
}
```

### Multiple adapters

```go
import (
    "github.com/edwin-luijten/go_flysystem/adapter"
)

func main() {
    a, err := adapter.NewLocal("./_testdata")
    if err != nil {
    	panic(err)
    }
    
    b, err = adapter.NewLocal("./_testdata")
    if err != nil {
        panic(err)
    }
    
    fs, err = flysystem.New(a, b)
    
    // Write
    err = fs.Write("test.txt", []byte("hello"))
    if err != nil {
        t.Log(err)
        t.Fail()
    }
}
```