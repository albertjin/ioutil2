package ioutil2

import (
    "compress/gzip"
    "io"
    "os"
)

// A closure wrap for reading a gzip file. The stream opened is closed after use.
func ReadGzip(filename string, handle func(io.Reader)) (err error) {
    f, err := os.Open(filename)
    if err != nil {
        return
    }
    defer f.Close()

    z, err := gzip.NewReader(f)
    if err != nil {
        return
    }
    defer z.Close()

    handle(z)
    return
}

// A closure wrap for writing a gzip file. The stream opened is closed after use.
func WriteGzip(filename string, level int, handle func(io.Writer)) (err error) {
    f, err := os.Create(filename)
    if err != nil {
        return
    }
    defer f.Close()

    z, err := gzip.NewWriterLevel(f, level)
    if err != nil {
        return
    }
    defer z.Close()

    handle(z)
    return
}
