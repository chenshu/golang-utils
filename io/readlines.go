package io

import (
    "io"
    "os"
    "bufio"
    "bytes"
)

func ReadLines(path string) (lines []string, err error) {
    var file *os.File
    if file, err = os.OpenFile(path, os.O_RDONLY, 0644); err != nil {
        return
    }
    defer file.Close()

    var reader *bufio.Reader = bufio.NewReader(file)
    var buffer *bytes.Buffer = bytes.NewBuffer(make([]byte, 0))
    var line []byte
    var isPrefix bool
    for {
        if line, isPrefix, err = reader.ReadLine(); err != nil {
            break
        }
        buffer.Write(line)
        if !isPrefix {
            lines = append(lines, buffer.String())
            buffer.Reset()
        }
    }
    if err == io.EOF {
        err = nil
    }
    return
}
